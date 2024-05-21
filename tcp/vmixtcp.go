package vmixtcp

import (
	"bufio"
	"context"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/FlowingSPDG/vmix-go/common/models"
)

var (
	ErrFailedToReadCommand = errors.New("failed to read command")
	ErrFailedToReadStatus  = errors.New("failed to read status")
	ErrStatusNotOK         = errors.New("status is not OK")
	ErrFailedToReadLength  = errors.New("failed to read length")
	ErrFailedToParseLength = errors.New("failed to parse length")
	ErrFailedToReadXML     = errors.New("failed to read XML")
	ErrFailedToUnmarshal   = errors.New("failed to unmarshal XML")
)

// Vmix main object
type Vmix struct {
	lock *sync.RWMutex // TODO: for goroutine safety
	conn net.Conn      // TCP connection

	callbacks callbacks
}

type callbacks struct {
	version     func(*VersionResponse)
	tally       func(*TallyResponse)
	function    func(*FunctionResponse)
	acts        func(*ActsResponse)
	xml         func(*XMLResponse)
	xmltext     func(*XMLTextResponse)
	subscribe   func(*SubscribeResponse)
	unsubscribe func(*UnsubscribeResponse)
}

// New vmix instance.
func New(dest string) (*Vmix, error) {
	c, err := net.Dial("tcp", dest+":8099")
	if err != nil {
		return nil, err
	}

	vmix := &Vmix{
		lock: &sync.RWMutex{},
		conn: c,
		callbacks: callbacks{
			version:     func(*VersionResponse) {},
			tally:       func(*TallyResponse) {},
			function:    func(*FunctionResponse) {},
			acts:        func(*ActsResponse) {},
			xml:         func(*XMLResponse) {},
			xmltext:     func(*XMLTextResponse) {},
			subscribe:   func(*SubscribeResponse) {},
			unsubscribe: func(*UnsubscribeResponse) {},
		},
	}

	return vmix, nil
}

func (v *Vmix) readCommand(rd *bufio.Reader) (string, error) {
	command, err := rd.ReadString(' ')
	if err != nil {
		return "", ErrFailedToReadCommand
	}
	command = strings.TrimSpace(command)
	return command, nil
}

func (v *Vmix) readStatus(rd *bufio.Reader) error {
	status, err := rd.ReadString(' ')
	if err != nil {
		return ErrFailedToReadStatus
	}
	status = strings.TrimSpace(status)
	if status != statusOK {
		return ErrStatusNotOK
	}
	return nil
}

func (v *Vmix) readLength(rd *bufio.Reader) (int, error) {
	length, _, err := rd.ReadLine()
	if err != nil {
		return 0, ErrFailedToReadLength
	}
	i, err := strconv.Atoi(strings.TrimSpace(string(length)))
	if err != nil {
		return 0, ErrFailedToParseLength
	}
	return i, nil
}

func (v *Vmix) readXML(length int, rd *bufio.Reader) (*models.APIXML, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(rd, b); err != nil {
		return nil, ErrFailedToReadXML
	}
	api := models.APIXML{}
	if err := xml.Unmarshal(b, &api); err != nil {
		return nil, ErrFailedToUnmarshal
	}
	return &api, nil
}

// Run Start vMix TCP API Instance
func (v *Vmix) Run(ctx context.Context) error {
	r := bufio.NewReader(v.conn)
	for {
		command, err := v.readCommand(r)
		if err != nil {
			log.Println("Failed to read command:", err)
			continue
		}

		switch command {
		case commandVersion:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			version, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read response:", err)
				continue
			}
			resp := VersionResponse{
				Version: string(version),
			}
			v.callbacks.version(&resp)

		case commandTally:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			tallies, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read tallies:", err)
				continue
			}
			resp := TallyResponse{
				Tally: encodeTallies(tallies),
			}
			v.callbacks.tally(&resp)

		case commandFunction:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			response, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read response:", err)
				continue
			}
			resp := FunctionResponse{
				Response: string(response),
			}
			v.callbacks.function(&resp)

		case commandActs:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			response, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read response:", err)
				continue
			}
			resp := ActsResponse{
				Response: string(response),
			}
			v.callbacks.acts(&resp)

		case commandXML:
			length, err := v.readLength(r)
			if err != nil {
				log.Println("Unknown parse XML length:", err)
				continue
			}
			api, err := v.readXML(length, r)
			if err != nil {
				log.Println("Failed to read XML:", err)
				continue
			}

			resp := XMLResponse{
				XML: api,
			}
			v.callbacks.xml(&resp)

		case commandXMLText:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			xmltext, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := XMLTextResponse{
				XMLText: string(xmltext),
			}
			v.callbacks.xmltext(&resp)

		case commandSubscribe:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			respCommand, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := SubscribeResponse{
				Command: string(respCommand),
			}
			v.callbacks.subscribe(&resp)

		case commandUnsubscribe:
			if err := v.readStatus(r); err != nil {
				log.Println("Failed to read status:", err)
				continue
			}
			respCommand, _, err := r.ReadLine()
			if err != nil {
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := UnsubscribeResponse{
				Command: string(respCommand),
			}
			v.callbacks.unsubscribe(&resp)

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

// Close connection. Calls QUIT command before connection closure.
func (v *Vmix) Close() error {
	if err := v.Quit(); err != nil {
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
func (v *Vmix) XMLPath(xpath string) error {
	_, err := v.conn.Write(newXMLTEXTCommand(xpath))
	if err != nil {
		return err
	}
	return nil
}

// TALLY Get tally status
func (v *Vmix) Tally() error {
	_, err := v.conn.Write(newTALLYCommand())
	if err != nil {
		return err
	}
	return nil
}

// FUNCTION Send function
func (v *Vmix) Function(funcname string) error {
	_, err := v.conn.Write(newFUNCTIONCommand(funcname))
	if err != nil {
		return err
	}
	return nil
}

// SUBSCRIBE Event
func (v *Vmix) Subscribe(event, option string) error {
	_, err := v.conn.Write(newSUBSCRIBECommand(event, option))
	if err != nil {
		return err
	}
	return nil
}

// UNSUBSCRIBE from event.
func (v *Vmix) Unsubscribe(event string) error {
	_, err := v.conn.Write(newUNSUBSCRIBECommand(event))
	if err != nil {
		return err
	}
	return nil
}

// QUIT Sends QUIT sigal
func (v *Vmix) Quit() error {
	_, err := v.conn.Write(newQUITCommand())
	if err != nil {
		return err
	}
	return nil
}
