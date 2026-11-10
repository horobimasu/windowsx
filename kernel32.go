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
	enumResourceNamesW     = kernel32.NewProc("EnumResourceNamesW")
	findResourceW          = kernel32.NewProc("FindResourceW")
	loadResource           = kernel32.NewProc("LoadResource")
	lockResource           = kernel32.NewProc("LockResource")
	sizeofResource         = kernel32.NewProc("SizeofResource")
	globalAlloc            = kernel32.NewProc("GlobalAlloc")
	globalFree             = kernel32.NewProc("GlobalFree")
	globalUnlock           = kernel32.NewProc("GlobalUnlock")
	globalLock             = kernel32.NewProc("GlobalLock")
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

func EnumResourceNames(handle uintptr, kind uint32, callback uintptr, lParam uintptr) error {
	ret, _, err := syscall.SyscallN(
		enumResourceNamesW.Addr(),
		handle,
		uintptr(kind),
		callback,
		lParam,
	)

	if ret == 0 {
		return errnoErr(err)
	}

	return nil
}

func FindResource(handle uintptr, name *uint16, kind uint32) (uintptr, error) {
	ret, _, err := syscall.SyscallN(
		findResourceW.Addr(),
		handle,
		uintptr(unsafe.Pointer(name)),
		uintptr(kind),
	)

	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func LoadResource(handle uintptr, info uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(loadResource.Addr(), handle, info)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func LockResource(resource uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(lockResource.Addr(), resource)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func SizeOfResource(handle uintptr, resource uintptr) (uint32, error) {
	ret, _, err := syscall.SyscallN(sizeofResource.Addr(), handle, resource)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return uint32(ret), nil
}

func GlobalAlloc(flags uint32, size uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(globalAlloc.Addr(), uintptr(flags), size)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func GlobalFree(flags uint32, size uintptr) (uintptr, error) {
	ret, _, err := syscall.SyscallN(globalFree.Addr(), uintptr(flags), size)
	if ret == 0 {
		return 0, errnoErr(err)
	}

	return ret, nil
}

func GlobalUnlock(handle uintptr) error {
	ret, _, err := syscall.SyscallN(globalUnlock.Addr(), handle)
	if ret == 0 && err != 0 {
		return errnoErr(err)
	}

	return nil
}

func GlobalLock(handle uintptr) (unsafe.Pointer, error) {
	ret, _, err := syscall.SyscallN(globalLock.Addr(), handle)

	ptr := unsafe.Pointer(ret)
	if ptr == nil {
		return nil, errnoErr(err)
	}

	return ptr, nil
}
