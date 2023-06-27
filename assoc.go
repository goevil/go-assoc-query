package assoc

import (
	"syscall"
	"unsafe"
)

var (
	shlwapiDll       = syscall.NewLazyDLL("shlwapi.dll")
	assocQueryString = shlwapiDll.NewProc("AssocQueryStringW")
)

func AssocQueryString(flags uint32, str uint32, pszAssoc *uint16, pszExtra *uint16, pszOut *uint16, pcchOut *uint32) error {
	ret, _, _ := assocQueryString.Call(
		uintptr(flags),
		uintptr(str),
		uintptr(unsafe.Pointer(pszAssoc)),
		uintptr(unsafe.Pointer(pszExtra)),
		uintptr(unsafe.Pointer(pszOut)),
		uintptr(unsafe.Pointer(pcchOut)),
	)
	if ret != 0 {
		return syscall.Errno(ret)
	}
	return nil
}
