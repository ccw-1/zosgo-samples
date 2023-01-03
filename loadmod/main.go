package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
	"unsafe"
	"zosgo-samples/zutils"
)

type PLIST struct {
	list  [4]uint32
	parm1 [80]byte
	parm2 [80]byte
	parm3 [80]byte
	parm4 [80]byte
}

func DoPrint(rd *os.File, c chan int) {
	fmt.Printf("DoPrint\n")
	var buffer [4096]byte
	n, err := rd.Read(buffer[:])
	if err != nil {
		fmt.Printf("Read()  %v\n", err)
		return
	}
	for n > 0 {
		zutils.EtoA(buffer[:n])
		fmt.Print(string(buffer[:n]))
		n, err = rd.Read(buffer[:])
		if err != nil {
			fmt.Printf("Read()  %v\n", err)
			return
		}
	}
	c <- 55
}

func main() {
	// zutils.Trace = true
	ic := make(chan int)
	rd, wr, err := os.Pipe()
	if err != nil {
		fmt.Printf("Pipe() %v\n", err)
		os.Exit(1)
	}
	wrstr := fmt.Sprintf("%d", wr.Fd())
	runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x71b<<4, []uintptr{uintptr(unsafe.Pointer(&([]byte("_BPXK_JOBLOG"))[0])), uintptr(unsafe.Pointer(&([]byte(wrstr))[0])), 1})
	go DoPrint(rd, ic)

	siz := (int((reflect.TypeOf((*PLIST)(nil)).Elem()).Size()))
	plist := (*PLIST)(unsafe.Pointer(zutils.Malloc24(siz)))
	copy(plist.parm1[:], "AMODE 24 parameter 1")
	copy(plist.parm2[:], "AMODE 24 parameter 2")
	copy(plist.parm3[:], "AMODE 24 parameter 3")
	copy(plist.parm4[:], "AMODE 24 parameter 4")
	zutils.AtoE(plist.parm1[:])
	zutils.AtoE(plist.parm2[:])
	zutils.AtoE(plist.parm3[:])
	zutils.AtoE(plist.parm4[:])
	plist.list[0] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist.parm1[0])))
	plist.list[1] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist.parm2[0])))
	plist.list[2] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist.parm3[0])))
	plist.list[3] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist.parm4[0])))
	plist.list[3] |= uint32(0x80000000)

	mod := zutils.LoadMod("A24")
	fmt.Printf("Test 24-bit load module A24\n")
	if uintptr(unsafe.Pointer(mod)) != 0 {
		RC := mod.Call(uintptr(unsafe.Pointer(plist)))
		if RC != 0 {
			fmt.Printf("RC=0x%x\n", RC)
		}
		mod.Free()
	} else {
		fmt.Printf("Failed to load module A24\n")
	}
	zutils.Free(unsafe.Pointer(plist))

	siz = (int((reflect.TypeOf((*PLIST)(nil)).Elem()).Size()))
	plist31 := (*PLIST)(unsafe.Pointer(zutils.Malloc24(siz)))
	copy(plist31.parm1[:], "AMODE 31 parameter 1")
	copy(plist31.parm2[:], "AMODE 31 parameter 2")
	copy(plist31.parm3[:], "AMODE 31 parameter 3")
	copy(plist31.parm4[:], "AMODE 31 parameter 4")
	zutils.AtoE(plist31.parm1[:])
	zutils.AtoE(plist31.parm2[:])
	zutils.AtoE(plist31.parm3[:])
	zutils.AtoE(plist31.parm4[:])
	plist31.list[0] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist31.parm1[0])))
	plist31.list[1] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist31.parm2[0])))
	plist31.list[2] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist31.parm3[0])))
	plist31.list[3] = uint32(0x0ffffffff & uintptr(unsafe.Pointer(&plist31.parm4[0])))
	plist31.list[3] |= uint32(0x80000000)
	mod2 := zutils.LoadMod("A31")
	fmt.Printf("Test 31-bit load module A31\n")
	if uintptr(unsafe.Pointer(mod2)) != 0 {
		RC := mod2.Call(uintptr(unsafe.Pointer(plist31)))
		if RC != 0 {
			fmt.Printf("RC=0x%x\n", RC)
		}
		mod2.Free()
	} else {
		fmt.Printf("Failed to load module A31\n")
	}
	zutils.Free(unsafe.Pointer(plist31))

	msg := []byte("========================= My AMODE64 Argument String ==============================")
	zutils.AtoE(msg[:])
	var msglen uint32 = uint32(len(msg))

	var parm [2]unsafe.Pointer
	parm[0] = unsafe.Pointer(&msg[0])
	parm[1] = unsafe.Pointer(&msglen)

	fmt.Printf("Test 64-bit load module A64\n")
	mod4 := zutils.LoadMod("A64")
	if uintptr(unsafe.Pointer(mod4)) != 0 {
		RC := mod4.Call(uintptr(unsafe.Pointer(&parm[0])))
		if RC != 0 {
			fmt.Printf("RC=0x%x\n", RC)
		}
		mod4.Free()
	} else {
		fmt.Printf("Failed to load module A64\n")
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	select {
	case <-ic:
	case <-timeout:
	}
}
