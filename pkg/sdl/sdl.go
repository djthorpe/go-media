package sdl

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"unsafe"

	// Packages
	sdl "github.com/veandco/go-sdl2/sdl"

	// Namespace imports
	. "github.com/mutablelogic/go-media"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Context struct {
	evt      map[uint32]func(unsafe.Pointer)
	audio    map[sdl.AudioDeviceID]*Audio
	video    []*Video
	window   *sdl.Window
	surface  *sdl.Surface
	renderer *sdl.Renderer
	logger   *logger
}

type logger struct {
	fn LogFn
}

var (
	pinner runtime.Pinner
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new SDL object which can output audio and video. Call the main
// function for the remainder of the program.
func New(t Type, main func(ctx *Context) error, opt ...Opt) error {
	var o opts
	var flags uint32
	if t.Is(VIDEO) {
		flags |= sdl.INIT_VIDEO
	}
	if t.Is(AUDIO) {
		flags |= sdl.INIT_AUDIO
	}

	// Set title
	if name, err := os.Executable(); err == nil {
		o.title = filepath.Base(name)
	} else {
		o.title = "SDL"
	}

	// Get options
	for _, opt := range opt {
		if err := opt(&o); err != nil {
			return err
		}
	}

	// Initialize SDL
	if err := sdl.Init(flags); err != nil {
		return err
	}

	// Create a context
	ctx := &Context{
		evt:   make(map[uint32]func(unsafe.Pointer)),
		audio: make(map[sdl.AudioDeviceID]*Audio),
	}
	pinner.Pin(ctx)

	// Set log function
	if o.logfn != nil {
		ctx.logger = &logger{o.logfn}
		if o.verbose {
			sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
		} else {
			sdl.LogSetAllPriority(sdl.LOG_PRIORITY_WARN)
		}
		sdl.LogSetOutputFunction(func(userData any, category int, pri sdl.LogPriority, message string) {
			ctx := userData.(*Context)
			pristr := "OTHER"
			switch pri {
			case sdl.LOG_PRIORITY_VERBOSE:
				pristr = "VERBOSE"
			case sdl.LOG_PRIORITY_DEBUG:
				pristr = "DEBUG"
			case sdl.LOG_PRIORITY_INFO:
				pristr = "INFO"
			case sdl.LOG_PRIORITY_WARN:
				pristr = "WARN"
			case sdl.LOG_PRIORITY_ERROR:
				pristr = "ERROR"
			case sdl.LOG_PRIORITY_CRITICAL:
				pristr = "CRITICAL"
			}
			ctx.logger.fn(fmt.Sprintf("[%s] %s", pristr, message))
		}, ctx)
	}

	var result error
	sdl.Main(func() {
		switch {
		case t.Is(VIDEO) && o.surface:
			if surface, renderer, err := ctx.createOffscreenRenderer(int32(o.width), int32(o.height)); err != nil {
				result = err
			} else {
				ctx.surface = surface
				ctx.renderer = renderer
			}
		case t.Is(VIDEO):
			if window, renderer, err := ctx.createWindowRenderer(o.title, int32(o.width), int32(o.height)); err != nil {
				result = err
			} else {
				ctx.window = window
				ctx.renderer = renderer
			}
		}

		// Run the main function
		if result == nil && main != nil {
			result = main(ctx)
		}

		// Close the context
		result = errors.Join(result, ctx.close())
	})

	// Return any errors
	return result
}

// Close the SDL object and release all resources
func (s *Context) close() error {
	var result error

	// Close all audio devices
	for _, audio := range s.audio {
		result = errors.Join(result, audio.Close())
	}

	// Close all video devices
	for _, video := range s.video {
		result = errors.Join(result, video.Close())
	}

	// Close renderer, window and surface
	if s.renderer != nil {
		result = errors.Join(result, s.renderer.Destroy())
	}
	if s.window != nil {
		result = errors.Join(result, s.window.Destroy())
	}
	if s.surface != nil {
		s.surface.Free()
	}

	// Quit SDL
	sdl.Quit()

	// Return any errors
	return result
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Enumerate capture and playback devices. Currently works for audio only.
func (ctx *Context) Devices(t Type) []*Device {
	var devices []*Device
	if t&(INPUT|OUTPUT) == 0 {
		t |= INPUT | OUTPUT
	}
	if t.Is(AUDIO | INPUT) {
		for i := 0; i < sdl.GetNumAudioDevices(true); i++ {
			devices = append(devices, newDevice(AUDIO|INPUT, sdl.GetAudioDeviceName(i, true)))
		}
	}
	if t.Is(AUDIO | OUTPUT) {
		for i := 0; i < sdl.GetNumAudioDevices(false); i++ {
			devices = append(devices, newDevice(AUDIO|OUTPUT, sdl.GetAudioDeviceName(i, false)))
		}
	}
	return devices
}

func (ctx *Context) Register(fn func(userInfo unsafe.Pointer)) uint32 {
	evt := sdl.RegisterEvents(1)
	ctx.evt[evt] = fn
	return evt
}

func (ctx *Context) Post(evt uint32, userInfo unsafe.Pointer) {
	sdl.PushEvent(&sdl.UserEvent{
		Type:  evt,
		Data1: userInfo,
	})
}

func (ctx *Context) Run(parent context.Context) error {
	var wg sync.WaitGroup
	var result error

	// Register an event which quits the application when context is cancelled
	evtCancel := sdl.RegisterEvents(1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-parent.Done()
		sdl.PushEvent(&sdl.UserEvent{
			Type: evtCancel,
		})
	}()

	// Unpause the audio
	for _, audio := range ctx.audio {
		sdl.PauseAudioDevice(audio.device, false)
	}

	// Start the runloop
	quit := false
	for {
		if quit {
			break
		}

		// Wait on an event
		var evt sdl.Event
		sdl.Do(func() {
			evt = sdl.WaitEvent()
		})

		// Handle cancel, custom, keyboard and quit events
		switch evt := evt.(type) {
		case *sdl.UserEvent:
			if evt.GetType() == evtCancel {
				quit = true
				if parent.Err() != nil {
					result = errors.Join(result, parent.Err())
				}
			} else if fn, exists := ctx.evt[evt.GetType()]; exists {
				fn(evt.Data1)
			}
		case *sdl.QuitEvent:
			quit = true
		case *sdl.AudioDeviceEvent:
			if evt.Type == sdl.AUDIODEVICEADDED {
				sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Audio device added: index=%d capture=%d name=%s", evt.Which, evt.IsCapture, sdl.GetAudioDeviceName(int(evt.Which), evt.IsCapture != 0))
			} else if evt.Type == sdl.AUDIODEVICEREMOVED {
				sdl.LogDebug(sdl.LOG_CATEGORY_APPLICATION, "Audio device removed: index=%d capture=%d name=%s", evt.Which, evt.IsCapture, sdl.GetAudioDeviceName(int(evt.Which), evt.IsCapture != 0))
			}
		default:
			sdl.LogWarn(sdl.LOG_CATEGORY_APPLICATION, "Unhandled event: %T(%d)", evt, evt.GetType())
		}
	}

	// Pause the audio
	for _, audio := range ctx.audio {
		sdl.PauseAudioDevice(audio.device, true)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Return any errors (including context errors)
	return result
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Create software surface (RGB 32-bit)
func (ctx *Context) createOffscreenRenderer(w, h int32) (*sdl.Surface, *sdl.Renderer, error) {
	var surface *sdl.Surface
	var renderer *sdl.Renderer
	var result error
	sdl.Do(func() {
		if surface, result = sdl.CreateRGBSurface(0, w, h, 32, 0, 0, 0, 0); result != nil {
			// NOOP
		} else if renderer, result = sdl.CreateSoftwareRenderer(surface); result != nil {
			surface.Free()
		}
	})

	// Return surface and renderer
	return surface, renderer, result
}

// Attach renderer to a window
func (ctx *Context) createWindowRenderer(title string, w, h int32) (*sdl.Window, *sdl.Renderer, error) {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var result error

	sdl.Do(func() {
		if window, result = sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, w, h, sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS); result != nil {
			// NOOP
		} else if renderer, result = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); result != nil {
			window.Destroy()
		}
	})
	return window, renderer, result
}
