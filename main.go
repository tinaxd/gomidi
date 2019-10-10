package main

import (
	"time"

	"github.com/tinaxd/gowinmidi/util"
)

func main() {
	var midiH = &util.HMIDIOUT{}
	util.MidiOutOpen(midiH, 0)

	util.MidiOutShortMsg(midiH, 0x007f3c90)
	time.Sleep(1000 * time.Millisecond)
	util.MidiOutShortMsg(midiH, 0x007f3c80)

	util.MidiOutClose(midiH)
}
