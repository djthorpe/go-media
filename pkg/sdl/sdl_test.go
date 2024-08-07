package sdl_test

import (
	"testing"

	// Packages
	"github.com/mutablelogic/go-media"
	"github.com/mutablelogic/go-media/pkg/ffmpeg"
	"github.com/mutablelogic/go-media/pkg/sdl"
	"github.com/stretchr/testify/assert"
)

func Test_audio_001(t *testing.T) {
	assert := assert.New(t)

	// Initialize SDL
	sdl, err := sdl.New(media.AUDIO)
	if !assert.NoError(err) {
		t.SkipNow()
	}
	t.Cleanup(func() {
		assert.NoError(sdl.Close())
	})
}

func Test_audio_002(t *testing.T) {
	assert := assert.New(t)

	// Initialize SDL
	sdl, err := sdl.New(media.AUDIO)
	if !assert.NoError(err) {
		t.SkipNow()
	}
	t.Cleanup(func() {
		assert.NoError(sdl.Close())
	})

	// Enumerate all audio devices
	devices := sdl.Devices(media.AUDIO)
	t.Log(devices)
}

func Test_audio_003(t *testing.T) {
	assert := assert.New(t)

	// Initialize SDL
	s, err := sdl.New(media.AUDIO)
	if !assert.NoError(err) {
		t.SkipNow()
	}
	t.Cleanup(func() {
		assert.NoError(s.Close())
	})

	// Capture from default audio device
	audio, err := s.NewAudio(ffmpeg.AudioPar("flt", "mono", 16000), sdl.DefaultAudioSink)
	if !assert.NoError(err) {
		t.SkipNow()
	}

	t.Log(audio)
}
