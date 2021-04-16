package vmixhttp

import (
	"testing"
)

// General vMix features
// vmix.go or models.go
func TestRefresh(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}

	if err := vmix.Refresh(); err != nil {
		t.Fatal(err)
	}
}

//browser.go
func TestBrowserBack(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserBack(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserForward(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserForward(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserKeyboardDisabled(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserKeyboardDisabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserKeyboardEnabled(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserKeyboardEnabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserMouseDisabled(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserMouseDisabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserMouseEnabled(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserMouseEnabled(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserNavigate(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.BrowserNavigate(0, "http://google.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserReload(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
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
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOff("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceAutoNextOn(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOn("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceAutoNextOnOff(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceAutoNextOnOff("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceNextRow(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourceNextRow("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourcePreviousRow(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.DataSourcePreviousRow("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataSourceSelectRow(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
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
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ActivatorRefresh()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallManagerShowHide(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.CallManagerShowHide()
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeyPress(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.KeyPress("A")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendKeys(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
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
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDICommand(0, "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDISelectSourceByIndex(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDISelectSourceByIndex(0, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDISelectSourceByName(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDISelectSourceByName(0, "NDI")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDIStartRecording(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDIStartRecording(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNDIStopRecording(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.NDIStopRecording(0)
	if err != nil {
		t.Fatal(err)
	}
}

// output.go
func TestFullscreen(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Fullscreen()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFullscreenOff(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.FullscreenOff()
	if err != nil {
		t.Fatal(err)
	}
}
func TestFullscreenOn(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.FullscreenOn()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutput2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutput2(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutput3(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutput3(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutput4(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutput4(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutputExternal2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutputExternal2(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutputFullscreen(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutputFullscreen(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetOutputFullscreen2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetOutputFullscreen2(0, "Output")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSnapshot(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Snapshot("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSnapshotInput(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SnapshotInput(0, "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartExternal(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartExternal()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartMultiCorder(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartMultiCorder()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartRecording(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartRecording()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStopExternal(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartStopExternal()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStopMultiCorder(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartStopMultiCorder()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStopRecording(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartStopRecording()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStopStreaming(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartStopStreaming(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartStreaming(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StartStreaming(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopExternal(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StopExternal()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopMultiCorder(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StopMultiCorder()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopRecording(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StopRecording()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopStreaming(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StopStreaming(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamingSetKey(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StreamingSetKey("KEY")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamingSetPassword(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StreamingSetPassword("Pass")
	if err != nil {
		t.Fatal(err)
	}
}
func TestStreamingSetURL(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StreamingSetURL("rtmp://example.com")
	if err != nil {
		t.Fatal(err)
	}
}
func TestStreamingSetUsername(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.StreamingSetUsername("Username")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteDurationToRecordingLog(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.WriteDurationToRecordingLog("")
	if err != nil {
		t.Fatal(err)
	}
}

// preset.go
func TestLastPreset(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.LastPreset()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenPreset(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.OpenPreset("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSavePreset(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SavePreset("")
	if err != nil {
		t.Fatal(err)
	}
}

// script.go
func TestScriptStart(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ScriptStart("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestScriptStartDynamic(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ScriptStartDynamic(`Function=Cut
Sleep 1000
Function=OverlayInput1Out
`)
	if err != nil {
		t.Fatal(err)
	}
}

func TestScriptStop(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ScriptStop("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestScriptStopAll(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ScriptStopAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestScriptStopDynamic(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.ScriptStopDynamic()
	if err != nil {
		t.Fatal(err)
	}
}

// transition.go
func TestFade(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Fade(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestZoom(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Zoom(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWipe(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Wipe(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSlide(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Slide(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFly(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Fly(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrossZoom(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.CrossZoom(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFlyRotate(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.FlyRotate(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCube(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Cube(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCubeZoom(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.CubeZoom(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerticalWipe(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VerticalWipe(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerticalSlide(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VerticalSlide(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMerge(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Merge(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWipeReverse(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.WipeReverse(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSlideReverse(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SlideReverse(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerticalWipeReverse(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VerticalWipeReverse(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerticalSlideReverse(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.VerticalSlideReverse(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCut(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Cut(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCutDirect(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.CutDirect(nil, 500)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFadeToBlack(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.FadeToBlack()
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuickPlay(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.QuickPlay(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetFader(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetFader(255)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionDuration1(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionDuration1(300)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionDuration2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionDuration2(300)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionDuration3(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionDuration3(300)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionDuration4(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionDuration4(300)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionEffect1(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionEffect1("Cut")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionEffect2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionEffect2("Fade")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionEffect3(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionEffect3("Merge")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTransitionEffect4(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.SetTransitionEffect4("Wipe")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStinger1(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Stinger1(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStinger2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Stinger2(0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransition1(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Transition1()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransition2(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Transition2()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransition3(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Transition3()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransition4(t *testing.T) {
	vmix, err := NewVmixHTTP("localhost", 8088)
	if err != nil {
		t.Fatal(err)
	}
	err = vmix.Transition4()
	if err != nil {
		t.Fatal(err)
	}
}
