package windowsx

func RGB(r uint8, g uint8, b uint8) uint32 {
	return uint32(r) | uint32(g)<<8 | uint32(b)<<16
}
