package vmixhttp

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/FlowingSPDG/vmix-go/common/models"
)

// Client vMix HTTP API main object
type Client struct {
	addr *url.URL `xml:"-"` // vmix API destination.

	// Available informations in /api endpoint (XML).
	XMLName xml.Name `xml:"vmix"`
	Version string   `xml:"version"` // vmix Version. e.g. "23.0.0.31"
	Edition string   `xml:"edition"` // vmix Edition. e.g. "4K"
	Preset  string   `xml:"preset"`  // vmix profile directory. e.g. "C:\my-profile.vmix"
	// Scenes slice
	Inputs struct {
		Input []models.Input `xml:"input"`
	} `xml:"inputs"`
	// Overlays slice
	Overlays struct {
		Overlay []models.Overlay `xml:"overlay"`
	} `xml:"overlays"`
	Preview       uint `xml:"preview"`     // Preview scene number
	Active        uint `xml:"active"`      // Active scene number
	IsFadeToBlack bool `xml:"fadeToBlack"` // FTB activated or not
	// vmix transition
	Transitions struct {
		Transition []models.Transition `xml:"transition"`
	} `xml:"transitions"`
	Recording   bool `xml:"recording"`   // Recording enabled
	External    bool `xml:"external"`    // External output enabled
	Streaming   bool `xml:"streaming"`   // RTMP Streaming enabled
	PlayList    bool `xml:"playList"`    // Playlist enabled
	MultiCorder bool `xml:"multiCorder"` // MultiCorder enabled
	FullScreen  bool `xml:"fullscreen"`  // FullScreen enabled
	Mix         []struct {
		Number  uint `xml:"number,attr"` // Mix number(2~4)
		Preview uint `xml:"preview"`
		Active  uint `xml:"active"`
	} `xml:"mix"`
	// Audio?
	Audios struct {
		Master []models.Audio `xml:"master"`
	} `xml:"audio"`
}

// SendFunction sends request to /api?Function=funcname&Key=Value...
func (v *Client) SendFunction(funcname string, params map[string]string) error {
	q := v.addr.Query()
	q.Add("Function", funcname)
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
	}
	req := *v.addr
	queries := q.Encode()
	req.RawQuery = queries
	resp, err := http.Get(req.String())
	if err != nil {
		return fmt.Errorf("Failed to send function... %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusInternalServerError {
		return fmt.Errorf("vMix returned Internal error")
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to Read body... %v", err)
	}
	return nil
}

// Refresh Inputs
func (v *Client) Refresh() error {
	resp, err := http.Get(v.addr.String())
	if err != nil {
		return fmt.Errorf("Failed to connect vmix... %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to Read body... %v", err)
	}
	vnew := Client{}
	//fmt.Printf("body : %v\n", string(body))
	err = xml.Unmarshal(body, &vnew)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal XML... %v", err)
	}
	vnew.addr = v.addr
	v = &vnew
	return nil
}

// NewClient Creates new vMix HTTP API instance
func NewClient(host string, port int) (*Client, error) {
	u := &url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        &url.Userinfo{},
		Host:        fmt.Sprintf("%s:%d", host, port),
		Path:        "",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	u.Path = path.Join(u.Path, "/api")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect vmix... %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to Read body... %v", err)
	}
	v := Client{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal XML... %v", err)
	}
	v.addr = u
	return &v, nil
}
