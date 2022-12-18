package zutils

import (
	"runtime"
	"unsafe"
)

func Malloc31(size int) (ret unsafe.Pointer) {
	ret = unsafe.Pointer(runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x7fd<<4,
		[]uintptr{uintptr(size)}))
	return
}
func Free(ptr unsafe.Pointer) {
	runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x059<<4,
		[]uintptr{uintptr(ptr)})
}

func EtoA(record []byte) {
	sz := len(record)
	runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x6e3<<4, // __e2a_l
		[]uintptr{uintptr(unsafe.Pointer(&record[0])), uintptr(sz)})
}

func AtoE(record []byte) {
	sz := len(record)
	runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x741<<4, // __a2e_l
		[]uintptr{uintptr(unsafe.Pointer(&record[0])), uintptr(sz)})
}

//go:noescape
func Bpxcall(plist []unsafe.Pointer, bpx_offset int64)

//go:nosplit
func IefssreqX(parm unsafe.Pointer, branch_ptr unsafe.Pointer, save_area unsafe.Pointer) uintptr

func Iefssreq(parm unsafe.Pointer, dsa unsafe.Pointer) (ret uintptr) {
	branch_ptr := unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(0) + 16))) + 296))) + 20))))
	ret = IefssreqX(parm, branch_ptr, dsa)
	return
}
