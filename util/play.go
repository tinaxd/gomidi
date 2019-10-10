package util

import (
	"syscall"
	"unsafe"
)

type HMIDIOUT struct {
	unused int32
}

func MidiOutOpen(hMidiOut *HMIDIOUT, devid uint) {
	hMidiOutPtr := unsafe.Pointer(&hMidiOut)
	var proc = winmmDll.findProc("midiOutOpen")

	_, _, lastErr := proc.Call(uintptr(nil), uintptr(devid), uintptr(0), uintptr(0), uintptr(0))
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		panic(lastErr)
	}
}

func MidiOutClose(hMidiOut *HMIDIOUT) {
	hMidiOutPtr := unsafe.Pointer(&hMidiOut)
	var proc = winmmDll.findProc("midiOutClose")

	_, _, lastErr := proc.Call(uintptr(hMidiOutPtr))
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		panic(lastErr)
	}
}

// MidiOutShortMsg sends msg to the midi device.
// msg is little-endian.
func MidiOutShortMsg(hMidiOut *HMIDIOUT, msg uint64) {
	hMidiOutPtr := unsafe.Pointer(&hMidiOut)
	var proc = winmmDll.findProc("midiOutShortMsg")

	_, _, lastErr := proc.Call(uintptr(hMidiOutPtr), uintptr(msg))
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		panic(lastErr)
	}
}
