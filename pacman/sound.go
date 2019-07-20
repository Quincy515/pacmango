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

	s.sirenPlayer.SetVolume(0.2)
	s.eatFruitPlayer.SetVolume(0.1)
	s.wailPlayer.SetVolume(0.05)
	s.eatGhostPlayer.SetVolume(0.05)
	s.deathPlayer.SetVolume(0.05)
	return s
}

func (s *sounds) newPlayer(b []byte) *audio.Player {
	p, err := audio.NewPlayer(s.audioContext, s.load(b))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func (s *sounds) load(b []byte) *wav.Stream {
	stream, err := wav.Decode(s.audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func (s *sounds) playSiren() {
	if s.canPlaySiren() && !s.sirenPlayer.IsPlaying() {
		s.sirenPlayer.Rewind()
		s.sirenPlayer.Play()
	}
}

func (s *sounds) canPlaySiren() bool {
	d := [...]*audio.Player{
		s.wailPlayer,
		s.deathPlayer,
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
