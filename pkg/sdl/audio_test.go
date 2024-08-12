package sdl_test

import (
	"context"
	"errors"
	"testing"
	"time"

	// Packages
	"github.com/mutablelogic/go-media"
	"github.com/mutablelogic/go-media/pkg/ffmpeg"
	"github.com/mutablelogic/go-media/pkg/sdl"
	"github.com/stretchr/testify/assert"
)

func Test_audio_001(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(sdl.New(media.AUDIO, func(ctx *sdl.Context) error {
		// TODO
		return nil
	}))
}

func Test_audio_002(t *testing.T) {
	assert := assert.New(t)

	// Enumerate all audio devices
	assert.NoError(sdl.New(media.AUDIO, func(ctx *sdl.Context) error {
		devices := ctx.Devices(media.AUDIO)
		t.Log(devices)
		return nil
	}))
}

func Test_audio_003(t *testing.T) {
	assert := assert.New(t)

	// Output to default audio device
	assert.NoError(sdl.New(media.AUDIO, func(ctx *sdl.Context) error {
		audio, err := ctx.NewAudio(ffmpeg.AudioPar("flt", "mono", 16000), sdl.DefaultAudioSink, 0)
		if err != nil {
			return err
		}
		t.Log(audio)
		return nil
	}))

}

func Test_audio_004(t *testing.T) {
	assert := assert.New(t)

	// Capture from default audio device
	assert.NoError(sdl.New(media.AUDIO, func(ctx *sdl.Context) error {
		_, err := ctx.NewAudio(ffmpeg.AudioPar("flt", "mono", 16000), sdl.DefaultAudioSource, 20*time.Millisecond)
		if err != nil {
			return err
		}

		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Start capturing in runloop
		if err := ctx.Run(timeout); !errors.Is(err, context.DeadlineExceeded) {
			return err
		}

		// Return success
		return nil
	}, sdl.OptLog(true, t.Log), sdl.OptTitle("Test")))
}
