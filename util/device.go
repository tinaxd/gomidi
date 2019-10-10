package util

import (
	"fmt"
	"syscall"
	"unsafe"
)

// MidiOutCaps describes capabilities of a midi output device
type MidiOutCaps struct {
	wMid           uint16
	wPid           uint16
	vDriverVersion uint32
	szPname        uint16
	wTechnology    uint16
	wVoices        uint16
	wNotes         uint16
	wChannelMask   uint16
	dwSupport      uint32
}

// GetMidiOutputDevNum returns the number of available midi output devices
func GetMidiOutputDevNum() uint {
	var proc = winmmDll.findProc("midiOutGetNumDevs")

	r1, _, lastErr := proc.Call()
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		panic(lastErr)
	}
	ret := (uint)(r1)
	return ret
}

func GetMidiOutputCapInfo(devid uint) MidiOutCaps {
	var proc = winmmDll.findProc("midiOutGetDevCapsW")

	devidPtr := (uintptr)(devid)
	retStruct := MidiOutCaps{}
	structPtr := unsafe.Pointer(&retStruct)

	r1, _, lastErr := proc.Call(devidPtr, uintptr(structPtr), uintptr(50))
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		panic(lastErr)
	}

	fmt.Printf("%v", r1)

	return retStruct
}
