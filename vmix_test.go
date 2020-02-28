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

// datasources.go
func TestDataSourceAutoNextOff(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOff("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceAutoNextOn(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOn("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceAutoNextOnOff(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOnOff("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceNextRow(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceNextRow("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourcePreviousRow(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourcePreviousRow("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceSelectRow(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceSelectRow("")
	if err != nil {
		t.Fatal(err)
	}
}

// general.go
func TestActivatorRefresh(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ActivatorRefresh()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallManagerShowHide(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.CallManagerShowHide()
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeyPress(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.KeyPress("A")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendKeys(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SendKeys("{ENTER}")
	if err != nil {
		t.Fatal(err)
	}
}

// ndi.go
func TestNDICommand(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDICommand(0, "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDISelectSourceByIndex(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDISelectSourceByIndex(0, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDISelectSourceByName(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDISelectSourceByName(0, "NDI")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDIStartRecording(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDIStartRecording(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDIStopRecording(t *testing.T) {
	vmix, err := NewVmix("http://localhost:8088")
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDIStopRecording(0)
	if err != nil {
		t.Fatal(err)
	}
}
