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
