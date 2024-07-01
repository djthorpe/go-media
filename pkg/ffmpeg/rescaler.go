package ffmpeg

import (
	"errors"

	// Packages

	ff "github.com/mutablelogic/go-media/sys/ffmpeg61"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type rescaler struct {
	src_pix_fmt ff.AVPixelFormat
	src_width   int
	src_height  int
	ctx         *ff.SWSContext
	flags       ff.SWSFlag
	force       bool
	dest        *ff.AVFrame
}

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new rescaler which will rescale the input frame to the
// specified format, width and height.
func NewRescaler(par *Par, force bool) (*rescaler, error) {
	rescaler := new(rescaler)

	// Check parameters
	if par == nil || par.CodecType() != ff.AVMEDIA_TYPE_AUDIO {
		return nil, errors.New("invalid codec type")
	}
	if par.SampleFormat() == ff.AV_SAMPLE_FMT_NONE {
		return nil, errors.New("invalid sample format parameters")
	}
	ch := par.ChannelLayout()
	if !ff.AVUtil_channel_layout_check(&ch) {
		return nil, errors.New("invalid channel layout parameters")
	}

	// Apply options
	options.par.SetCodecType(ff.AVMEDIA_TYPE_VIDEO)
	options.par.SetPixelFormat(format)
	options.par.SetWidth(640)
	options.par.SetHeight(480)
	for _, o := range opt {
		if err := o(options); err != nil {
			return nil, err
		}
	}

	// Check parameters
	if options.par.PixelFormat() == ff.AV_PIX_FMT_NONE {
		return nil, errors.New("invalid parameters")
	}

	// Create a destimation frame
	dest := ff.AVUtil_frame_alloc()
	if dest == nil {
		return nil, errors.New("failed to allocate frame")
	}

	// Set force flag
	rescaler.force = options.force

	// Set parameters
	dest.SetPixFmt(options.par.PixelFormat())
	dest.SetWidth(options.par.Width())
	dest.SetHeight(options.par.Height())

	// Allocate buffer
	if err := ff.AVUtil_frame_get_buffer(dest, false); err != nil {
		ff.AVUtil_frame_free(dest)
		return nil, err
	} else {
		rescaler.dest = dest
		rescaler.flags = ff.SWS_POINT
	}

	// Return success
	return rescaler, nil
}

// Release resources
func (r *rescaler) Close() error {
	if r.ctx != nil {
		ff.SWScale_free_context(r.ctx)
		r.ctx = nil
	}
	if r.dest != nil {
		ff.AVUtil_frame_free(r.dest)
		r.dest = nil
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Scale the source image and return the destination image
func (r *rescaler) Frame(src *ff.AVFrame) (*ff.AVFrame, error) {
	// If source is null then return null (no flushing)
	if src == nil {
		return nil, nil
	}

	// Simply return the frame if it matches the destination format
	if matchesVideoFormat(src, r.dest) && !r.force {
		return src, nil
	}

	// Allocate a context
	if r.ctx == nil || r.src_pix_fmt != src.PixFmt() || r.src_width != src.Width() || r.src_height != src.Height() {
		// Release existing scaling context, if any
		if r.ctx != nil {
			ff.SWScale_free_context(r.ctx)
		}
		// Create a new scaling context
		ctx := ff.SWScale_get_context(
			src.Width(), src.Height(), src.PixFmt(), // source
			r.dest.Width(), r.dest.Height(), r.dest.PixFmt(), // destination
			r.flags, nil, nil, nil,
		)
		if ctx == nil {
			return nil, errors.New("failed to allocate swscale context")
		} else {
			r.ctx = ctx
			r.src_pix_fmt = src.PixFmt()
			r.src_width = src.Width()
			r.src_height = src.Height()
		}
	}

	// Rescale the image
	if err := ff.SWScale_scale_frame(r.ctx, r.dest, src, false); err != nil {
		return nil, err
	}

	// Copy parameters from the source frame
	if err := ff.AVUtil_frame_copy_props(r.dest, src); err != nil {
		return nil, err
	}

	// Return the destination frame
	return r.dest, nil
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Returns true if the pixel format, width and height of the source
// and destination frames match
func matchesVideoFormat(src, dest *ff.AVFrame) bool {
	if src.PixFmt() == dest.PixFmt() && src.Width() == dest.Width() && src.Height() == dest.Height() {
		return true
	}
	return false
}
