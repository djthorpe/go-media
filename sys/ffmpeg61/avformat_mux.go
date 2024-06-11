package ffmpeg

import (
	"errors"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////
// CGO

/*
#cgo pkg-config: libavformat
#include <libavformat/avformat.h>
*/
import "C"

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Open an output stream.
func AVFormat_open_writer(writer *AVIOContextEx, format *AVOutputFormat, filename string) (*AVFormatContext, error) {
	// TODO
	return nil, errors.New("not implemented")
}

// Open an output file.
func AVFormat_create_file(filename string, format *AVOutputFormat) (*AVFormatContext, error) {
	var ctx *AVFormatContext
	if err := AVError(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), (*C.struct_AVOutputFormat)(format), nil, C.CString(filename))); err != 0 {
		return nil, err
	} else {
		return ctx, nil
	}
}

func AVFormat_close_writer(ctx *AVFormatContext) error {
	var result error

	octx := (*C.struct_AVFormatContext)(ctx)
	if octx.oformat.flags&C.int(AVFMT_NOFILE) == 0 && octx.pb != nil {
		if err := AVError(C.avio_closep(&octx.pb)); err != 0 {
			result = errors.Join(result, err)
		}
	}
	C.avformat_free_context(octx)

	// Return any errors
	return result
}

// Allocate an AVFormatContext for an output format.
func AVFormat_alloc_output_context2(ctx **AVFormatContext, format *AVOutputFormat, filename string) error {
	var cFilename *C.char
	if filename != "" {
		cFilename = C.CString(filename)
	}
	defer C.free(unsafe.Pointer(cFilename))
	if err := AVError(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), (*C.struct_AVOutputFormat)(format), nil, cFilename)); err != 0 {
		return err
	}

	// Return success
	return nil
}

// Allocate the stream private data and initialize the codec, but do not write the header.
// May optionally be used before avformat_write_header() to initialize stream parameters before actually writing the header.
func AVFormat_init_output(ctx *AVFormatContext, options **AVDictionary) error {
	if err := AVError(C.avformat_init_output((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(options)))); err != 0 {
		return err
	} else {
		return nil
	}
}

// Allocate the stream private data and write the stream header to an output media file.
func AVFormat_write_header(ctx *AVFormatContext, options **AVDictionary) error {
	if err := AVError(C.avformat_write_header((*C.struct_AVFormatContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(options)))); err != 0 {
		return err
	} else {
		return nil
	}
	// TODO:
	// AVSTREAM_INIT_IN_WRITE_HEADER
	// AVSTREAM_INIT_IN_INIT_OUTPUT
}

// Write a packet to an output media file. Returns true if flushed and there is
// no more data to flush.
func AVFormat_write_frame(ctx *AVFormatContext, pkt *AVPacket) (bool, error) {
	if err := AVError(C.av_write_frame((*C.struct_AVFormatContext)(ctx), (*C.struct_AVPacket)(pkt))); err < 0 {
		return false, err
	} else if err == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// Write a packet to an output media file ensuring correct interleaving.
func AVFormat_interleaved_write_frame(ctx *AVFormatContext, pkt *AVPacket) error {
	if err := AVError(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ctx), (*C.struct_AVPacket)(pkt))); err != 0 {
		return err
	} else {
		return nil
	}
}

// Write the stream trailer to an output media file and free the file private data.
func AVFormat_write_trailer(ctx *AVFormatContext) error {
	if err := AVError(C.av_write_trailer((*C.struct_AVFormatContext)(ctx))); err != 0 {
		return err
	} else {
		return nil
	}
}

// Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AVFormat_guess_format(format, filename, mime_type string) *AVOutputFormat {
	var cFilename, cFormat, cMimeType *C.char
	if format != "" {
		cFormat = C.CString(format)
	}
	if filename != "" {
		cFilename = C.CString(filename)
	}
	if mime_type != "" {
		cMimeType = C.CString(mime_type)
	}
	defer C.free(unsafe.Pointer(cFormat))
	defer C.free(unsafe.Pointer(cFilename))
	defer C.free(unsafe.Pointer(cMimeType))
	return (*AVOutputFormat)(C.av_guess_format(cFormat, cFilename, cMimeType))
}