package media

import (
	"fmt"
	"io"
	"runtime"

	// Package imports
	version "github.com/mutablelogic/go-media/pkg/version"
	ff "github.com/mutablelogic/go-media/sys/ffmpeg61"

	// Namespace imports
	. "github.com/djthorpe/go-errors"
)

////////////////////////////////////////////////////////////////////////////
// TYPES

type manager struct {
}

var _ Manager = (*manager)(nil)

////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewManager() Manager {
	return new(manager)
}

////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Return the list of matching input formats, optionally filtering by name,
// extension or mimetype File extensions should be prefixed with a dot,
// e.g. ".mp4". The media type can be NONE (for any) or combinations of
// STREAM, DEVICE.
func (manager *manager) InputFormats(t MediaType, filter ...string) []Format {
	var result []Format

	// Iterate over all input formats
	if t == NONE || t.Is(FILE) {
		var opaque uintptr
		for {
			demuxer := ff.AVFormat_demuxer_iterate(&opaque)
			if demuxer == nil {
				break
			}
			if matchesInput(demuxer, t, filter...) {
				result = append(result, newInputFormat(demuxer, FILE))
			}
		}
	}

	if t == NONE || t.Is(DEVICE) {
		// Iterate over all device inputs
		audio := ff.AVDevice_input_audio_device_first()
		for {
			if audio == nil {
				break
			}
			if matchesInput(audio, t, filter...) {
				result = append(result, newInputFormat(audio, AUDIO|DEVICE))
			}
			audio = ff.AVDevice_input_audio_device_next(audio)
		}

		video := ff.AVDevice_input_video_device_first()
		for {
			if video == nil {
				break
			}
			if matchesInput(video, t, filter...) {
				result = append(result, newInputFormat(video, VIDEO|DEVICE))
			}
			video = ff.AVDevice_input_video_device_next(video)
		}
	}

	// Return success
	return result
}

// Return the list of matching output formats, optionally filtering by name,
// extension or mimetype File extensions should be prefixed with a dot,
// e.g. ".mp4". The media type can be NONE (for any) or combinations of
// STREAM, DEVICE.
func (manager *manager) OutputFormats(t MediaType, filter ...string) []Format {
	var result []Format

	// Iterate over all output formats
	if t == NONE || t.Is(FILE) {
		var opaque uintptr
		for {
			muxer := ff.AVFormat_muxer_iterate(&opaque)
			if muxer == nil {
				break
			}
			if matchesOutput(muxer, t, filter...) {
				result = append(result, newOutputFormat(muxer, FILE))
			}
		}
	}

	// Iterate over all device outputs
	if t == NONE || t.Is(DEVICE) {
		audio := ff.AVDevice_output_audio_device_first()
		for {
			if audio == nil {
				break
			}
			if matchesOutput(audio, t, filter...) {
				result = append(result, newOutputFormat(audio, AUDIO|DEVICE))
			}
			audio = ff.AVDevice_output_audio_device_next(audio)
		}

		video := ff.AVDevice_output_video_device_first()
		for {
			if video == nil {
				break
			}
			if matchesOutput(video, t, filter...) {
				result = append(result, newOutputFormat(video, VIDEO|DEVICE))
			}
			video = ff.AVDevice_output_video_device_next(video)
		}
	}

	// Return success
	return result
}

// Return supported input devices for a given input format. Sometimes
// (ie, AVFoundation) there is a option which provides the input
// devices and this function returns an empty string instead. Go figure!
func (manager *manager) InputDevices(format string) []Device {
	input := ff.AVFormat_find_input_format(format)
	if input == nil {
		return nil
	}

	device_list, err := ff.AVDevice_list_input_sources(input, format, nil)
	if err != nil {
		panic(err)
	}
	if device_list == nil {
		return nil
	}
	defer ff.AVDevice_free_list_devices(device_list)

	// Iterate over devices
	result := make([]Device, 0, device_list.NumDevices())
	for i, device := range device_list.Devices() {
		fmt.Println(i, device)
	}

	return result
}

// Return supported output devices for a given name
func (manager *manager) OutputDevices(format string) []Device {
	panic("TODO")
}

// Return all supported channel layouts
func (manager *manager) ChannelLayouts() []Metadata {
	var result []Metadata
	var iter uintptr
	for {
		ch := ff.AVUtil_channel_layout_standard(&iter)
		if ch == nil {
			break
		}
		if name, err := ff.AVUtil_channel_layout_describe(ch); err != nil {
			continue
		} else {
			result = append(result, newMetadata(name, ch))
		}
	}
	return result
}

// Return all supported sample formats
func (manager *manager) SampleFormats() []Metadata {
	var result []Metadata
	var opaque uintptr
	for {
		samplefmt := ff.AVUtil_next_sample_fmt(&opaque)
		if samplefmt == ff.AV_SAMPLE_FMT_NONE {
			break
		}
		if name := ff.AVUtil_get_sample_fmt_name(samplefmt); name != "" {
			result = append(result, newMetadata(name, samplefmt))
		}
	}
	return result
}

// Return all supported  pixel formats
func (manager *manager) PixelFormats() []Metadata {
	var result []Metadata
	var opaque uintptr
	for {
		pixfmt := ff.AVUtil_next_pixel_fmt(&opaque)
		if pixfmt == ff.AV_PIX_FMT_NONE {
			break
		}
		if name := ff.AVUtil_get_pix_fmt_name(pixfmt); name != "" {
			result = append(result, newMetadata(name, pixfmt))
		}
	}
	return result
}

// Return all supported codecs
func (manager *manager) Codecs() []Metadata {
	var result []Metadata
	var iter uintptr
	for {
		codec := ff.AVCodec_iterate(&iter)
		if codec == nil {
			break
		}
		if name := codec.Name(); name != "" {
			result = append(result, newMetadata(name, newCodec(codec)))
		}
	}
	return result
}

// Return audio parameters for encoding
// ChannelLayout, SampleFormat, Samplerate
func (manager *manager) AudioParameters(channels string, samplefmt string, samplerate int) (AudioParameters, error) {
	return newAudioParametersEx(channels, samplefmt, samplerate)
}

// Return video parameters for encoding
// Width, Height, PixelFormat, Framerate
func (manager *manager) VideoParameters(width int, height int, pixelfmt string) (VideoParameters, error) {
	return newVideoParametersEx(width, height, pixelfmt)
}

// Open a media file or device for reading, from a path or url.
func (manager *manager) Open(url string, format Format, opts ...string) (Media, error) {
	return newMedia(url, format, opts...)
}

// Open a media stream for reading.
func (manager *manager) Read(r io.Reader, format Format, opts ...string) (Media, error) {
	return newReader(r, format, opts...)
}

// Create a media file for writing, from a path.
func (manager *manager) Create(string, Format) (Media, error) {
	return nil, ErrNotImplemented
}

// Create a media stream for writing.
func (manager *manager) Write(io.Writer, Format) (Media, error) {
	return nil, ErrNotImplemented
}

// Return version information for the media manager as a set of metadata
func (manager *manager) Version() []Metadata {
	metadata := []Metadata{
		newMetadata("libavcodec_version", ff.AVCodec_version()),
		newMetadata("libavformat_version", ff.AVFormat_version()),
		newMetadata("libavutil_version", ff.AVUtil_version()),
		newMetadata("libavdevice_version", ff.AVDevice_version()),
		//		newMetadata("libavfilter_version", ff.AVFilter_version()),
		newMetadata("libswscale_version", ff.SWScale_version()),
		newMetadata("libswresample_version", ff.SWResample_version()),
	}
	if version.GitSource != "" {
		metadata = append(metadata, newMetadata("git_source", version.GitSource))
	}
	if version.GitBranch != "" {
		metadata = append(metadata, newMetadata("git_branch", version.GitBranch))
	}
	if version.GitTag != "" {
		metadata = append(metadata, newMetadata("git_tag", version.GitTag))
	}
	if version.GoBuildTime != "" {
		metadata = append(metadata, newMetadata("go_build_time", version.GoBuildTime))
	}
	if runtime.Version() != "" {
		metadata = append(metadata, newMetadata("go_version", runtime.Version()))
		metadata = append(metadata, newMetadata("go_arch", runtime.GOOS+"/"+runtime.GOARCH))
	}
	return metadata
}
