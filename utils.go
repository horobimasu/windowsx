package windowsx

func toWinBOOL(theBool bool) uintptr {
	if theBool == true {
		return 1
	} else {
		return 0
	}
}
