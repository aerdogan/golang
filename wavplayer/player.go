package main

import (
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	mmsystem      = syscall.MustLoadDLL("winmm.dll")
	sndPlaySoundW = mmsystem.MustFindProc("sndPlaySoundW")
)

func main() {
	file, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(err)
	}

	s16, _ := syscall.UTF16PtrFromString(file)
	sndPlaySoundW.Call(uintptr(unsafe.Pointer(s16)), uintptr(0x0000))
}
