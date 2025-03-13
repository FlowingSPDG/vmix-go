package vmixhttp

import (
	"testing"
)

// TODO: vMix HTTPのインスタンスをテスト開始時に定義して使いまわす

// General vMix features
func TestRefresh(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}

	if err := vmix.Refresh(); err != nil {
		t.Fatal(err)
	}
}

// Audio tests
func TestAudio(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Audio(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAudioBus(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.AudioBus(0, "Master")
	if err != nil {
		t.Fatal(err)
	}
}

// Layer tests
func TestLayerOff(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.LayerOff(0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLayerOn(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.LayerOn(0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

// MultiView tests
func TestMultiViewOverlay(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.MultiViewOverlay(0, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetMultiViewOverlay(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetMultiViewOverlay(0, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
}

// VideoCall tests
func TestVideoCallConnect(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VideoCallConnect(0, "test", "password")
	if err != nil {
		t.Fatal(err)
	}
}

func TestZoomJoinMeeting(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ZoomJoinMeeting(0, "123456789", "password")
	if err != nil {
		t.Fatal(err)
	}
}

// VideoDelay tests
func TestVideoDelayStartRecording(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VideoDelayStartRecording(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoDelayStopRecording(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VideoDelayStopRecording(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
}

// Replay tests
func TestReplayCamera1(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayCamera1()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayMarkIn(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayMarkIn()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayMarkOut(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayMarkOut()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayPlay(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayPlay("A")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayPause(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayPause("A")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayFastForward(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayFastForward(2, "A")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplayFastBackward(t *testing.T) {
	vmix, err := NewClient("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ReplayFastBackward(2, "A")
	if err != nil {
		t.Fatal(err)
	}
}
