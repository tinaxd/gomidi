package util

/*
#cgo LDFLAGS: -lwinmm
#include <windows.h>
#include <mmsystem.h>

HMIDIOUT midi_out_open() {
	HMIDIOUT hMidiOut;

	midiOutOpen(&hMidiOut, MIDI_MAPPER, 0, 0, 0);
	return hMidiOut;
}

void midi_out_close(HMIDIOUT hMidiOut) {
	midiOutClose(hMidiOut);
}

void midi_out_short_msg(HMIDIOUT hMidiOut, DWORD msg) {
	midiOutShortMsg(hMidiOut, msg);
}

void midi_out_reset(HMIDIOUT hMidiOut) {
	midiOutReset(hMidiOut);
}

*/
import "C"

type HMIDIOUT struct {
	wrapped C.HMIDIOUT
}

func MidiOutOpen(devid uint) HMIDIOUT {
	r, err := C.midi_out_open()
	if err != nil {
		panic(err)
	}

	return HMIDIOUT{r}
}

func MidiOutClose(hmo *HMIDIOUT) {
	_, err := C.midi_out_close(hmo.wrapped)
	if err != nil {
		panic(err)
	}
}

func MidiOutReset(hmo *HMIDIOUT) {
	_, err := C.midi_out_reset(hmo.wrapped)
	if err != nil {
		panic(err)
	}
}

func MidiOutShortMsg(hmo *HMIDIOUT, msg int64) {
	_, err := C.midi_out_short_msg(hmo.wrapped, C.ulong(msg))
	if err != nil {
		panic(err)
	}
}
