package zutils

import (
	"runtime"
	"unsafe"
)

//go:nosplit
func A2e([]byte)

//go:nosplit
func E2a([]byte)

func Malloc31(size int) (ret unsafe.Pointer) {
	ret = unsafe.Pointer(runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x7fd<<4,
		[]uintptr{uintptr(size)}))
	return
}
func Free(ptr unsafe.Pointer) {
	runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x059<<4,
		[]uintptr{uintptr(ptr)})
}

//go:nosplit
func IefssreqX(parm unsafe.Pointer, branch_ptr unsafe.Pointer, save_area unsafe.Pointer) uintptr

func Iefssreq(parm unsafe.Pointer, dsa unsafe.Pointer) (ret uintptr) {
	branch_ptr := unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(0) + 16))) + 296))) + 20))))
	ret = IefssreqX(parm, branch_ptr, dsa)
	return
}
