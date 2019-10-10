package util

import "syscall"

type midiLib struct {
	dll *syscall.DLL
}

var (
	winmmDll = loadDLL("winmm.dll")
)

func loadDLL(name string) *midiLib {
	dll, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	return &midiLib{dll}
}

func (c *midiLib) findProc(name string) *syscall.Proc {
	proc, err := c.dll.FindProc(name)
	if err != nil {
		panic(err)
	}
	return proc
}
