package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var sounds sound

type sound struct {
	Ctrl      *beep.Ctrl
	Ambiance  beep.StreamSeekCloser
	Stone     beep.StreamSeekCloser
	Capture   beep.StreamSeekCloser
	WinIA     beep.StreamSeekCloser
	WinPlayer beep.StreamSeekCloser
	Winvsia   beep.StreamSeekCloser
}

func loadAudio() {
	file, err := os.Open("sounds/stone.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ := mp3.Decode(file)
	sounds.Stone = s

	file, err = os.Open("sounds/capture.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ = mp3.Decode(file)
	sounds.Capture = s

	file, err = os.Open("sounds/ambiance.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ = mp3.Decode(file)
	sounds.Ambiance = s

	file, err = os.Open("sounds/winia.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ = mp3.Decode(file)
	sounds.WinIA = s

	file, err = os.Open("sounds/winplayer.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ = mp3.Decode(file)
	sounds.WinPlayer = s

	file, err = os.Open("sounds/winvsia.mp3")
	if err != nil {
		log.Fatal(err)
	}
	s, format, _ = mp3.Decode(file)
	sounds.Winvsia = s
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/40))
}

func playAmbiance() {
	sounds.Ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, sounds.Ambiance)}
	playing := make(chan struct{})
	speaker.Play(beep.Seq(sounds.Ctrl, beep.Callback(func() {
		close(playing)
	})))
	<-playing
	sounds.Ambiance.Seek(0)
}

func playSound(sound beep.StreamSeekCloser) {
	if sound.Position() != 0 {
		return
	}
	sound.Seek(0)
	playing := make(chan struct{})
	speaker.Play(beep.Seq(sound, beep.Callback(func() {
		close(playing)
	})))
	<-playing
	sound.Seek(0)
}
