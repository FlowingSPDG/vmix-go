package vmixgo

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

// NewVmixHTTP Creates new vMix HTTP API instance
func NewVmixHTTP(host string, port int) (*VmixHTTPClient, error) {
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
	v := VmixHTTPClient{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal XML... %v", err)
	}
	v.addr = u
	return &v, nil
}
