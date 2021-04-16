package models

type Input struct {
	// Common properties
	Name         string `xml:",chardata"`
	Key          string `xml:"key,attr"`
	Number       uint   `xml:"number,attr"`
	SceneType    string `xml:"type,attr"`
	Title        string `xml:"title,attr"` // same as Name??
	ShortTitle   string `xml:"shorttite,attr"`
	State        string `xml:"state,attr"` // Paused | Running
	AttrPosition int    `xml:"position,attr"`
	Duration     int    `xml:"duration,attr"`
	Loop         bool   `xml:"loop,attr"`

	// Sound related
	Muted       bool    `xml:"muted,attr"`
	Volume      float64 `xml:"volume,attr"`
	Balance     float64 `xml:"balance,attr"`
	Solo        bool    `xml:"solo,attr"`
	AudioBusses string  `xml:"audiobusses,attr"`
	MeterF1     float64 `xml:"meterF1,attr"`
	MeterF2     float64 `xml:"meterF2,attr"`
	GainDb      string  `xml:"gainDb,attr"`
	Position    struct {
		Text  string `xml:",chardata"`
		PanX  string `xml:"panX,attr"`
		PanY  string `xml:"panY,attr"`
		ZoomX string `xml:"zoomX,attr"`
		ZoomY string `xml:"zoomY,attr"`
	} `xml:"position"`

	// vMix Instant Replay
	Replay struct {
		Text      string  `xml:",chardata"`
		Live      bool    `xml:"live,attr"`
		Recording bool    `xml:"recording,attr"`
		Events    int     `xml:"events,attr"`
		CameraA   string  `xml:"cameraA,attr"`
		CameraB   string  `xml:"cameraB,attr"`
		Speed     float64 `xml:"speed,attr"`
		Timecode  string  `xml:"timecode"` // time.Time - e.g. 2020-08-14T16:23:13.832
	} `xml:"replay"`

	// Multi view
	Overlay []struct {
		Text     string `xml:",chardata"`
		Index    int    `xml:"index,attr"`
		Key      string `xml:"key,attr"`
		Position struct {
			Text  string  `xml:",chardata"`
			PanX  float64 `xml:"panX,attr"`
			PanY  float64 `xml:"panY,attr"`
			ZoomX float64 `xml:"zoomX,attr"`
			ZoomY float64 `xml:"zoomY,attr"`
		} `xml:"position"`
	} `xml:"overlay"`
}

type Overlay struct {
	Number uint `xml:"number,attr"`
}

type Audio struct {
	Volume           float64 `xml:"volume,attr"`
	Muted            bool    `xml:"muted,attr"`
	MeterF1          float64 `xml:"meterF1,attr"`
	MeterF2          float64 `xml:"meterF2,attr"`
	HeadphonesVolume float64 `xml:"headphonesVolume,attr"`
}

type Transition struct {
	Number   uint   `xml:"number,attr"`
	Effect   string `xml:"effect,attr"`
	Duration uint   `xml:"duration,attr"`
}
