package windowsx

import (
	"syscall"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	setCursorPos = user32.NewProc("SetCursorPos")
	blockInput   = user32.NewProc("BlockInput")
)

func SetCursorPosition(x int, y int) error {
	ret, _, err := syscall.SyscallN(setCursorPos.Addr(), uintptr(x), uintptr(y))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func BlockInput(block bool) error {
	blockIt := toWinBOOL(block)

	ret, _, err := syscall.SyscallN(blockInput.Addr(), uintptr(blockIt))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}
