package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	setCursorPos             = user32.NewProc("SetCursorPos")
	getCursorPos             = user32.NewProc("GetCursorPos")
	blockInput               = user32.NewProc("BlockInput")
	setWinEventHook          = user32.NewProc("SetWinEventHook")
	unhookWinEvent           = user32.NewProc("UnhookWinEvent")
	getMessageW              = user32.NewProc("GetMessageW")
	peekMessageW             = user32.NewProc("PeekMessageW")
	postThreadMessageW       = user32.NewProc("PostThreadMessageW")
	getWindowTextW           = user32.NewProc("GetWindowTextW")
	getWindowTextLengthW     = user32.NewProc("GetWindowTextLengthW")
	findWindowW              = user32.NewProc("FindWindowW")
	setWindowTextW           = user32.NewProc("SetWindowTextW")
	moveWindow               = user32.NewProc("MoveWindow")
	setWindowPos             = user32.NewProc("SetWindowPos")
	getSystemMetrics         = user32.NewProc("GetSystemMetrics")
	getWindowRect            = user32.NewProc("GetWindowRect")
	openClipboard            = user32.NewProc("OpenClipboard")
	closeClipboard           = user32.NewProc("CloseClipboard")
	emptyClipboard           = user32.NewProc("EmptyClipboard")
	setClipboardData         = user32.NewProc("SetClipboardData")
	getClipboardData         = user32.NewProc("GetClipboardData")
	registerClipboardFormatW = user32.NewProc("RegisterClipboardFormatW")
	enumDisplaySettingsW     = user32.NewProc("EnumDisplaySettingsW")
	changeDisplaySettingsExW = user32.NewProc("ChangeDisplaySettingsExW")
	postMessageW             = user32.NewProc("PostMessageW")
	setProcessDPIAware       = user32.NewProc("SetProcessDPIAware")
	getCursorInfo            = user32.NewProc("GetCursorInfo")
	getIconInfo              = user32.NewProc("GetIconInfo")
	drawIconEx               = user32.NewProc("DrawIconEx")
	getWindowDisplayAffinity = user32.NewProc("GetWindowDisplayAffinity")
	setWindowDisplayAffinity = user32.NewProc("SetWindowDisplayAffinity")
	sendInput                = user32.NewProc("SendInput")
	registerClassExW         = user32.NewProc("RegisterClassExW")
	createWindowExW          = user32.NewProc("CreateWindowExW")
	defWindowProcW           = user32.NewProc("DefWindowProcW")
	showWindow               = user32.NewProc("ShowWindow")
	updateWindow             = user32.NewProc("UpdateWindow")
	destroyWindow            = user32.NewProc("DestroyWindow")
)

func SetCursorPosition(x int, y int) error {
	ret, _, err := syscall.SyscallN(setCursorPos.Addr(), uintptr(x), uintptr(y))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetCursorPos(point *Point) error {
	ret, _, err := syscall.SyscallN(getCursorPos.Addr(), uintptr(unsafe.Pointer(point)))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func BlockInput(block bool) error {
	ret, _, err := syscall.SyscallN(blockInput.Addr(), toWinBOOL(block))
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
	ret, _, err := syscall.SyscallN(
		moveWindow.Addr(),
		handle,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		toWinBOOL(repaint),
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

func OpenClipboard(handle uintptr) error {
	ret, _, err := syscall.SyscallN(openClipboard.Addr(), handle)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func CloseClipboard() error {
	ret, _, err := syscall.SyscallN(closeClipboard.Addr())
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func EmptyClipboard() error {
	ret, _, err := syscall.SyscallN(emptyClipboard.Addr())
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func SetClipboardData(format uint32, handle uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(setClipboardData.Addr(), uintptr(format), handle)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func GetClipboardData(format uint32) (uintptr, error) {
	ret, _, err := syscall.SyscallN(getClipboardData.Addr(), uintptr(format))
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func RegisterClipboardFormat(name *uint16) (uint32, error) {
	ret, _, err := syscall.SyscallN(registerClipboardFormatW.Addr(), uintptr(unsafe.Pointer(name)))
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return uint32(ret), nil
}

func EnumDisplaySettings(deviceName *uint16, mode uint32, device *DeviceMode) error {
	device.Size = uint16(unsafe.Sizeof(*device))
	device.DriverExtra = 0

	ret, _, err := syscall.SyscallN(
		enumDisplaySettingsW.Addr(),
		uintptr(unsafe.Pointer(deviceName)),
		uintptr(mode),
		uintptr(unsafe.Pointer(device)),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func ChangeDisplaySettingsEx(
	deviceName *uint16,
	device *DeviceMode,
	handle uintptr,
	flags uint32,
	lParam uintptr,
) int32 {
	ret, _, _ := syscall.SyscallN(
		changeDisplaySettingsExW.Addr(),
		uintptr(unsafe.Pointer(deviceName)),
		uintptr(unsafe.Pointer(device)),
		handle,
		uintptr(flags),
		lParam,
	)

	return int32(ret)
}

func PostMessage(handle uintptr, message uint32, wParam uintptr, lParam uintptr) error {
	ret, _, err := syscall.SyscallN(postMessageW.Addr(), handle, uintptr(message), wParam, lParam)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func SetProcessDPIAware() error {
	ret, _, err := syscall.SyscallN(setProcessDPIAware.Addr())
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetCursorInfo(info *CursorInfo) error {
	info.Size = uint32(unsafe.Sizeof(*info))

	ret, _, err := syscall.SyscallN(getCursorInfo.Addr(), uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetIconInfo(icon uintptr, info *IconInfo) error {
	ret, _, err := syscall.SyscallN(getIconInfo.Addr(), icon, uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func DrawIconEx(
	handle uintptr,
	left int32,
	top int32,
	icon uintptr,
	xWidth int32,
	yWidth int32,
	animated uint32,
	flickerFree uintptr,
	flags uint32,
) error {
	ret, _, err := syscall.SyscallN(
		drawIconEx.Addr(),
		handle,
		uintptr(left),
		uintptr(top),
		icon,
		uintptr(xWidth),
		uintptr(yWidth),
		uintptr(animated),
		flickerFree,
		uintptr(flags),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func GetWindowDisplayAffinity(handle uintptr) (uint32, error) {
	var affinity uint32

	ret, _, err := syscall.SyscallN(
		getWindowDisplayAffinity.Addr(),
		handle,
		uintptr(unsafe.Pointer(&affinity)),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return affinity, nil
}

func SetWindowDisplayAffinity(handle uintptr, affinity uint32) error {
	ret, _, err := syscall.SyscallN(setWindowDisplayAffinity.Addr(), handle, uintptr(affinity))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func SendInput(inputs []Input) (uint32, error) {
	ret, _, err := syscall.SyscallN(
		sendInput.Addr(),
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		unsafe.Sizeof(inputs[0]),
	)

	sent := uint32(ret)
	if sent == 0 {
		return 0, errnoErr(err)
	}

	return sent, nil
}

func RegisterClassEx(class *WindowClassEx) (uint16, error) {
	class.Size = uint32(unsafe.Sizeof(*class))

	ret, _, err := syscall.SyscallN(registerClassExW.Addr(), uintptr(unsafe.Pointer(class)))

	atom := uint16(ret)
	if atom == 0 {
		return 0, errnoErr(err)
	}

	return atom, nil
}

func CreateWindowEx(
	exStyle uint32,
	className *uint16,
	windowName *uint16,
	style uint32,
	x int32,
	y int32,
	width int32,
	height int32,
	parent uintptr,
	menu uintptr,
	instance uintptr,
	param uintptr,
) (uintptr, error) {
	ret, _, err := syscall.SyscallN(createWindowExW.Addr(),
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		parent,
		menu,
		instance,
		param,
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func DefWindowProc(handle uintptr, message uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := syscall.SyscallN(defWindowProcW.Addr(), handle, uintptr(message), wParam, lParam)
	return ret
}

func ShowWindow(handle uintptr, flag int) bool {
	ret, _, _ := syscall.SyscallN(showWindow.Addr(), handle, uintptr(flag))
	return ret != 0
}

func UpdateWindow(handle uintptr) error {
	ret, _, err := syscall.SyscallN(updateWindow.Addr(), handle)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func DestroyWindow(handle uintptr) error {
	ret, _, err := syscall.SyscallN(destroyWindow.Addr(), handle)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}
