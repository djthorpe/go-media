package sdl_test

import (
	"testing"

	// Packages
	"github.com/mutablelogic/go-media"
	ffmpeg "github.com/mutablelogic/go-media/pkg/ffmpeg"
	"github.com/mutablelogic/go-media/pkg/sdl"
	"github.com/stretchr/testify/assert"
)

func Test_video_001(t *testing.T) {
	assert := assert.New(t)

	// Initialize SDL with off-screen renderer
	err := sdl.New(media.VIDEO, nil, sdl.OptTitle(t.Name()), sdl.OptWindowSize(640, 480), sdl.OptOffscreenRenderer())
	if !assert.NoError(err) {
		t.SkipNow()
	}
}

func Test_video_002(t *testing.T) {
	assert := assert.New(t)

	main := func(sdl *sdl.Context) error {
		// Create a new video texture
		video, err := sdl.NewVideo(ffmpeg.VideoPar("yuv420p", "640x480", 25))
		if err != nil {
			return err
		}
		t.Log(video)
		return nil
	}

	// Initialize SDL with off-screen renderer
	err := sdl.New(media.VIDEO, main, sdl.OptTitle(t.Name()), sdl.OptWindowSize(640, 480), sdl.OptOffscreenRenderer())
	if !assert.NoError(err) {
		t.SkipNow()
	}
}
