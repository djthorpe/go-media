package ffmpeg

import (
	"encoding/json"

	// Packages
	ff "github.com/mutablelogic/go-media/sys/ffmpeg61"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Packet ff.AVPacket

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (packet *Packet) String() string {
	data, _ := json.MarshalIndent((*ff.AVPacket)(packet), "", "  ")
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Return the timestamp in seconds, or PTS_UNDEFINED if the timestamp
// is undefined or timebase is not set
func (packet *Packet) Ts() float64 {
	if packet == nil {
		return PTS_UNDEFINED
	}
	if pts := (*ff.AVPacket)(packet).Pts(); pts == ff.AV_NOPTS_VALUE {
		return PTS_UNDEFINED
	} else if tb := (*ff.AVPacket)(packet).TimeBase(); tb.Num() == 0 || tb.Den() == 0 {
		return PTS_UNDEFINED
	} else {
		return ff.AVUtil_rational_q2d(tb) * float64(pts)
	}
}
