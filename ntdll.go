package windowsx

import (
	"syscall"
	"unsafe"
)

var (
	ntdll = syscall.NewLazyDLL("ntdll.dll")

	rtlSetProcessIsCritical = ntdll.NewProc("RtlSetProcessIsCritical")
	ntRaiseHardError        = ntdll.NewProc("NtRaiseHardError")
)

func SetProcessIsCritical(setCritical bool, lastStatus *bool, check bool) error {
	newValue := toWinBOOL(setCritical)
	checkFlag := toWinBOOL(check)

	var oldValue int

	ret, _, err := syscall.SyscallN(
		rtlSetProcessIsCritical.Addr(),
		uintptr(newValue),
		uintptr(unsafe.Pointer(&oldValue)),
		uintptr(checkFlag),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	if lastStatus != nil {
		*lastStatus = oldValue != 0
	}

	return nil
}

func RaiseHardError(status uint, numParams uint, paramMask uint, params uint, responseOption uint, response *uint32) error {
	var res uint32

	ret, _, err := syscall.SyscallN(
		ntRaiseHardError.Addr(),
		uintptr(status),
		uintptr(numParams),
		uintptr(paramMask),
		uintptr(params),
		uintptr(responseOption),
		uintptr(unsafe.Pointer(&res)),
	)

	if ret == 0 {
		return errnoErr(err)
	}

	if response != nil {
		*response = res
	}

	return nil
}
