package vmixgo

import (
	"testing"
)

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
