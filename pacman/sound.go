package pacman

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	pacsounds "github.com/kgosse/pacmanresources/sounds"
	"log"
)

type sounds struct {
	audioContext *audio.Context
	sirenPlayer  *audio.Player
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

	s.sirenPlayer.SetVolume(0.2)
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
	if !s.sirenPlayer.IsPlaying() {
		s.sirenPlayer.Rewind()
		s.sirenPlayer.Play()
	}
}
