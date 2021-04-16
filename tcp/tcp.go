package vmixtcp

import (
	"bufio"
	"fmt"
	"net"
)

var (
	// Terminate letter
	Terminate = []byte("\r\n")
)

// Client vMix TCP API main object
type Client struct {
	host         string
	port         int
	conn         net.Conn
	cbHandler    map[string][]func(*Response)
	tallyHandler []func(*TallyResponse)
}

// NewClient Get new instance of vMix TCP API Client. Host:localhost Port:8099
func NewClient(host string, port int) (*Client, error) {
	if port == 0 || host == "" {
		return nil, fmt.Errorf("Invalid arguments")
	}
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(c)
	sc.Split(scanCRLF)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return &Client{
		host:         host,
		port:         port,
		conn:         c,
		cbHandler:    make(map[string][]func(*Response)),
		tallyHandler: make([]func(*TallyResponse), 0),
	}, nil
}

// Close connection.
func (c *Client) Close() error {
	if _, err := c.conn.Write([]byte("QUIT")); err != nil {
		return err
	}
	return c.conn.Close()
}
