package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	setCursorPos         = user32.NewProc("SetCursorPos")
	blockInput           = user32.NewProc("BlockInput")
	setWinEventHook      = user32.NewProc("SetWinEventHook")
	unhookWinEvent       = user32.NewProc("UnhookWinEvent")
	getClassNameW        = user32.NewProc("GetClassNameW")
	getMessageW          = user32.NewProc("GetMessageW")
	peekMessageW         = user32.NewProc("PeekMessageW")
	postThreadMessageW   = user32.NewProc("PostThreadMessageW")
	getWindowTextW       = user32.NewProc("GetWindowTextW")
	getWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")
	findWindowW          = user32.NewProc("FindWindowW")
	setWindowTextW       = user32.NewProc("SetWindowTextW")
	moveWindow           = user32.NewProc("MoveWindow")
	setWindowPos         = user32.NewProc("SetWindowPos")
	getSystemMetrics     = user32.NewProc("GetSystemMetrics")
	getWindowRect        = user32.NewProc("GetWindowRect")
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

func SetEventHook(
	eventMin uint32,
	eventMax uint32,
	hook uintptr,
	proc uintptr,
	procId uint32,
	threadId uint32,
	flags uint32,
) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		setWinEventHook.Addr(),
		uintptr(eventMin),
		uintptr(eventMax),
		hook,
		proc,
		uintptr(procId),
		uintptr(threadId),
		uintptr(flags),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return uintptr(ret), nil
}

func UnhookEvent(hook uintptr) error {
	ret, _, err := syscall.SyscallN(unhookWinEvent.Addr(), hook)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetClassName(handle uintptr) (string, error) {
	buffer := make([]uint16, 256)

	ret, _, err := syscall.SyscallN(
		getClassNameW.Addr(),
		handle,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
	)

	if ret == 0 {
		return "", errnoErr(err)
	}

	return syscall.UTF16ToString(buffer[:ret]), nil
}

func GetMessage(message *Msg, handle uintptr, filterMin uint32, filterMax uint32) (int32, error) {
	ret, _, err := syscall.SyscallN(
		getMessageW.Addr(),
		uintptr(unsafe.Pointer(message)),
		uintptr(handle),
		uintptr(filterMin),
		uintptr(filterMax),
	)

	retInt := int32(ret)
	if retInt == -1 {
		return retInt, errnoErr(err)
	}

	return retInt, nil
}

func PeekMessage(
	message *Msg,
	handle uintptr,
	filterMin uint32,
	filterMax uint32,
	removeMessage uint32,
) bool {
	ret, _, _ := syscall.SyscallN(
		peekMessageW.Addr(),
		uintptr(unsafe.Pointer(message)),
		uintptr(handle),
		uintptr(filterMin),
		uintptr(filterMax),
		uintptr(removeMessage),
	)

	return ret != 0
}

func PostThreadMessage(threadId uint32, message uint32, wParam uintptr, lParam uintptr) error {
	ret, _, err := syscall.SyscallN(
		postThreadMessageW.Addr(),
		uintptr(threadId),
		uintptr(message),
		wParam,
		lParam,
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetWindowText(handle uintptr) (string, error) {
	buffer := make([]uint16, 256)

	ret, _, err := syscall.SyscallN(
		getWindowTextW.Addr(),
		handle,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
	)

	if ret == 0 {
		return "", errnoErr(err)
	}

	return syscall.UTF16ToString(buffer[:ret]), nil
}

func GetWindowTextLength(handle uintptr) int32 {
	ret, _, _ := syscall.SyscallN(getWindowTextLengthW.Addr(), handle)
	return int32(ret)
}

func FindWindow(className *uint16, windowName *uint16) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		findWindowW.Addr(),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func SetWindowText(handle uintptr, text *uint16) error {
	ret, _, err := syscall.SyscallN(
		setWindowTextW.Addr(),
		handle,
		uintptr(unsafe.Pointer(text)),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func MoveWindow(handle uintptr, x int32, y int32, width int32, height int32, repaint bool) error {
	shouldRepain := toWinBOOL(repaint)

	ret, _, err := syscall.SyscallN(
		moveWindow.Addr(),
		handle,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(shouldRepain),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func SetWindowPos(
	handle uintptr,
	insertAfter uintptr,
	x int32,
	y int32,
	width int32,
	height int32,
	flags uint32,
) error {
	ret, _, err := syscall.SyscallN(
		setWindowPos.Addr(),
		handle,
		insertAfter,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetSystemMetrics(index uint32) int32 {
	ret, _, _ := syscall.SyscallN(getSystemMetrics.Addr(), uintptr(index))
	return int32(ret)
}

func GetWindowRect(handle uintptr) (Rect, error) {
	var rect Rect

	ret, _, err := syscall.SyscallN(
		getWindowRect.Addr(),
		handle,
		uintptr(unsafe.Pointer(&rect)),
	)

	if ret == 0 {
		return rect, errnoErr(err)
	}

	return rect, nil
}
