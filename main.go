package main

import (
	"time"

	"github.com/tinaxd/gowinmidi/util"
)

func main() {

	midiH := util.MidiOutOpen(0)

	util.MidiOutShortMsg(&midiH, 0x007f3c90)
	time.Sleep(1000 * time.Millisecond)
	util.MidiOutShortMsg(&midiH, 0x007f3c80)

	util.MidiOutReset(&midiH)
	util.MidiOutClose(&midiH)
}
