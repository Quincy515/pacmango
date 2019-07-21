package pacman

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	pacsounds "github.com/kgosse/pacmanresources/sounds"
	"log"
)

type sounds struct {
	audioContext   *audio.Context
	sirenPlayer    *audio.Player
	eatFruitPlayer *audio.Player
	wailPlayer     *audio.Player
	eatGhostPlayer *audio.Player
	deathPlayer    *audio.Player
	entrancePlayer *audio.Player
	applausePlayer *audio.Player
	on             bool
}

const (
	sampleRate = 44100
)

func newSounds() *sounds {
	audioContext, err := audio.NewContext(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
	s := &sounds{
		audioContext: audioContext,
	}

	s.sirenPlayer = s.newPlayer(pacsounds.Siren_wav)
	s.eatFruitPlayer = s.newPlayer(pacsounds.EatFruit_wav)
	s.wailPlayer = s.newPlayer(pacsounds.Wail_wav)
	s.eatGhostPlayer = s.newPlayer(pacsounds.EatGhost_wav)
	s.deathPlayer = s.newPlayer(pacsounds.Death_wav)
	s.entrancePlayer = s.newPlayer(pacsounds.Beginning_wav)
	s.applausePlayer = s.newPlayer(pacsounds.Applause_wav)

	s.sirenPlayer.SetVolume(0.2)
	s.eatFruitPlayer.SetVolume(0.1)
	s.wailPlayer.SetVolume(0.05)
	s.eatGhostPlayer.SetVolume(0.05)
	s.deathPlayer.SetVolume(0.05)
	s.toggleSound()
	return s
}

func (s *sounds) newPlayer(b []byte) *audio.Player {
	p, err := audio.NewPlayer(s.audioContext, s.load(b))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func (s *sounds) toggleSound() {
	if s.on {
		s.on = false
		s.sirenPlayer.SetVolume(0)
		s.eatFruitPlayer.SetVolume(0)
		s.wailPlayer.SetVolume(0)
		s.eatGhostPlayer.SetVolume(0)
		s.deathPlayer.SetVolume(0)
		s.entrancePlayer.SetVolume(0)
		s.applausePlayer.SetVolume(0)
	} else {
		s.on = true
		s.sirenPlayer.SetVolume(0.2)
		s.eatGhostPlayer.SetVolume(0.1)
		s.wailPlayer.SetVolume(0.05)
		s.eatGhostPlayer.SetVolume(0.05)
		s.deathPlayer.SetVolume(0.05)
		s.entrancePlayer.SetVolume(0.2)
		s.applausePlayer.SetVolume(0.2)
	}
}

func (s *sounds) turnOff() {
	s.on = false
	s.sirenPlayer.SetVolume(0.2)
	s.eatGhostPlayer.SetVolume(0.1)
	s.wailPlayer.SetVolume(0.05)
	s.eatGhostPlayer.SetVolume(0.05)
	s.deathPlayer.SetVolume(0.05)
}

func (s *sounds) load(b []byte) *wav.Stream {
	stream, err := wav.Decode(s.audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func (s *sounds) status() string {
	var r string
	if s.on {
		r = "off"
	} else {
		r = "on"
	}
	return r
}

func (s *sounds) playSiren(won bool) {
	if !won && s.canPlaySiren() && !s.sirenPlayer.IsPlaying() {
		s.sirenPlayer.Rewind()
		s.sirenPlayer.Play()
	}
}

func (s *sounds) canPlaySiren() bool {
	d := [...]*audio.Player{
		s.wailPlayer,
		s.deathPlayer,
		s.entrancePlayer,
		s.applausePlayer,
	}

	for _, v := range d {
		if v.IsPlaying() {
			return false
		}
	}
	return true
}

func (s *sounds) playEatFruit() {
	if !s.eatFruitPlayer.IsPlaying() {
		s.eatFruitPlayer.Rewind()
		s.eatFruitPlayer.Play()
	}
}

func (s *sounds) playWail() {
	s.sirenPlayer.Pause()
	s.wailPlayer.Rewind()
	s.wailPlayer.Play()
}

func (s *sounds) playEatGhost() {
	if !s.eatGhostPlayer.IsPlaying() {
		s.eatGhostPlayer.Rewind()
		s.eatGhostPlayer.Play()
	}
}

func (s *sounds) playDeath() {
	s.sirenPlayer.Pause()
	if !s.sirenPlayer.IsPlaying() {
		s.deathPlayer.Rewind()
		s.deathPlayer.Play()
	}
}

func (s *sounds) pause() {
	d := [...]*audio.Player{
		s.wailPlayer,
		s.deathPlayer,
		s.sirenPlayer,
		s.eatGhostPlayer,
		s.eatFruitPlayer,
		s.entrancePlayer,
		s.applausePlayer,
	}

	for _, v := range d {
		v.Pause()
	}
}

func (s *sounds) playEntrance() {
	s.sirenPlayer.Pause()
	s.entrancePlayer.Rewind()
	s.entrancePlayer.Play()
}

func (s *sounds) playApplause() {
	s.sirenPlayer.Pause()
	if s.applausePlayer.IsPlaying() {
		return
	}
	s.applausePlayer.Rewind()
	s.applausePlayer.Play()
}
