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
	"time"

	"github.com/FlowingSPDG/vmix-go/common/models"
)

var (
	ErrAlreadyConnected           = errors.New("already connected")
	ErrDisconnected               = errors.New("disconnected")
	ErrFailedToInitiateConnection = errors.New("failed to initiate connection")
	ErrNotConnected               = errors.New("not connected to vMix")
	ErrFailedToReadCommand        = errors.New("failed to read command")
	ErrFailedToReadStatus         = errors.New("failed to read status")
	ErrStatusNotOK                = errors.New("status is not OK")
	ErrFailedToReadLength         = errors.New("failed to read length")
	ErrFailedToParseLength        = errors.New("failed to parse length")
	ErrFailedToReadXML            = errors.New("failed to read XML")
	ErrFailedToUnmarshal          = errors.New("failed to unmarshal XML")
)

// vmix main instance.
type vmix struct {
	dest      string // vMix destination
	connected bool
	mutex     *sync.Mutex   // for goroutine safety
	conn      net.Conn      // TCP connection
	reader    *bufio.Reader // buffered reader

	callbacks callbacks
}

// Vmix
// vMix TCP API Main interface.
// You always need to call Connect() before Run().
// You can call other methods before Run() since connection buffer holds the response from vMix but method needs be called after initiating TCP connection.
// So I highly recommend you to send commands inside of OnVersion callback since Version command is the first command that vMix sends after connection establish.
type Vmix interface {
	IsConnected() bool

	Connect() error                // Connects vMix TCP API. You need to call this before Run().
	Run(ctx context.Context) error // Start Receiving TCP packet with vMix. You need to call this after Connect(). You can call other methods before Run() since connection buffer holds the response from vMix.
	Close() error                  // Close connection. Wraps Quit() and conn.Close().

	// Send commands
	Tally() error
	Function(name string, query string) error
	Acts(name string, input ...int) error
	XML() error
	XMLText(xpath string) error
	Subscribe(event, command string) error
	Unsubscribe(command string) error
	Quit() error // Normally you do not need to call this. Instead, call Close() for connection closure.

	// Callbacks. Since vMix TCP API does not respond to the command, you need to register callbacks to receive responses.
	OnVersion(func(*VersionResponse))
	OnTally(func(*TallyResponse))
	OnFunction(func(*FunctionResponse))
	OnActs(func(*ActsResponse))
	OnXML(func(*XMLResponse))
	OnXMLText(func(*XMLTextResponse))
	OnSubscribe(func(*SubscribeResponse))
	OnUnsubscribe(func(*UnsubscribeResponse))
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
func New(dest string) Vmix {
	return &vmix{
		dest:      dest,
		connected: false,
		mutex:     &sync.Mutex{},
		conn:      nil,
		reader:    nil,
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
}

func (v *vmix) IsConnected() bool {
	return v.connected
}

func (v *vmix) Connect() error {
	if v.connected {
		return ErrAlreadyConnected
	}
	// Start connecting to vmix TCP API
	host := net.JoinHostPort(v.dest, "8099")
	d := net.Dialer{
		Timeout: 5 * time.Second,
	}
	conn, err := d.Dial("tcp", host)
	if err != nil {
		return ErrFailedToInitiateConnection
	}
	v.connected = true
	v.conn = conn
	// Since bufio.Reader buffers the response, client can call other methods between Connect() and Run().
	v.reader = bufio.NewReader(v.conn)
	return nil
}

func (v *vmix) readCommand() (string, error) {
	command, err := v.reader.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return "", ErrDisconnected
		}
		return "", ErrFailedToReadCommand
	}
	command = strings.TrimSpace(command)
	return command, nil
}

func (v *vmix) readLine() (string, error) {
	line, _, err := v.reader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return "", ErrDisconnected
		}
		return "", err
	}
	return string(line), nil
}

func (v *vmix) readStatus() error {
	status, err := v.reader.ReadString(' ')
	if err != nil {
		if err == io.EOF {
			return ErrDisconnected
		}
		return ErrFailedToReadStatus
	}
	status = strings.TrimSpace(status)
	if status != statusOK {
		s, err := v.readLine()
		if err != nil {
			return err
		}
		return errors.Join(ErrStatusNotOK, errors.New(s))
	}
	return nil
}

func (v *vmix) readLength() (int, error) {
	length, err := v.readLine()
	if err != nil {
		if err == io.EOF {
			return 0, ErrDisconnected
		}
		return 0, ErrFailedToReadLength
	}
	i, err := strconv.Atoi(strings.TrimSpace(string(length)))
	if err != nil {
		return 0, ErrFailedToParseLength
	}
	return i, nil
}

func (v *vmix) readXML(length int) (*models.APIXML, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(v.reader, b); err != nil {
		if err == io.EOF {
			return nil, ErrDisconnected
		}
		return nil, ErrFailedToReadXML
	}
	api := models.APIXML{}
	if err := xml.Unmarshal(b, &api); err != nil {
		return nil, ErrFailedToUnmarshal
	}
	return &api, nil
}

// Run Start vMix TCP API Instance
func (v *vmix) Run(ctx context.Context) error {
	if !v.connected {
		return ErrNotConnected
	}

	go func() {
		<-ctx.Done()
		if err := v.Close(); err != nil {
			log.Println("Failed to close connection:", err)
		}
	}()

	for {
		command, err := v.readCommand()
		if err != nil {
			if err == ErrDisconnected {
				return v.Close()
			}
			log.Println("Failed to read command:", err)
			continue
		}

		switch command {
		case commandVersion:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			version, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read response:", err)
				continue
			}
			resp := VersionResponse{
				Version: string(version),
			}
			v.callbacks.version(&resp)

		case commandTally:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			tallies, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read tallies:", err)
				continue
			}
			resp := TallyResponse{
				Tally: encodeTallies([]byte(tallies)),
			}
			v.callbacks.tally(&resp)

		case commandFunction:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			response, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read response:", err)
				continue
			}
			resp := FunctionResponse{
				Response: string(response),
			}
			v.callbacks.function(&resp)

		case commandActs:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			response, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read response:", err)
				continue
			}
			resp := ActsResponse{
				Response: string(response),
			}
			v.callbacks.acts(&resp)

		case commandXML:
			length, err := v.readLength()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Unknown parse XML length:", err)
				continue
			}
			api, err := v.readXML(length)
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read XML:", err)
				continue
			}

			resp := XMLResponse{
				XML: api,
			}
			v.callbacks.xml(&resp)

		case commandXMLText:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			xmltext, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := XMLTextResponse{
				XMLText: string(xmltext),
			}
			v.callbacks.xmltext(&resp)

		case commandSubscribe:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			respCommand, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := SubscribeResponse{
				Command: string(respCommand),
			}
			v.callbacks.subscribe(&resp)

		case commandUnsubscribe:
			if err := v.readStatus(); err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read status:", err)
				continue
			}
			respCommand, err := v.readLine()
			if err != nil {
				if err == ErrDisconnected {
					v.Close()
					return err
				}
				log.Println("Failed to read XMLTEXT:", err)
				continue
			}
			resp := UnsubscribeResponse{
				Command: string(respCommand),
			}
			v.callbacks.unsubscribe(&resp)

		default:
		}
	}
}

func (v *vmix) send(command []byte) error {
	if !v.connected {
		return ErrNotConnected
	}
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if _, err := v.conn.Write(command); err != nil {
		return err
	}
	return nil
}

// Close connection. Calls QUIT command before connection closure.
func (v *vmix) Close() error {
	if err := v.Quit(); err != nil {
		return err
	}
	if err := v.conn.Close(); err != nil {
		return err
	}
	v.connected = false
	// ?
	// v.conn = nil
	// v.reader = nil
	return nil
}

// TALLY Get tally status
func (v *vmix) Tally() error {
	if err := v.send(newTallyCommand()); err != nil {
		return err
	}
	return nil
}

// FUNCTION Send function
func (v *vmix) Function(name string, query string) error {
	if err := v.send(newFunctionCommand(name, query)); err != nil {
		return err
	}
	return nil
}

// Acts Send ACTS command
func (v *vmix) Acts(name string, input ...int) error {
	if err := v.send(newActsCommand(name, input...)); err != nil {
		return err
	}
	return nil
}

// XML Gets XML data. Same as HTTP API.
func (v *vmix) XML() error {
	if err := v.send(newXMLCommand()); err != nil {
		return err
	}
	return nil
}

// XMLText Gets XML data from specified XPATH
func (v *vmix) XMLText(xpath string) error {
	if err := v.send(newXMLTEXTCommand(xpath)); err != nil {
		return err
	}
	return nil
}

// SUBSCRIBE Event
func (v *vmix) Subscribe(event, command string) error {
	if err := v.send(newSubscribeCommand(event, command)); err != nil {
		return err
	}
	return nil
}

// UNSUBSCRIBE from event.
func (v *vmix) Unsubscribe(command string) error {
	if err := v.send(newUnsubscribeCommand(command)); err != nil {
		return err
	}
	return nil
}

// QUIT Sends QUIT sigal
func (v *vmix) Quit() error {
	if err := v.send(newQuitCommand()); err != nil {
		return err
	}
	return nil
}
