package vmixtcp

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// Terminate letter
	Terminate = "\r\n"
)

// Vmix main object
type Vmix struct {
	conn       net.Conn
	subscriber net.Conn
	cbhandler  map[string][]func(*Response)
}

// New vmix instance. TODO:Support context
func New(dest string) (*Vmix, error) {
	vmix := &Vmix{}
	vmix.cbhandler = make(map[string][]func(*Response))
	c, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}

	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := c.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return nil, fmt.Errorf("Unknown ERR : %v", Resps[2:])
	}

	log.Printf("vMix TCP API Initialized... : %s\n", Resp)

	vmix.conn = c

	// SUBSCRIBE related...
	subscriber, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}
	vmix.subscriber = subscriber

	return vmix, nil
}

func (v *Vmix) Write(b []byte) (n int, err error) {
	fmt.Printf("Sending %s\n", b)
	return v.conn.Write(append(b, []byte(Terminate)...))
}

// Close connection
func (v *Vmix) Close() error {
	if err := v.QUIT(); err != nil {
		return err
	}
	if err := v.conn.Close(); err != nil {
		return err
	}
	if err := v.subscriber.Close(); err != nil {
		return err
	}
	return nil
}

// Run start
func (v *Vmix) Run(ctx context.Context) error {
	log.Println("RUNNING")
	reader := bufio.NewReader(v.subscriber)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("EOF")
			return err
		} else if err != nil {
			log.Printf("Unknown error on subscriber : %v\n", err)
			return err
		}
		data = strings.ReplaceAll(data, Terminate, " ")
		// log.Println("SUBSCRIBER DATA :", data)
		responses := strings.Split(string(data), " ") // Split response by space
		if len(responses) < 3 {
			log.Println("Unknown length data :", responses)
			continue
		}
		resp := &Response{}
		resp.Command = responses[0]
		resp.StatusOrLength = responses[1]
		resp.Response = responses[2]
		if len(responses) >= 4 {
			resp.Data = responses[3]
		}
		if v, ok := v.cbhandler[resp.Command]; ok {
			for _, f := range v {
				go f(resp)
			}
		}

		select {
		case <-ctx.Done():
			log.Println("DONE")
			if err := v.Close(); err != nil {
				return err
			}
			return ctx.Err()
		default:
			continue
		}
	}
}

// XML Gets XML data. Same as HTTP API.
func (v *Vmix) XML() (string, error) {
	_, err := v.conn.Write([]byte(EVENT_XML + Terminate))
	if err != nil {
		return "", err
	}
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	sc := bufio.NewScanner(v.conn)
	sc.Scan()
	Resp := strings.Split(sc.Text(), " ")
	if len(Resp) != 2 {
		return "", fmt.Errorf("Unknown XML Response length: %v", Resp)
	}
	size, err := strconv.Atoi(Resp[1]) // get XML size
	if err != nil {
		return "", fmt.Errorf("Unknown XML Response: %v", Resp)
	}

	BodyBuffer := make([]byte, size) // allocate memory
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, err = v.conn.Read(BodyBuffer)
	if err != nil {
		return "", err
	}
	r := strings.ReplaceAll(strings.ReplaceAll(string(BodyBuffer), string('\r'), ""), string('\n'), "")
	return r, nil
}

// XMLPATH Gets XML data from specified XPATH
func (v *Vmix) XMLPATH(xpath string) (string, error) {
	_, err := v.conn.Write([]byte(EVENT_XMLTEXT + " " + xpath + Terminate))
	if err != nil {
		return "", err
	}
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := v.conn.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

// TALLY Get tally status
func (v *Vmix) TALLY() (string, error) {
	_, err := v.conn.Write([]byte(EVENT_TALLY + Terminate))
	if err != nil {
		return "", err
	}
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1011) // Maximum possible length = 9 + 1000 + 2 = 1011 bytes
	RespLength, _ := v.conn.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[4:])
	}

	return Resp, nil
}

// FUNCTION Send function
func (v *Vmix) FUNCTION(funcname string) (string, error) {
	_, err := v.conn.Write([]byte(fmt.Sprintf("%s %s%s", EVENT_FUNCTION, funcname, Terminate)))
	if err != nil {
		return "", err
	}
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := v.conn.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

// SUBSCRIBE Event
func (v *Vmix) SUBSCRIBE(command string) (string, error) {
	_, err := v.subscriber.Write([]byte(fmt.Sprintf("%s %s%s", EVENT_SUBSCRIBE, command, Terminate)))
	if err != nil {
		return "", err
	}
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := v.subscriber.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

// UNSUBSCRIBE from event.
func (v *Vmix) UNSUBSCRIBE(command string) (string, error) {
	_, err := v.subscriber.Write([]byte(fmt.Sprintf("%s %s%s", EVENT_UNSUBSCRIBE, command, Terminate)))
	if err != nil {
		return "", err
	}
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := v.subscriber.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	if Resps[1] != "OK" {
		return "", fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return Resp, nil
}

// QUIT Sends QUIT sigal
func (v *Vmix) QUIT() error {
	_, err := v.conn.Write([]byte(fmt.Sprintf("%s %s", EVENT_QUIT, Terminate)))
	if err != nil {
		return err
	}
	v.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	RespBuffer := make([]byte, 1024)
	RespLength, _ := v.conn.Read(RespBuffer)

	Resp := strings.ReplaceAll(string(RespBuffer[:RespLength]), Terminate, "")
	Resps := strings.Split(Resp, " ")

	// check slice length
	if len(Resps) < 2 {
		return fmt.Errorf("Unknown response length : %v", Resps)
	}
	if Resps[1] != "OK" {
		return fmt.Errorf("Unknown ERR : %v", Resps[3:])
	}

	return nil
}
