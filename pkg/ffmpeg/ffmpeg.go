package ffmpeg

import (
	ff "github.com/mutablelogic/go-media/sys/ffmpeg61"
)

////////////////////////////////////////////////////////////////////////////
// TYPES

type MediaType int

////////////////////////////////////////////////////////////////////////////
// GLOBALS

// Media Types
const (
	AUDIO = MediaType(ff.AVMEDIA_TYPE_AUDIO)
	VIDEO = MediaType(ff.AVMEDIA_TYPE_VIDEO)
)
