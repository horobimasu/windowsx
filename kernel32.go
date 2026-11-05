package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	findFirstStreamW = kernel32.NewProc("FindFirstStreamW")
	findNextStreamW  = kernel32.NewProc("FindNextStreamW")
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
