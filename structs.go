package windowsx

import "unsafe"

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

type CursorInfo struct { // CURSORINFO
	Size      uint32
	Flags     uint32
	Cursor    uintptr
	ScreenPos Point
}

type IconInfo struct { // ICONINFO
	Icon     uint32
	XHotspot uint32
	YHotspot uint32
	Mask     uintptr
	Color    uintptr
}

type BitmapInfoHeader struct { // BITMAPINFOHEADER
	Size         uint32
	Width        int32
	Height       int32
	Planes       uint16
	BitCount     uint16
	Compression  uint32
	SizeImage    uint32
	XPels        int32
	YPels        int32
	ClrUsed      uint32
	ClrImportant uint32
}

type MouseInput struct { // MOUSEINPUT
	X         int32
	Y         int32
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type KeyboardInput struct { // KEYBDINPUT
	Code      uint16 // Vk
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type Input struct { // INPUT
	Type  uint32
	union [unsafe.Sizeof(MouseInput{}) / unsafe.Sizeof(uintptr(0))]uintptr
}

type WindowClassEx struct { // WNDCLASSEXA
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}
