package windowsx

type Point struct {
	X int32
	Y int32
}

type Msg struct {
	Handle  uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Point   Point
}

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type FindStreamData struct { // WIN32_FIND_STREAM_DATA
	StreamSize int64
	StreamName [260 + 36]uint16 // MAX_PATH
}

type DeviceMode struct { // DEVMODEW
	DeviceName         [32]uint16
	SpecVersion        uint16
	DriverVersion      uint16
	Size               uint16
	DriverExtra        uint16
	Fields             uint32
	Position           Point
	DisplayOrientation uint32
	DisplayFixedOutput uint32
	Color              int16
	Duplex             int16
	YResolution        int16
	TtOption           int16
	Collate            int16
	FormName           [32]uint16
	LogPixels          uint16
	BitsPerPel         uint32
	PelsWidth          uint32
	PelsHeight         uint32
	DisplayFlags       uint32
	DisplayFrequency   uint32
}
