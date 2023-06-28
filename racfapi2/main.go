package main

import (
	"fmt"
	"unsafe"
	"zosgo-samples/zutils"
)

type Pkis64CplEntry struct {
	FieldName  [12]byte
	FieldLen   int32
	FieldValue [4096]byte
}
type Pkis64GencertMap struct {
	Eyecatch   [8]byte
	CertplLen  int32
	_          [4]byte
	CertplAddr unsafe.Pointer
	CertidAddr unsafe.Pointer
}

type PLIST struct {
	list         [14]unsafe.Pointer
	Workarea     [128]uint64
	SafRcAlet    int32
	SafRc        int32
	RacfRcAlet   int32
	RacfRc       int32
	RacfRsnAlet  int32
	RacfRsn      int32
	NumParms     uint64
	Func         uint16
	Attributes   uint32
	LogStringLen byte
	LogString    [255]byte
	ParmVer      uint32
	FuncParml    [4096]byte
	CaDomain     [4096]byte
}

const (
	Pkis64Gencert      = 0x0001
	Pkis64Export       = 0x0002
	Pkis64Queryreqs    = 0x0003
	Pkis64Reqdetails   = 0x0004
	Pkis64Modifyreqs   = 0x0005
	Pkis64Querycerts   = 0x0006
	Pkis64Certdetails  = 0x0007
	Pkis64Modifycerts  = 0x0008
	Pkis64Reqcert      = 0x0009
	Pkis64Verify       = 0x000A
	Pkis64Revoke       = 0x000B
	Pkis64Genrenew     = 0x000C
	Pkis64Reqrenew     = 0x000D
	Pkis64Respond      = 0x000E
	Pkis64Scepreq      = 0x000F
	Pkis64Preregister  = 0x0010
	Pkis64Qrecover     = 0x0011
	Pkis64Synch_create = 0x80000000
)

func main() {
	var plist64 PLIST
	plist64.list[0] = unsafe.Pointer(&plist64.Workarea)
	plist64.list[1] = unsafe.Pointer(&plist64.SafRcAlet)
	plist64.list[2] = unsafe.Pointer(&plist64.SafRc)
	plist64.list[3] = unsafe.Pointer(&plist64.RacfRcAlet)
	plist64.list[4] = unsafe.Pointer(&plist64.RacfRc)
	plist64.list[5] = unsafe.Pointer(&plist64.RacfRsnAlet)
	plist64.list[6] = unsafe.Pointer(&plist64.RacfRsn)
	plist64.list[7] = unsafe.Pointer(&plist64.NumParms)
	plist64.list[8] = unsafe.Pointer(&plist64.Func)
	plist64.list[9] = unsafe.Pointer(&plist64.Attributes)
	plist64.list[10] = unsafe.Pointer(&plist64.LogStringLen)
	plist64.list[11] = unsafe.Pointer(&plist64.ParmVer)
	plist64.list[12] = unsafe.Pointer(&plist64.FuncParml)
	plist64.list[13] = unsafe.Pointer(&plist64.CaDomain)
	plist64.Func = Pkis64Gencert
	plist64.Attributes = Pkis64Synch_create
	copy(plist64.LogString[:], "ZOSGO   ")
	zutils.AtoE(plist64.LogString[:])
	plist64.LogStringLen = 255

	// LOAD IRRSPX64
	mod := zutils.LoadMod("IRRSPX64")
	if uintptr(unsafe.Pointer(mod)) != 0 {
		if mod.Amode == zutils.Amode64 {
			fmt.Printf("IRRSPX64 is Amode64\n")
		} else {
			fmt.Printf("IRRSPX64 Amode number %v\n", mod.Amode)
		}
		RC := mod.Call(uintptr(unsafe.Pointer(&plist64)))
		if RC == 0 {
			fmt.Printf("SafRC %d RacfRc %d Reason %d\n", plist64.SafRc, plist64.RacfRc, plist64.RacfRsn)
		} else {
			fmt.Printf("Call rc=0x%x\n", RC)
		}
		mod.Free()
	} else {
		fmt.Printf("Failed to load IRRSPX64\n")
	}
	// FREE PARM STORAGE
}
