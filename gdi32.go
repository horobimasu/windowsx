package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	createCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	createDIBSection   = gdi32.NewProc("CreateDIBSection")
	selectObject       = gdi32.NewProc("SelectObject")
	deleteObject       = gdi32.NewProc("DeleteObject")
	deleteDC           = gdi32.NewProc("DeleteDC")
)

func CreateCompatibleDc(handle uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(createCompatibleDC.Addr(), handle)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func CreateDibSection(
	handle uintptr,
	info *BitmapInfoHeader,
	usage uint32,
	bits unsafe.Pointer,
	section uintptr,
	offset uint32,
) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		createDIBSection.Addr(),
		handle,
		uintptr(unsafe.Pointer(info)),
		uintptr(usage),
		uintptr(unsafe.Pointer(bits)),
		section,
		uintptr(offset),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func SelectObject(handle uintptr, object uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(selectObject.Addr(), handle, object)
	if ret == 0 || ret == ^uintptr(0) { // HGDI_ERROR
		return 0, errnoErr(err)
	}

	return ret, nil
}

func DeleteObject(handle uintptr) error {
	ret, _, err := syscall.SyscallN(deleteObject.Addr(), handle)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func DeleteDc(handle uintptr) error {
	ret, _, err := syscall.SyscallN(deleteDC.Addr(), handle)
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}
