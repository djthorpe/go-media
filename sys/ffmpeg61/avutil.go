package ffmpeg

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// CGO

/*
#cgo pkg-config: libavutil
#include <libavutil/avutil.h>
#include <libavutil/dict.h>
*/
import "C"

////////////////////////////////////////////////////////////////////////////////
// TYPES

type (
	AVClass           C.AVClass
	AVError           C.int
	AVDictionary      C.struct_AVDictionary
	AVDictionaryEntry C.struct_AVDictionaryEntry
	AVDictionaryFlag  C.int
	AVMediaType       C.enum_AVMediaType
	AVRational        C.AVRational
	AVPixelFormat     C.enum_AVPixelFormat
)

type jsonAVClass struct {
	ClassName string `json:"class_name"`
}

type jsonAVDictionary struct {
	Count int                  `json:"count"`
	Elems []*AVDictionaryEntry `json:"elems,omitempty"`
}

type jsonAVDictionaryEntry struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	AVMEDIA_TYPE_UNKNOWN    AVMediaType = C.AVMEDIA_TYPE_UNKNOWN ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO      AVMediaType = C.AVMEDIA_TYPE_VIDEO
	AVMEDIA_TYPE_AUDIO      AVMediaType = C.AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA       AVMediaType = C.AVMEDIA_TYPE_DATA ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE   AVMediaType = C.AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT AVMediaType = C.AVMEDIA_TYPE_ATTACHMENT ///< Opaque data information usually sparse
)

const (
	// Only get an entry with exact-case key match.
	AV_DICT_MATCH_CASE AVDictionaryFlag = C.AV_DICT_MATCH_CASE

	// Return first entry in a dictionary whose first part corresponds to the search key, ignoring the suffix of the found key string.
	AV_DICT_IGNORE_SUFFIX AVDictionaryFlag = C.AV_DICT_IGNORE_SUFFIX

	// Take ownership of  key that has been allocated with av_malloc()
	AV_DICT_DONT_STRDUP_KEY AVDictionaryFlag = C.AV_DICT_DONT_STRDUP_KEY

	// Take ownership of  value that has been allocated with av_malloc()
	AV_DICT_DONT_STRDUP_VAL AVDictionaryFlag = C.AV_DICT_DONT_STRDUP_VAL

	// Don't overwrite existing entries.
	AV_DICT_DONT_OVERWRITE AVDictionaryFlag = C.AV_DICT_DONT_OVERWRITE

	// Append to existing key.
	AV_DICT_APPEND AVDictionaryFlag = C.AV_DICT_APPEND

	// Allow to store several equal keys in the dictionary.
	AV_DICT_MULTIKEY AVDictionaryFlag = C.AV_DICT_MULTIKEY
)

////////////////////////////////////////////////////////////////////////////////
// JSON OUTPUT

func (ctx AVClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonAVClass{
		ClassName: C.GoString(ctx.class_name),
	})
}

func (ctx *AVDictionary) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonAVDictionary{
		Count: AVUtil_dict_count(ctx),
		Elems: AVUtil_dict_entries(ctx),
	})
}

func (ctx *AVDictionaryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonAVDictionaryEntry{
		Key:   ctx.Key(),
		Value: ctx.Value(),
	})
}

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (ctx AVClass) String() string {
	if str, err := json.MarshalIndent(ctx, "", "  "); err != nil {
		return err.Error()
	} else {
		return string(str)
	}
}

func (ctx *AVDictionary) String() string {
	if str, err := json.MarshalIndent(ctx, "", "  "); err != nil {
		return err.Error()
	} else {
		return string(str)
	}
}

func (v AVMediaType) String() string {
	switch v {
	case AVMEDIA_TYPE_UNKNOWN:
		return "AVMEDIA_TYPE_UNKNOWN"
	case AVMEDIA_TYPE_VIDEO:
		return "AVMEDIA_TYPE_VIDEO"
	case AVMEDIA_TYPE_AUDIO:
		return "AVMEDIA_TYPE_AUDIO"
	case AVMEDIA_TYPE_DATA:
		return "AVMEDIA_TYPE_DATA"
	case AVMEDIA_TYPE_SUBTITLE:
		return "AVMEDIA_TYPE_SUBTITLE"
	case AVMEDIA_TYPE_ATTACHMENT:
		return "AVMEDIA_TYPE_ATTACHMENT"
	}
	return "[AVMediaType]"
}
