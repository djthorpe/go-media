package ffmpeg

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"sync"
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
// TYPES

type (
	AVFormatContext C.struct_AVFormatContext
	AVInputFormat   C.struct_AVInputFormat
	AVOutputFormat  C.struct_AVOutputFormat
)

////////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	once_init, once_deinit sync.Once
)

////////////////////////////////////////////////////////////////////////////////
// INIT AND DEINIT

// Register and Deregister
func AVFormatInit() {
	once_init.Do(func() {
		C.avformat_network_init()
	})
}

func AVFormatDeinit() {
	once_deinit.Do(func() {
		C.avformat_network_deinit()
	})
}

func AVFormatVersion() uint {
	return uint(C.avformat_version())
}

////////////////////////////////////////////////////////////////////////////////
// GET FORMATS

// AllMuxers returns all registered multiplexers
func AllMuxers() []*AVOutputFormat {
	muxers := make([]*AVOutputFormat, 0)
	ptr := unsafe.Pointer(nil)
	for {
		if muxer := C.av_muxer_iterate(&ptr); muxer == nil {
			break
		} else {
			muxers = append(muxers, (*AVOutputFormat)(muxer))
		}
	}
	return muxers
}

// AllDemuxers returns all registered demultiplexers
func AllDemuxers() []*AVInputFormat {
	demuxers := make([]*AVInputFormat, 0)
	ptr := unsafe.Pointer(nil)
	for {
		if demuxer := C.av_demuxer_iterate(&ptr); demuxer == nil {
			break
		} else {
			demuxers = append(demuxers, (*AVInputFormat)(demuxer))
		}
	}
	return demuxers
}

////////////////////////////////////////////////////////////////////////////////
// AVFormatContext

// NewAVFormatContext creates a new empty format context
func NewAVFormatContext() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

// NewAVFormatOutputContext creates a new format context with
// context populated with output parameters
func NewAVFormatOutputContext(filename string, output_format *AVOutputFormat) (*AVFormatContext, error) {
	filename_ := C.CString(filename)
	defer C.free(unsafe.Pointer(filename_))
	ctx := (*C.AVFormatContext)(unsafe.Pointer(nil))
	if err := AVError(C.avformat_alloc_output_context2(
		&ctx,
		(*C.AVOutputFormat)(output_format),
		nil,
		filename_,
	)); err != 0 {
		return nil, err
	} else {
		return (*AVFormatContext)(ctx), nil
	}
}

// Free AVFormatContext
func (this *AVFormatContext) Free() {
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	C.avformat_free_context(ctx)
}

// Open Input
func (this *AVFormatContext) OpenInput(filename string, input_format *AVInputFormat) error {
	filename_ := C.CString(filename)
	defer C.free(unsafe.Pointer(filename_))
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	dict := new(AVDictionary)
	if err := AVError(C.avformat_open_input(
		&ctx,
		filename_,
		(*C.struct_AVInputFormat)(input_format),
		(**C.struct_AVDictionary)(unsafe.Pointer(dict)),
	)); err != 0 {
		return err
	} else {
		return nil
	}
}

// Open Input URL
func (this *AVFormatContext) OpenInputUrl(url string, input_format *AVInputFormat) error {
	url_ := C.CString(url)
	defer C.free(unsafe.Pointer(url_))
	ctx := (*C.AVFormatContext)(this)
	dict := new(AVDictionary)
	if err := AVError(C.avformat_open_input(
		&ctx,
		url_,
		(*C.struct_AVInputFormat)(input_format),
		(**C.struct_AVDictionary)(unsafe.Pointer(dict)),
	)); err != 0 {
		return err
	} else {
		return nil
	}
}

// Open Input with AVIO Context
func (this *AVFormatContext) OpenInputIO(io *AVIOContext, input_format *AVInputFormat) error {
	ctx := (*C.AVFormatContext)(this)
	this.SetIOContext(io)
	if err := AVError(C.avformat_open_input(&ctx, nil, (*C.struct_AVInputFormat)(input_format), nil)); err != 0 {
		this.SetIOContext(nil)
		return err
	} else {
		return nil
	}
}

// Close Input
func (this *AVFormatContext) CloseInput() {
	ctx := (*C.AVFormatContext)(this)
	C.avformat_close_input(&ctx)
}

// Write header
func (this *AVFormatContext) WriteHeader(dict *AVDictionary) error {
	ctx := (*C.AVFormatContext)(this)
	dictctx := (**C.AVDictionary)(nil)
	if dict != nil {
		dictctx = &dict.ctx
	}
	if ret := AVError(C.avformat_write_header(ctx, dictctx)); ret != 0 {
		return ret
	} else {
		return nil
	}
}

// Write trailer
func (this *AVFormatContext) WriteTrailer() error {
	ctx := (*C.AVFormatContext)(this)
	if ret := AVError(C.av_write_trailer(ctx)); ret != 0 {
		return ret
	} else {
		return nil
	}
}

// Return Metadata Dictionary
func (this *AVFormatContext) Metadata() *AVDictionary {
	return &AVDictionary{ctx: this.metadata}
}

// Find Stream Info
func (this *AVFormatContext) FindStreamInfo() error {
	ctx := (*C.AVFormatContext)(this)
	if err := AVError(C.avformat_find_stream_info(ctx, nil)); err != 0 {
		return err
	} else {
		return nil
	}
}

// Return Filename
func (this *AVFormatContext) Filename() string {
	return C.GoString(&this.filename[0])
}

// Return URL
func (this *AVFormatContext) Url() *url.URL {
	url_ := C.GoString(this.url)
	if url_ == "" {
		return nil
	} else if url, err := url.Parse(url_); err != nil {
		return nil
	} else {
		return url
	}
}

// Return flags
func (this *AVFormatContext) Flags() AVFormatFlag {
	ctx := (*C.AVFormatContext)(this)
	return AVFormatFlag(ctx.flags)
}

// Return number of streams
func (this *AVFormatContext) NumStreams() uint {
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	return uint(ctx.nb_streams)
}

// Return Streams
func (this *AVFormatContext) Streams() []*AVStream {
	var streams []*AVStream

	// Get context
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))

	// Make a fake slice
	if nb_streams := this.NumStreams(); nb_streams > 0 {
		// Make a fake slice
		sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&streams)))
		sliceHeader.Cap = int(nb_streams)
		sliceHeader.Len = int(nb_streams)
		sliceHeader.Data = uintptr(unsafe.Pointer(ctx.streams))
	}
	return streams
}

// Return Input Format
func (this *AVFormatContext) InputFormat() *AVInputFormat {
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	return (*AVInputFormat)(ctx.iformat)
}

// Return Output Format
func (this *AVFormatContext) OutputFormat() *AVOutputFormat {
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	return (*AVOutputFormat)(ctx.oformat)
}

// IOContext returns the current IO context
func (this *AVFormatContext) IOContext() *AVIOContext {
	ctx := (*C.AVFormatContext)(this)
	return (*AVIOContext)(ctx.pb)
}

// SetIOContext sets the current IO context
func (this *AVFormatContext) SetIOContext(avio *AVIOContext) {
	ctx := (*C.AVFormatContext)(this)
	ctx.pb = (*C.AVIOContext)(avio)
}

func (this *AVFormatContext) String() string {
	str := "<AVFormatContext"
	if fn := this.Filename(); fn != "" {
		str += fmt.Sprintf(" filename=%q", fn)
	}
	if u := this.Url(); u != nil {
		str += fmt.Sprintf(" url=%q", u.String())
	}
	if io := this.IOContext(); io != nil {
		str += " io=" + fmt.Sprint(io)
	}
	if ifmt := this.InputFormat(); ifmt != nil {
		str += " iformat=" + fmt.Sprint(ifmt)
	}
	if ofmt := this.OutputFormat(); ofmt != nil {
		str += " oformat=" + fmt.Sprint(ofmt)
	}
	str += " num_streams=" + fmt.Sprint(this.NumStreams())
	str += " metadata=" + fmt.Sprint(this.Metadata())
	return str + ">"
}

func (this *AVFormatContext) Dump(index int) {
	str := ""
	if url := this.Url(); url != nil {
		str = url.String()
	}
	if this.OutputFormat() != nil {
		AVDumpFormat(this, index, str, true)
	} else if this.InputFormat() != nil {
		AVDumpFormat(this, index, str, false)
	}
}

func (this *AVFormatContext) ReadPacket(packet *AVPacket) error {
	ctx := (*C.AVFormatContext)(unsafe.Pointer(this))
	packetctx := (*C.AVPacket)(packet)
	if ret := int(C.av_read_frame(ctx, packetctx)); ret >= 0 {
		return nil
	} else {
		return io.EOF
	}
}

func (this *AVFormatContext) WritePacket(packet *AVPacket, out *AVFormatContext) error {
	i := (*C.AVFormatContext)(unsafe.Pointer(this))
	o := (*C.AVFormatContext)(unsafe.Pointer(out))
	p := (*C.AVPacket)(packet)

	/* Get streams */
	in_stream := *(i.streams)  // TODO
	out_stream := *(o.streams) // TODO

	/* Adjust packet params for output */
	p.pts = C.av_rescale_q_rnd(p.pts, in_stream.time_base, out_stream.time_base, C.AV_ROUND_NEAR_INF|C.AV_ROUND_PASS_MINMAX)
	p.dts = C.av_rescale_q_rnd(p.dts, in_stream.time_base, out_stream.time_base, C.AV_ROUND_NEAR_INF|C.AV_ROUND_PASS_MINMAX)
	p.duration = C.av_rescale_q(p.duration, in_stream.time_base, out_stream.time_base)
	p.pos = -1

	/* Write packet */
	if ret := AVError(C.av_interleaved_write_frame(o, p)); ret == 0 {
		return nil
	} else {
		return ret
	}
}

////////////////////////////////////////////////////////////////////////////////
// AVInputFormat and AVOutputFormat

// Return input formats
func EnumerateInputFormats() []*AVInputFormat {
	a := make([]*AVInputFormat, 0, 100)
	p := unsafe.Pointer(uintptr(0))
	for {
		if iformat := (*AVInputFormat)(C.av_demuxer_iterate(&p)); iformat == nil {
			break
		} else {
			a = append(a, iformat)
		}
	}
	return a
}

// Return output formats
func EnumerateOutputFormats() []*AVOutputFormat {
	a := make([]*AVOutputFormat, 0, 100)
	p := unsafe.Pointer(uintptr(0))
	for {
		if oformat := (*AVOutputFormat)(C.av_muxer_iterate(&p)); oformat == nil {
			break
		} else {
			a = append(a, oformat)
		}
	}
	return a
}

func (this *AVInputFormat) Name() string {
	return C.GoString(this.name)
}

func (this *AVInputFormat) Description() string {
	return C.GoString(this.long_name)
}

func (this *AVInputFormat) Ext() string {
	return C.GoString(this.extensions)
}

func (this *AVInputFormat) MimeType() string {
	return C.GoString(this.mime_type)
}

func (this *AVInputFormat) Flags() AVFormatFlag {
	return AVFormatFlag(this.flags)
}

func (this *AVOutputFormat) Name() string {
	return C.GoString(this.name)
}

func (this *AVOutputFormat) Description() string {
	return C.GoString(this.long_name)
}

func (this *AVOutputFormat) Ext() string {
	return C.GoString(this.extensions)
}

func (this *AVOutputFormat) MimeType() string {
	return C.GoString(this.mime_type)
}

func (this *AVOutputFormat) Flags() AVFormatFlag {
	return AVFormatFlag(this.flags)
}

func (this *AVInputFormat) String() string {
	str := "<AVInputFormat"
	str += " name=" + strconv.Quote(this.Name())
	str += " description=" + strconv.Quote(this.Description())
	str += " ext=" + strconv.Quote(this.Ext())
	str += " mime_type=" + strconv.Quote(this.MimeType())
	str += " flags=" + fmt.Sprint(this.Flags())
	return str + ">"
}

func (this *AVOutputFormat) String() string {
	str := "<AVOutputFormat"
	str += " name=" + strconv.Quote(this.Name())
	str += " description=" + strconv.Quote(this.Description())
	str += " ext=" + strconv.Quote(this.Ext())
	str += " mime_type=" + strconv.Quote(this.MimeType())
	str += " flags=" + fmt.Sprint(this.Flags())
	return str + ">"
}

////////////////////////////////////////////////////////////////////////////////
// UTILITY METHODS

func AVDumpFormat(ctx *AVFormatContext, index int, filename string, is_output bool) {
	filename_ := C.CString(filename)
	defer C.free(unsafe.Pointer(filename_))
	is_output_ := 0
	if is_output {
		is_output_ = 1
	}
	C.av_dump_format((*C.AVFormatContext)(ctx), C.int(index), filename_, C.int(is_output_))
}
