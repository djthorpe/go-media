package sdl

import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"

	// Packages
	media "github.com/mutablelogic/go-media"
	ffmpeg "github.com/mutablelogic/go-media/pkg/ffmpeg"
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

type Audio struct {
	t      media.Type
	device sdl.AudioDeviceID
}

// Audio callback function
type audioCallbackFn func(data unsafe.Pointer, length C.int)

//////////////////////////////////////////////////////////////////////////////
// GLOBALS

/*
	SDL Supported audio formats:
	AUDIO_S8         // signed 8-bit samples
	AUDIO_U8         // unsigned 8-bit samples
	AUDIO_S16LSB     // signed 16-bit samples in little-endian byte order
	AUDIO_S16MSB     // signed 16-bit samples in big-endian byte order
	AUDIO_S16SYS     // signed 16-bit samples in native byte order
	AUDIO_S16        // AUDIO_S16LSB
	AUDIO_U16LSB     // unsigned 16-bit samples in little-endian byte order
	AUDIO_U16MSB     // unsigned 16-bit samples in big-endian byte order
	AUDIO_U16SYS     // unsigned 16-bit samples in native byte order
	AUDIO_U16        // AUDIO_U16LSB
	AUDIO_S32LSB     // 32-bit integer samples in little-endian byte order
	AUDIO_S32MSB     // 32-bit integer samples in big-endian byte order
	AUDIO_S32SYS     // 32-bit integer samples in native byte order
	AUDIO_S32        // AUDIO_S32LSB
	AUDIO_F32LSB     // 32-bit floating point samples in little-endian byte order
	AUDIO_F32MSB     // 32-bit floating point samples in big-endian byte order
	AUDIO_F32SYS     // 32-bit floating point samples in native byte order
	AUDIO_F32        // AUDIO_F32LSB
*/

var (
	mapAudio = map[string]sdl.AudioFormat{
		"u8":  sdl.AUDIO_U8,
		"s8":  sdl.AUDIO_S8,
		"u16": sdl.AUDIO_U16,
		"s16": sdl.AUDIO_S16,
		"flt": sdl.AUDIO_F32,
	}
	mapDevice []audioCallbackFn
)

//////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewAudio creates a new audio device with the specified parameters.
func (s *Context) NewAudio(par *ffmpeg.Par, device *Device, dur time.Duration) (*Audio, error) {
	if device == nil || !device.Type().Is(AUDIO) {
		return nil, ErrBadParameter.With("invalid device type")
	}
	if par == nil || !par.Type().Is(AUDIO) {
		return nil, ErrBadParameter.With("invalid audio parameters")
	}
	if dur < 0 {
		return nil, ErrBadParameter.With("invalid duration")
	}

	// Get the audio format
	src_format := fmt.Sprint(par.SampleFormat())
	format, exists := mapAudio[src_format]
	if !exists {
		return nil, ErrBadParameter.Withf("unsupported sample format %q", src_format)
	}

	// Create a new audio device
	audio := new(Audio)

	// Set the audio parameters
	var desired, obtained sdl.AudioSpec
	desired.Freq = int32(par.Samplerate())
	desired.Format = format
	desired.Channels = uint8(par.ChannelLayout().NumChannels())
	desired.Samples = uint16(dur.Seconds() * float64(par.Samplerate()))
	desired.Callback = sdl.AudioCallback(C.audio_callback)

	// Create a callback
	switch format {
	case sdl.AUDIO_U8:
		mapDevice = append(mapDevice, audio.callbackUint8)
	case sdl.AUDIO_S8:
		mapDevice = append(mapDevice, audio.callbackInt8)
	case sdl.AUDIO_U16:
		mapDevice = append(mapDevice, audio.callbackUint16)
	case sdl.AUDIO_S16:
		mapDevice = append(mapDevice, audio.callbackInt16)
	case sdl.AUDIO_F32:
		mapDevice = append(mapDevice, audio.callbackFloat32)
	}
	desired.UserData = unsafe.Pointer(uintptr(len(mapDevice)))

	// Open audio device
	if id, err := sdl.OpenAudioDevice(device.Name(), device.Type().Is(INPUT), &desired, &obtained, 0); err != nil {
		return nil, err
	} else if desired.Freq != obtained.Freq {
		return nil, ErrBadParameter.Withf("sample rate %d not supported", desired.Freq)
	} else if desired.Format != obtained.Format {
		return nil, ErrBadParameter.Withf("unsupported sample format %q", src_format)
	} else if desired.Channels != obtained.Channels {
		return nil, ErrBadParameter.Withf("number of channels %d not supported", desired.Channels)
	} else {
		audio.t = device.Type()
		audio.device = id
	}

	// Add to map
	if _, exists := s.audio[audio.device]; exists {
		sdl.CloseAudioDevice(audio.device)
		return nil, ErrDuplicateEntry.Withf("audio device %v already exists", audio.device)
	} else {
		s.audio[audio.device] = audio
	}

	// Return success
	return audio, nil
}

func (a *Audio) Close() error {
	var result error

	// Close the audio device
	sdl.CloseAudioDevice(a.device)

	// Return any errors
	return result
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (a *Audio) MarshalJSON() ([]byte, error) {
	type j struct {
		Device sdl.AudioDeviceID `json:"device"`
		Type   media.Type        `json:"type"`
	}
	return json.Marshal(j{
		Device: a.device,
		Type:   a.t,
	})
}

func (a *Audio) String() string {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// callback function to capture audio data
//
//export audio_callback
func audio_callback(userData unsafe.Pointer, data unsafe.Pointer, length C.int) {
	idx := int(uintptr(userData)) - 1
	if idx < len(mapDevice) && mapDevice[idx] != nil {
		mapDevice[idx](data, length)
	}
}

func (a *Audio) callbackUint8(data unsafe.Pointer, length C.int) {
	s := cUint8Slice(data, length)
	fmt.Println(s)
}

func (a *Audio) callbackInt8(data unsafe.Pointer, length C.int) {
	s := cInt8Slice(data, length)
	fmt.Println(s)
}

func (a *Audio) callbackUint16(data unsafe.Pointer, length C.int) {
	s := cUint16Slice(data, length>>1)
	fmt.Println(s)
}

func (a *Audio) callbackInt16(data unsafe.Pointer, length C.int) {
	s := cInt16Slice(data, length>>1)
	fmt.Println(s)
}

func (a *Audio) callbackFloat32(data unsafe.Pointer, length C.int) {
	s := cFloat32Slice(data, length>>2)
	fmt.Println(s)
}

///////////////////////////////////////////////////////////////////////////////
// REALLY PRIVATE METHODS

func cUint8Slice(p unsafe.Pointer, sz C.int) []uint8 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint8)(p)[:int(sz)]
}

func cInt8Slice(p unsafe.Pointer, sz C.int) []int8 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]int8)(p)[:int(sz)]
}

func cUint16Slice(p unsafe.Pointer, sz C.int) []uint16 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint16)(p)[:int(sz)]
}

func cInt16Slice(p unsafe.Pointer, sz C.int) []int16 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]int16)(p)[:int(sz)]
}

func cUint32Slice(p unsafe.Pointer, sz C.int) []uint32 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint32)(p)[:int(sz)]
}

func cInt32Slice(p unsafe.Pointer, sz C.int) []int32 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]int32)(p)[:int(sz)]
}

func cFloat32Slice(p unsafe.Pointer, sz C.int) []float32 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]float32)(p)[:int(sz)]
}
