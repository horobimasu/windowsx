package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	powrprof = syscall.NewLazyDLL("powrprof.dll")

	powerGetActiveScheme   = powrprof.NewProc("PowerGetActiveScheme")
	powerSetActiveScheme   = powrprof.NewProc("PowerSetActiveScheme")
	powerWriteACValueIndex = powrprof.NewProc("PowerWriteACValueIndex")
	powerWriteDCValueIndex = powrprof.NewProc("PowerWriteDCValueIndex")
	powerReadACValueIndex  = powrprof.NewProc("PowerReadACValueIndex")
	powerReadDCValueIndex  = powrprof.NewProc("PowerReadDCValueIndex")
)

func PowerGetActiveScheme(key uintptr, policy **syscall.GUID) error {
	ret, _, _ := syscall.SyscallN(powerGetActiveScheme.Addr(), key, uintptr(unsafe.Pointer(policy)))
	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}

func PowerSetActiveScheme(key uintptr, scheme *syscall.GUID) error {
	ret, _, _ := syscall.SyscallN(powerSetActiveScheme.Addr(), key, uintptr(unsafe.Pointer(scheme)))
	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}

func PowerWriteACValueIndex(
	key uintptr,
	scheme *syscall.GUID,
	group *syscall.GUID,
	setting *syscall.GUID,
	index uint32,
) error {
	ret, _, _ := syscall.SyscallN(
		powerWriteACValueIndex.Addr(),
		key,
		uintptr(unsafe.Pointer(scheme)),
		uintptr(unsafe.Pointer(group)),
		uintptr(unsafe.Pointer(setting)),
		uintptr(index),
	)

	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}

func PowerWriteDCValueIndex(
	key uintptr,
	scheme *syscall.GUID,
	group *syscall.GUID,
	setting *syscall.GUID,
	index uint32,
) error {
	ret, _, _ := syscall.SyscallN(
		powerWriteDCValueIndex.Addr(),
		key,
		uintptr(unsafe.Pointer(scheme)),
		uintptr(unsafe.Pointer(group)),
		uintptr(unsafe.Pointer(setting)),
		uintptr(index),
	)

	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}

func PowerReadACValueIndex(
	key uintptr,
	scheme *syscall.GUID,
	group *syscall.GUID,
	setting *syscall.GUID,
	index *uint32,
) error {
	ret, _, _ := syscall.SyscallN(
		powerReadACValueIndex.Addr(),
		key,
		uintptr(unsafe.Pointer(scheme)),
		uintptr(unsafe.Pointer(group)),
		uintptr(unsafe.Pointer(setting)),
		uintptr(unsafe.Pointer(index)),
	)

	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}

func PowerReadDCValueIndex(
	key uintptr,
	scheme *syscall.GUID,
	group *syscall.GUID,
	setting *syscall.GUID,
	index *uint32,
) error {
	ret, _, _ := syscall.SyscallN(
		powerReadDCValueIndex.Addr(),
		key,
		uintptr(unsafe.Pointer(scheme)),
		uintptr(unsafe.Pointer(group)),
		uintptr(unsafe.Pointer(setting)),
		uintptr(unsafe.Pointer(index)),
	)

	if ret != 0 {
		return errnoErr(syscall.Errno(ret))
	}

	return nil
}
