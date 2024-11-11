package models

import (
	"encoding/xml"
	"time"

	"golang.org/x/xerrors"
)

type VmixBool bool

func (b *VmixBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	switch v {
	case "True":
		*b = true
	case "False":
		*b = false
	default:
		return xerrors.Errorf("invalid bool value: %s", v)
	}
	return nil
}

// APIXML vMix XML API Response
type APIXML struct {
	XMLName     xml.Name    `xml:"vmix"`
	Text        string      `xml:",chardata"`
	Version     string      `xml:"version"`
	Preset      string      `xml:"preset"`
	Edition     string      `xml:"edition"`
	Inputs      Inputs      `xml:"inputs"`
	Overlays    Overlays    `xml:"overlays"`
	Preview     int         `xml:"preview"`
	Active      int         `xml:"active"`
	FadeToBlack VmixBool    `xml:"fadeToBlack"`
	Transitions Transitions `xml:"transitions"`
	Recording   VmixBool    `xml:"recording"`
	External    VmixBool    `xml:"external"`
	Streaming   VmixBool    `xml:"streaming"`
	PlayList    string      `xml:"playList"`
	MultiCorder string      `xml:"multiCorder"`
	Fullscreen  VmixBool    `xml:"fullscreen"`
	Mix         []Mix       `xml:"mix"`
	Audio       Audios      `xml:"audio"`
	Dynamic     Dynaminc    `xml:"dynamic"`
}

type Inputs struct {
	Text  string  `xml:",chardata"`
	Input []Input `xml:"input"`
}

type Input struct {
	// Common properties
	Key        string   `xml:"key,attr"`
	Number     uint     `xml:"number,attr"`
	Type       string   `xml:"type,attr"`
	Title      string   `xml:"title,attr"` // same as Name??
	ShortTitle string   `xml:"shorttite,attr"`
	State      string   `xml:"state,attr"` // Paused | Running
	Position   int      `xml:"position,attr"`
	Duration   int      `xml:"duration,attr"`
	Loop       VmixBool `xml:"loop,attr"`
	Name       string   `xml:",chardata"`

	// Sound related
	Muted       VmixBool `xml:"muted,attr"`
	Volume      float64  `xml:"volume,attr"`
	Balance     float64  `xml:"balance,attr"`
	Solo        VmixBool `xml:"soloPFL,attr"`
	AudioBusses string   `xml:"audiobusses,attr"` // Comma separated list of busses
	MeterF1     float64  `xml:"meterF1,attr"`
	MeterF2     float64  `xml:"meterF2,attr"`
	GainDb      float64  `xml:"gainDb,attr"`

	// vMix Instant Replay
	// if type == "Replay" or "ReplayPreview"
	Replay InputReplay `xml:"replay"`

	// Layers
	Overlay []InputOverlay `xml:"overlay"`
}

type InputReplay struct {
	Text        string              `xml:",chardata"`
	Live        VmixBool            `xml:"live,attr"`
	Recording   VmixBool            `xml:"recording,attr"`
	ChannelMode string              `xml:"channelMode,attr"` // "A" | "B" | "AB"
	Events      int                 `xml:"events,attr"`
	EventsA     int                 `xml:"eventsA,attr"`
	EventsB     int                 `xml:"eventsB,attr"`
	CameraA     string              `xml:"cameraA,attr"`
	CameraB     string              `xml:"cameraB,attr"`
	Speed       float64             `xml:"speed,attr"`
	SpeedA      float64             `xml:"speedA,attr"`
	SpeedB      float64             `xml:"speedB,attr"`
	Timecode    InputReplayTimecode `xml:"timecode"`
	TimecodeA   InputReplayTimecode `xml:"timecodeA"`
	TimecodeB   InputReplayTimecode `xml:"timecodeB"`
}

// InputReplayTimecode - e.g. 2020-08-14T16:23:13.832
type InputReplayTimecode time.Time

func (t *InputReplayTimecode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	tt, err := time.Parse("2006-01-02T15:04:05.000", v)
	if err != nil {
		return xerrors.Errorf("failed to parse time: %w", err)
	}
	*t = InputReplayTimecode(tt)
	return nil
}

type InputOverlay struct {
	Text     string               `xml:",chardata"`
	Index    int                  `xml:"index,attr"`
	Key      string               `xml:"key,attr"`
	Position InputOverlayPosition `xml:"position"`
}

type InputOverlayPosition struct {
	Text  string  `xml:",chardata"`
	PanX  float64 `xml:"panX,attr"`
	PanY  float64 `xml:"panY,attr"`
	ZoomX float64 `xml:"zoomX,attr"`
	ZoomY float64 `xml:"zoomY,attr"`
}

type Overlays struct {
	Text    string    `xml:",chardata"`
	Overlay []Overlay `xml:"overlay"`
}

type Overlay struct {
	Input  int  `xml:",chardata"`
	Number uint `xml:"number,attr"`
}

type Audios struct {
	Text   string `xml:",chardata"`
	Master Audio  `xml:"master"`
	BusA   Audio  `xml:"busA"`
	BusB   Audio  `xml:"busB"`
	BusC   Audio  `xml:"busC"`
	BusD   Audio  `xml:"busD"`
	BusE   Audio  `xml:"busE"`
	BusF   Audio  `xml:"busF"`
	BusG   Audio  `xml:"busG"`
}

type Audio struct {
	Text             string   `xml:",chardata"`
	Volume           float64  `xml:"volume,attr"`
	Muted            VmixBool `xml:"muted,attr"`
	MeterF1          float64  `xml:"meterF1,attr"`
	MeterF2          float64  `xml:"meterF2,attr"`
	HeadphonesVolume float64  `xml:"headphonesVolume,attr"`
}

type Transitions struct {
	Text       string       `xml:",chardata"`
	Transition []Transition `xml:"transition"`
}

type Transition struct {
	Number   uint   `xml:"number,attr"`
	Effect   string `xml:"effect,attr"`
	Duration uint   `xml:"duration,attr"`
}

type Mix struct {
	Text    string `xml:",chardata"`
	Number  string `xml:"number,attr"`
	Preview int    `xml:"preview"`
	Active  int    `xml:"active"`
}

type Dynaminc struct {
	Text   string `xml:",chardata"`
	Input1 string `xml:"input1"`
	Input2 string `xml:"input2"`
	Input3 string `xml:"input3"`
	Input4 string `xml:"input4"`
	Value1 string `xml:"value1"`
	Value2 string `xml:"value2"`
	Value3 string `xml:"value3"`
	Value4 string `xml:"value4"`
}
