package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	findFirstStreamW       = kernel32.NewProc("FindFirstStreamW")
	findNextStreamW        = kernel32.NewProc("FindNextStreamW")
	beginUpdateResourceW   = kernel32.NewProc("BeginUpdateResourceW")
	updateResourceW        = kernel32.NewProc("UpdateResourceW")
	endUpdateResourceW     = kernel32.NewProc("EndUpdateResourceW")
	enumResourceLanguagesW = kernel32.NewProc("EnumResourceLanguagesW")
)

func FindFirstStream(
	filePath *uint16,
	infoLevel uint32,
	streamName *string,
	flags uint32,
) (uintptr, error) {
	var streamData FindStreamData

	ret, _, err := syscall.SyscallN(
		findFirstStreamW.Addr(),
		uintptr(unsafe.Pointer(filePath)),
		uintptr(infoLevel),
		uintptr(unsafe.Pointer(&streamData)),
		uintptr(flags),
	)

	if ret == 0 {
		return ret, errnoErr(err)
	}

	if streamName != nil {
		*streamName = syscall.UTF16ToString(streamData.StreamName[:])
	}

	return ret, nil
}

func FindNextStream(handle uintptr, streamName *string) error {
	var streamData FindStreamData

	ret, _, err := syscall.SyscallN(
		findNextStreamW.Addr(),
		handle,
		uintptr(unsafe.Pointer(&streamData)),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	if streamName != nil {
		*streamName = syscall.UTF16ToString(streamData.StreamName[:])
	}

	return nil
}

func BeginUpdateResource(filePath *uint16, deleteExisting bool) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		beginUpdateResourceW.Addr(),
		uintptr(unsafe.Pointer(filePath)),
		toWinBOOL(deleteExisting),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func UpdateResource(handle uintptr, kind *uint16, name *uint16, language uint16, data []byte) error {
	var theData unsafe.Pointer
	if len(data) > 0 {
		theData = unsafe.Pointer(&data[0])
	}

	ret, _, err := syscall.SyscallN(
		updateResourceW.Addr(),
		handle,
		uintptr(unsafe.Pointer(kind)),
		uintptr(unsafe.Pointer(name)),
		uintptr(language),
		uintptr(theData),
		uintptr(len(data)),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func EndUpdateResource(handle uintptr, discard bool) error {
	ret, _, err := syscall.SyscallN(endUpdateResourceW.Addr(), handle, toWinBOOL(discard))
	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func EnumResourceLanguages(
	handle uintptr,
	kind uint32,
	name *uint16,
	callback uintptr,
	lParam uintptr,
) error {
	ret, _, err := syscall.SyscallN(
		enumResourceLanguagesW.Addr(),
		handle,
		uintptr(kind),
		uintptr(unsafe.Pointer(name)),
		callback,
		lParam,
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}
