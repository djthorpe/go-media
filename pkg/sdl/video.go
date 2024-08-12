package sdl

import (
	"errors"

	// Packages
	ffmpeg "github.com/mutablelogic/go-media/pkg/ffmpeg"
	ff "github.com/mutablelogic/go-media/sys/ffmpeg61"
	sdl "github.com/veandco/go-sdl2/sdl"

	// Namespace imports
	. "github.com/djthorpe/go-errors"
	. "github.com/mutablelogic/go-media"
)

////////////////////////////////////////////////////////////////////////////////
// CGO

/*
extern void audio_callback(void* userdata, void* data, int len);
*/
import "C"

//////////////////////////////////////////////////////////////////////////////
// TYPES

type Video struct {
	texture *sdl.Texture
}

//////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	mapVideo = map[ff.AVPixelFormat]uint32{
		ff.AV_PIX_FMT_YUV420P: sdl.PIXELFORMAT_IYUV,
	}
)

/*
static const TextureFormatEntry sdl_texture_format_map[] = {
		"yuv420p":               sdl.PIXELFORMAT_IYUV,   //
		"rgb8":                  sdl.PIXELFORMAT_RGB332, // AV_PIX_FMT_RGB8
		"rgb24":                 sdl.PIXELFORMAT_RGB24,  // AV_PIX_FMT_RGB24
		"bgr24":                 sdl.PIXELFORMAT_BGR24,  // AV_PIX_FMT_BGR24
		"yuyv422":               sdl.PIXELFORMAT_YUY2,   // AV_PIX_FMT_YUYV422
		"uyvy422":               sdl.PIXELFORMAT_UYVY,   // AV_PIX_FMT_UYVY422

	{ AV_PIX_FMT_RGB444,		 SDL_PIXELFORMAT_RGB444 },
	{ AV_PIX_FMT_RGB555,		 SDL_PIXELFORMAT_RGB555 },
	{ AV_PIX_FMT_BGR555,		 SDL_PIXELFORMAT_BGR555 },
	{ AV_PIX_FMT_RGB565,		 SDL_PIXELFORMAT_RGB565 },
	{ AV_PIX_FMT_BGR565,		 SDL_PIXELFORMAT_BGR565 },
	{ AV_PIX_FMT_BGR24,		  SDL_PIXELFORMAT_BGR24 },
	{ AV_PIX_FMT_0RGB32,		 SDL_PIXELFORMAT_RGB888 },
	{ AV_PIX_FMT_0BGR32,		 SDL_PIXELFORMAT_BGR888 },
	{ AV_PIX_FMT_NE(RGB0, 0BGR), SDL_PIXELFORMAT_RGBX8888 },
	{ AV_PIX_FMT_NE(BGR0, 0RGB), SDL_PIXELFORMAT_BGRX8888 },
	{ AV_PIX_FMT_RGB32,		  SDL_PIXELFORMAT_ARGB8888 },
	{ AV_PIX_FMT_RGB32_1,		SDL_PIXELFORMAT_RGBA8888 },
	{ AV_PIX_FMT_BGR32,		  SDL_PIXELFORMAT_ABGR8888 },
	{ AV_PIX_FMT_BGR32_1,		SDL_PIXELFORMAT_BGRA8888 },
};

	{ AV_PIX_FMT_RGB24,		  SDL_PIXELFORMAT_RGB24 },
	{ AV_PIX_FMT_RGB8,		   SDL_PIXELFORMAT_RGB332 },
	{ AV_PIX_FMT_YUV420P,		SDL_PIXELFORMAT_IYUV },

*/

//////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewVideo creates a new video texture device with the specified parameters.
func (s *Context) NewVideo(par *ffmpeg.Par) (*Video, error) {
	if par == nil || !par.Type().Is(VIDEO) {
		return nil, ErrBadParameter.With("invalid video parameters")
	}
	if s.renderer == nil {
		return nil, ErrInternalAppError.With("renderer")
	}

	// Create the texture for the video
	var video Video
	if pixfmt, exists := mapVideo[par.PixelFormat()]; !exists {
		return nil, ErrBadParameter.Withf("pixel format %v", par.PixelFormat())
	} else if surface, err := s.renderer.CreateTexture(pixfmt, sdl.TEXTUREACCESS_STREAMING, int32(par.Width()), int32(par.Height())); err != nil {
		return nil, err
	} else {
		video.texture = surface
	}

	// Add to list of videos
	s.video = append(s.video, &video)

	// Return success
	return &video, nil
}

// Release texture resources
func (v *Video) Close() error {
	var result error

	// Destroy the surface
	result = errors.Join(result, v.texture.Destroy())

	// Return any errors
	return result
}

//////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

/*
func (w *Window) Flush() error {
	if err := w.Renderer.Copy(w.Texture, nil, nil); err != nil {
		return err
	}
	w.Renderer.Present()
	return nil
}

func (w *Window) RenderFrame(frame *ffmpeg.Frame) error {
	if err := w.UpdateYUV(
		nil,
		frame.Bytes(0),
		frame.Stride(0),
		frame.Bytes(1),
		frame.Stride(1),
		frame.Bytes(2),
		frame.Stride(2),
	); err != nil {
		return err
	}

	// Return success
	return nil
}
*/
