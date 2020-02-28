package vmixgo

import (
	"testing"
)

// General vMix features
// vmix.go or models.go
func TestRefresh(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Refresh()
	if err != nil {
		t.Fatal(err)
	}
}

// transition.go
func TestCut(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Cut(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFade(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Fade(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

//browser.go
func TestBrowserBack(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserBack(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserForward(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserForward(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserKeyboardDisabled(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserKeyboardDisabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserKeyboardEnabled(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserKeyboardEnabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserMouseDisabled(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserMouseDisabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserMouseEnabled(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserMouseEnabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserNavigate(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserNavigate(0, "http://google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserReload(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserReload(0)
	if err != nil {
		t.Fatal(err)
	}
}
