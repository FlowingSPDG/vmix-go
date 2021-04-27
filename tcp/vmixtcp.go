package vmixtcp

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// Vmix main object
type Vmix struct {
	conn      net.Conn
	cbhandler map[string][]func(*Response)
}

// New vmix instance. TODO:Support context
func New(dest string) (*Vmix, error) {
	vmix := &Vmix{}
	vmix.cbhandler = make(map[string][]func(*Response))
	c, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}

	vmix.conn = c

	return vmix, nil
}

// Run Start vMix TCP API Instance
func (v *Vmix) Run(ctx context.Context) error {
	r := bufio.NewReader(v.conn)
	for {
		var prefix bool
		var line []byte
		var err error
		line, prefix, err = r.ReadLine()
		if err != nil {
			log.Println("Failed to read line:", err)
			return err
		}
		if prefix {
			l, _, e := r.ReadLine()
			if e != nil {
				log.Println("Failed to read line:", e)
				return e
			}
			line = append(line, l...)
		}

		resp := string(line)
		// log.Println("resp:", resp)

		resps := strings.Split(resp, " ")
		if len(resps) <= 0 {
			return fmt.Errorf("Failed to read line:%v", resps)
		}
		switch resps[0] {
		case EVENT_TALLY:
			if len(resps) != 3 {
				log.Println("Unknown response length:", resp)
				continue
			}
			for _, f := range v.cbhandler[EVENT_TALLY] {
				go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: resps[2], Data: ""})
			}
		case EVENT_FUNCTION:
			if len(resps) != 3 {
				log.Println("Unknown response length:", resp)
				continue
			}
			for _, f := range v.cbhandler[EVENT_FUNCTION] {
				go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: resps[2], Data: ""})
			}
		case EVENT_ACTS:
			if strings.HasPrefix(resp, EVENT_ACTS+" "+STATUS_ER) {
				log.Println("Failed to load ACTS:", resp)
				continue
			} else if strings.HasPrefix(resp, EVENT_ACTS+" "+STATUS_OK) {
				data := strings.ReplaceAll(resp, EVENT_ACTS+" "+STATUS_OK, "")
				for _, f := range v.cbhandler[EVENT_ACTS] {
					go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: data, Data: ""})
				}
			}
		case EVENT_XML:
			if len(resps) != 2 {
				log.Println("Unknown response length:", resp)
				continue
			}
			length, err := strconv.Atoi(resps[1])
			if err != nil {
				log.Println("Unknown parse XML length:", err)
				continue
			}
			b := make([]byte, length)
			read, err := v.conn.Read(b)
			if err != nil {
				log.Println("Failed to read XML", err)
				continue
			}
			for _, f := range v.cbhandler[EVENT_XML] {
				go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: string(b[:read]), Data: ""})
			}
		case EVENT_XMLTEXT:
			if strings.HasPrefix(resp, EVENT_XMLTEXT+" "+STATUS_ER) {
				log.Println("Failed to load XMLTEXT:", resp)
				continue
			} else if strings.HasPrefix(resp, EVENT_XMLTEXT+" "+STATUS_OK) {
				data := strings.ReplaceAll(resp, EVENT_XMLTEXT+" "+STATUS_OK, "")
				for _, f := range v.cbhandler[EVENT_XMLTEXT] {
					go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: data, Data: ""})
				}
			}
		case EVENT_SUBSCRIBE:
			if len(resps) != 3 && len(resps) != 4 {
				log.Println("Unknown response length:", resp)
				continue
			}
			for _, f := range v.cbhandler[EVENT_SUBSCRIBE] {
				r := &Response{Command: resps[0], StatusOrLength: resps[1], Response: resps[2], Data: ""}
				if len(resps) == 4 {
					r.Data = resps[3]
				}
				go f(r)
			}
		case EVENT_UNSUBSCRIBE:
			if len(resps) != 3 {
				log.Println("Unknown response length:", resp)
				continue
			}
			for _, f := range v.cbhandler[EVENT_UNSUBSCRIBE] {
				go f(&Response{Command: resps[0], StatusOrLength: resps[1], Response: resps[2], Data: ""})
			}
			log.Println("Unknown response:", resp)
		default:
		}

		select {
		case <-ctx.Done():
			if err := v.Close(); err != nil {
				return err
			}
			return ctx.Err()
		default:
			continue
		}
	}
}

func (v *Vmix) Write(b []byte) (n int, err error) {
	return v.conn.Write(b)
}

// Close connection
func (v *Vmix) Close() error {
	if err := v.QUIT(); err != nil {
		return err
	}
	if err := v.conn.Close(); err != nil {
		return err
	}
	return nil
}

// XML Gets XML data. Same as HTTP API.
func (v *Vmix) XML() error {
	_, err := v.conn.Write(newXMLCommand())
	if err != nil {
		return err
	}
	return nil
}

// XMLPATH Gets XML data from specified XPATH
func (v *Vmix) XMLPATH(xpath string) error {
	_, err := v.conn.Write(newXMLTEXTCommand(xpath))
	if err != nil {
		return err
	}
	return nil
}

// TALLY Get tally status
func (v *Vmix) TALLY() error {
	_, err := v.conn.Write(newTALLYCommand())
	if err != nil {
		return err
	}
	return nil
}

// FUNCTION Send function
func (v *Vmix) FUNCTION(funcname string) error {
	_, err := v.conn.Write(newFUNCTIONCommand(funcname))
	if err != nil {
		return err
	}
	return nil
}

// SUBSCRIBE Event
func (v *Vmix) SUBSCRIBE(event, option string) error {
	_, err := v.Write(newSUBSCRIBECommand(event, option))
	if err != nil {
		return err
	}
	return nil
}

// UNSUBSCRIBE from event.
func (v *Vmix) UNSUBSCRIBE(event string) error {
	_, err := v.Write(newUNSUBSCRIBECommand(event))
	if err != nil {
		return err
	}
	return nil
}

// QUIT Sends QUIT sigal
func (v *Vmix) QUIT() error {
	_, err := v.conn.Write(newQUITCommand())
	if err != nil {
		return err
	}
	return nil
}
