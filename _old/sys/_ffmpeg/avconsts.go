package ffmpeg

////////////////////////////////////////////////////////////////////////////////
// CGO

/*
#cgo pkg-config: libavcodec
#include <libavutil/samplefmt.h>
#include <libavcodec/avcodec.h>
*/
import "C"
import "strings"

////////////////////////////////////////////////////////////////////////////////
// TYPES

type (
	AVCodecId       C.enum_AVCodecID
	AVMediaType     int
	AVCodecCap      uint32
	AVDisposition   int
	AVFormatFlag    int
	AVPacketFlag    int
	AVPixelFormat   C.enum_AVPixelFormat
	AVSampleFormat  C.enum_AVSampleFormat
	AVPictureType   int
	AVIOFlag        int
	AVLogLevel      int
	AVChannelLayout uint64
)

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	AVMEDIA_TYPE_VIDEO AVMediaType = iota
	AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA // Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT                  // Opaque data information usually sparse
	AVMEDIA_TYPE_UNKNOWN    AVMediaType = -1 // Usually treated as AVMEDIA_TYPE_DATA
)

const (
	AVIO_FLAG_NONE       AVIOFlag = 0
	AVIO_FLAG_READ       AVIOFlag = 1
	AVIO_FLAG_WRITE      AVIOFlag = 2
	AVIO_FLAG_READ_WRITE AVIOFlag = (AVIO_FLAG_READ | AVIO_FLAG_WRITE)
)

const (
	AV_PICTURE_TYPE_NONE AVPictureType = iota
	AV_PICTURE_TYPE_I                  ///< Intra
	AV_PICTURE_TYPE_P                  ///< Predicted
	AV_PICTURE_TYPE_B                  ///< Bi-dir predicted
	AV_PICTURE_TYPE_S                  ///< S(GMC)-VOP MPEG-4
	AV_PICTURE_TYPE_SI                 ///< Switching Intra
	AV_PICTURE_TYPE_SP                 ///< Switching Predicted
	AV_PICTURE_TYPE_BI                 ///< BI type
)

const (
	AV_LOG_QUIET   AVLogLevel = -8
	AV_LOG_PANIC   AVLogLevel = 0
	AV_LOG_FATAL   AVLogLevel = 8
	AV_LOG_ERROR   AVLogLevel = 16
	AV_LOG_WARNING AVLogLevel = 24
	AV_LOG_INFO    AVLogLevel = 32
	AV_LOG_VERBOSE AVLogLevel = 40
	AV_LOG_DEBUG   AVLogLevel = 48
	AV_LOG_TRACE   AVLogLevel = 56
)

const (
	AV_DISPOSITION_DEFAULT          AVDisposition = 0x0001
	AV_DISPOSITION_DUB              AVDisposition = 0x0002
	AV_DISPOSITION_ORIGINAL         AVDisposition = 0x0004
	AV_DISPOSITION_COMMENT          AVDisposition = 0x0008
	AV_DISPOSITION_LYRICS           AVDisposition = 0x0010
	AV_DISPOSITION_KARAOKE          AVDisposition = 0x0020
	AV_DISPOSITION_FORCED           AVDisposition = 0x0040
	AV_DISPOSITION_HEARING_IMPAIRED AVDisposition = 0x0080 // Stream for hearing impaired audiences
	AV_DISPOSITION_VISUAL_IMPAIRED  AVDisposition = 0x0100 // Stream for visual impaired audiences
	AV_DISPOSITION_CLEAN_EFFECTS    AVDisposition = 0x0200 // Stream without voice
	AV_DISPOSITION_ATTACHED_PIC     AVDisposition = 0x0400
	AV_DISPOSITION_TIMED_THUMBNAILS AVDisposition = 0x0800
	AV_DISPOSITION_CAPTIONS         AVDisposition = 0x10000
	AV_DISPOSITION_DESCRIPTIONS     AVDisposition = 0x20000
	AV_DISPOSITION_METADATA         AVDisposition = 0x40000
	AV_DISPOSITION_DEPENDENT        AVDisposition = 0x80000  // Dependent audio stream (mix_type=0 in mpegts)
	AV_DISPOSITION_STILL_IMAGE      AVDisposition = 0x100000 // Still images in video stream (still_picture_flag=1 in mpegts)
	AV_DISPOSITION_NONE             AVDisposition = 0
	AV_DISPOSITION_MIN                            = AV_DISPOSITION_DEFAULT
	AV_DISPOSITION_MAX                            = AV_DISPOSITION_STILL_IMAGE
)

const (
	AV_CODEC_CAP_DRAW_HORIZ_BAND     AVCodecCap = (1 << 0) // Decoder can use draw_horiz_band callback
	AV_CODEC_CAP_DR1                 AVCodecCap = (1 << 1) // Codec uses get_buffer() for allocating buffers and supports custom allocators
	AV_CODEC_CAP_TRUNCATED           AVCodecCap = (1 << 3)
	AV_CODEC_CAP_DELAY               AVCodecCap = (1 << 5)   // Encoder or decoder requires flushing with NULL input at the end in order to give the complete and correct output
	AV_CODEC_CAP_SMALL_LAST_FRAME    AVCodecCap = (1 << 6)   // Codec can be fed a final frame with a smaller size
	AV_CODEC_CAP_SUBFRAMES           AVCodecCap = (1 << 8)   // Codec can output multiple frames per AVPacket Normally demuxers return one frame at a time, demuxers which do not do are connected to a parser to split what they return into proper frames
	AV_CODEC_CAP_EXPERIMENTAL        AVCodecCap = (1 << 9)   // Codec is experimental and is thus avoided in favor of non experimental encoders
	AV_CODEC_CAP_CHANNEL_CONF        AVCodecCap = (1 << 10)  // Codec should fill in channel configuration and samplerate instead of container
	AV_CODEC_CAP_FRAME_THREADS       AVCodecCap = (1 << 12)  // Codec supports frame-level multithreading
	AV_CODEC_CAP_SLICE_THREADS       AVCodecCap = (1 << 13)  // Codec supports slice-based (or partition-based) multithreading
	AV_CODEC_CAP_PARAM_CHANGE        AVCodecCap = (1 << 14)  // Codec supports changed parameters at any point
	AV_CODEC_CAP_AUTO_THREADS        AVCodecCap = (1 << 15)  // Codec supports avctx->thread_count == 0 (auto)
	AV_CODEC_CAP_VARIABLE_FRAME_SIZE AVCodecCap = (1 << 16)  // Audio encoder supports receiving a different number of samples in each call
	AV_CODEC_CAP_AVOID_PROBING       AVCodecCap = (1 << 17)  // Decoder is not a preferred choice for probing
	AV_CODEC_CAP_HARDWARE            AVCodecCap = (1 << 18)  // Codec is backed by a hardware implementation
	AV_CODEC_CAP_HYBRID              AVCodecCap = (1 << 19)  // Codec is potentially backed by a hardware implementation, but not necessarily
	AV_CODEC_CAP_INTRA_ONLY          AVCodecCap = 0x40000000 // Codec is intra only
	AV_CODEC_CAP_LOSSLESS            AVCodecCap = 0x80000000 // Codec is lossless
	AV_CODEC_CAP_NONE                AVCodecCap = 0
	AV_CODEC_CAP_MIN                 AVCodecCap = AV_CODEC_CAP_DRAW_HORIZ_BAND
	AV_CODEC_CAP_MAX                 AVCodecCap = AV_CODEC_CAP_LOSSLESS
)

const (
	AVFMT_NOFILE        AVFormatFlag = 0x0001
	AVFMT_NEEDNUMBER    AVFormatFlag = 0x0002    // Needs '%d' in filename
	AVFMT_SHOW_IDS      AVFormatFlag = 0x0008    // Show format stream IDs numbers
	AVFMT_GLOBALHEADER  AVFormatFlag = 0x0040    // Format wants global header
	AVFMT_NOTIMESTAMPS  AVFormatFlag = 0x0080    // Format does not need / have any timestamps
	AVFMT_GENERIC_INDEX AVFormatFlag = 0x0100    // Use generic index building code
	AVFMT_TS_DISCONT    AVFormatFlag = 0x0200    // Format allows timestamp discontinuities. Note, muxers always require valid (monotone) timestamps
	AVFMT_VARIABLE_FPS  AVFormatFlag = 0x0400    // Format allows variable fps
	AVFMT_NODIMENSIONS  AVFormatFlag = 0x0800    // Format does not need width/height
	AVFMT_NOSTREAMS     AVFormatFlag = 0x1000    // Format does not require any streams
	AVFMT_NOBINSEARCH   AVFormatFlag = 0x2000    // Format does not allow to fall back on binary search via read_timestamp
	AVFMT_NOGENSEARCH   AVFormatFlag = 0x4000    // Format does not allow to fall back on generic search
	AVFMT_NO_BYTE_SEEK  AVFormatFlag = 0x8000    // Format does not allow seeking by bytes
	AVFMT_ALLOW_FLUSH   AVFormatFlag = 0x10000   // Format allows flushing. If not set, the muxer will not receive a NULL packet in the write_packet function
	AVFMT_TS_NONSTRICT  AVFormatFlag = 0x20000   // Format does not require strictly increasing timestamps, but they must still be monotonic
	AVFMT_TS_NEGATIVE   AVFormatFlag = 0x40000   // Format allows muxing negative timestamps. If not set the timestamp will be shifted in av_write_frame and av_interleaved_write_frame so they start from 0. The user or muxer can override this through AVFormatContext.avoid_negative_ts
	AVFMT_SEEK_TO_PTS   AVFormatFlag = 0x4000000 // Seeking is based on PTS
	AVFMT_NONE          AVFormatFlag = 0
	AVFMT_MIN                        = AVFMT_NOFILE
	AVFMT_MAX                        = AVFMT_SEEK_TO_PTS
)

const (
	AV_CH_FRONT_LEFT            AVChannelLayout = 0x00000001
	AV_CH_FRONT_RIGHT           AVChannelLayout = 0x00000002
	AV_CH_FRONT_CENTER          AVChannelLayout = 0x00000004
	AV_CH_LOW_FREQUENCY         AVChannelLayout = 0x00000008
	AV_CH_BACK_LEFT             AVChannelLayout = 0x00000010
	AV_CH_BACK_RIGHT            AVChannelLayout = 0x00000020
	AV_CH_FRONT_LEFT_OF_CENTER  AVChannelLayout = 0x00000040
	AV_CH_FRONT_RIGHT_OF_CENTER AVChannelLayout = 0x00000080
	AV_CH_BACK_CENTER           AVChannelLayout = 0x00000100
	AV_CH_SIDE_LEFT             AVChannelLayout = 0x00000200
	AV_CH_SIDE_RIGHT            AVChannelLayout = 0x00000400
	AV_CH_TOP_CENTER            AVChannelLayout = 0x00000800
	AV_CH_TOP_FRONT_LEFT        AVChannelLayout = 0x00001000
	AV_CH_TOP_FRONT_CENTER      AVChannelLayout = 0x00002000
	AV_CH_TOP_FRONT_RIGHT       AVChannelLayout = 0x00004000
	AV_CH_TOP_BACK_LEFT         AVChannelLayout = 0x00008000
	AV_CH_TOP_BACK_CENTER       AVChannelLayout = 0x00010000
	AV_CH_TOP_BACK_RIGHT        AVChannelLayout = 0x00020000
	AV_CH_STEREO_LEFT           AVChannelLayout = 0x20000000
	AV_CH_STEREO_RIGHT          AVChannelLayout = 0x40000000
	AV_CH_WIDE_LEFT             AVChannelLayout = 0x0000000080000000
	AV_CH_WIDE_RIGHT            AVChannelLayout = 0x0000000100000000
)

const (
	AV_CH_LAYOUT_NONE              AVChannelLayout = 0
	AV_CH_LAYOUT_MIN                               = AV_CH_FRONT_LEFT
	AV_CH_LAYOUT_MAX                               = AV_CH_STEREO_RIGHT
	AV_CH_LAYOUT_MONO                              = (AV_CH_FRONT_CENTER)
	AV_CH_LAYOUT_STEREO                            = (AV_CH_FRONT_LEFT | AV_CH_FRONT_RIGHT)
	AV_CH_LAYOUT_2POINT1                           = (AV_CH_LAYOUT_STEREO | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_2_1                               = (AV_CH_LAYOUT_STEREO | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_SURROUND                          = (AV_CH_LAYOUT_STEREO | AV_CH_FRONT_CENTER)
	AV_CH_LAYOUT_3POINT1                           = (AV_CH_LAYOUT_SURROUND | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_4POINT0                           = (AV_CH_LAYOUT_SURROUND | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_4POINT1                           = (AV_CH_LAYOUT_4POINT0 | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_2_2                               = (AV_CH_LAYOUT_STEREO | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT)
	AV_CH_LAYOUT_QUAD                              = (AV_CH_LAYOUT_STEREO | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
	AV_CH_LAYOUT_5POINT0                           = (AV_CH_LAYOUT_SURROUND | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT)
	AV_CH_LAYOUT_5POINT1                           = (AV_CH_LAYOUT_5POINT0 | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_5POINT0_BACK                      = (AV_CH_LAYOUT_SURROUND | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
	AV_CH_LAYOUT_5POINT1_BACK                      = (AV_CH_LAYOUT_5POINT0_BACK | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_6POINT0                           = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_6POINT0_FRONT                     = (AV_CH_LAYOUT_2_2 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_LAYOUT_HEXAGONAL                         = (AV_CH_LAYOUT_5POINT0_BACK | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_6POINT1                           = (AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_6POINT1_BACK                      = (AV_CH_LAYOUT_5POINT1_BACK | AV_CH_BACK_CENTER)
	AV_CH_LAYOUT_6POINT1_FRONT                     = (AV_CH_LAYOUT_6POINT0_FRONT | AV_CH_LOW_FREQUENCY)
	AV_CH_LAYOUT_7POINT0                           = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
	AV_CH_LAYOUT_7POINT0_FRONT                     = (AV_CH_LAYOUT_5POINT0 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_LAYOUT_7POINT1                           = (AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
	AV_CH_LAYOUT_7POINT1_WIDE                      = (AV_CH_LAYOUT_5POINT1 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_LAYOUT_7POINT1_WIDE_BACK                 = (AV_CH_LAYOUT_5POINT1_BACK | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
	AV_CH_LAYOUT_OCTAGONAL                         = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_CENTER | AV_CH_BACK_RIGHT)
	AV_CH_LAYOUT_HEXADECAGONAL                     = (AV_CH_LAYOUT_OCTAGONAL | AV_CH_WIDE_LEFT | AV_CH_WIDE_RIGHT | AV_CH_TOP_BACK_LEFT | AV_CH_TOP_BACK_RIGHT | AV_CH_TOP_BACK_CENTER | AV_CH_TOP_FRONT_CENTER | AV_CH_TOP_FRONT_LEFT | AV_CH_TOP_FRONT_RIGHT)
	AV_CH_LAYOUT_STEREO_DOWNMIX                    = (AV_CH_STEREO_LEFT | AV_CH_STEREO_RIGHT)
)

const (
	AV_CODEC_ID_NONE AVCodecId = iota
	AV_CODEC_ID_MPEG1VIDEO
	AV_CODEC_ID_MPEG2VIDEO ///< preferred ID for MPEG-1/2 video decoding
	AV_CODEC_ID_H261
	AV_CODEC_ID_H263
	AV_CODEC_ID_RV10
	AV_CODEC_ID_RV20
	AV_CODEC_ID_MJPEG
	AV_CODEC_ID_MJPEGB
	AV_CODEC_ID_LJPEG
	AV_CODEC_ID_SP5X
	AV_CODEC_ID_JPEGLS
	AV_CODEC_ID_MPEG4
	AV_CODEC_ID_RAWVIDEO
	AV_CODEC_ID_MSMPEG4V1
	AV_CODEC_ID_MSMPEG4V2
	AV_CODEC_ID_MSMPEG4V3
	AV_CODEC_ID_WMV1
	AV_CODEC_ID_WMV2
	AV_CODEC_ID_H263P
	AV_CODEC_ID_H263I
	AV_CODEC_ID_FLV1
	AV_CODEC_ID_SVQ1
	AV_CODEC_ID_SVQ3
	AV_CODEC_ID_DVVIDEO
	AV_CODEC_ID_HUFFYUV
	AV_CODEC_ID_CYUV
	AV_CODEC_ID_H264
	AV_CODEC_ID_INDEO3
	AV_CODEC_ID_VP3
	AV_CODEC_ID_THEORA
	AV_CODEC_ID_ASV1
	AV_CODEC_ID_ASV2
	AV_CODEC_ID_FFV1
	AV_CODEC_ID_4XM
	AV_CODEC_ID_VCR1
	AV_CODEC_ID_CLJR
	AV_CODEC_ID_MDEC
	AV_CODEC_ID_ROQ
	AV_CODEC_ID_INTERPLAY_VIDEO
	AV_CODEC_ID_XAN_WC3
	AV_CODEC_ID_XAN_WC4
	AV_CODEC_ID_RPZA
	AV_CODEC_ID_CINEPAK
	AV_CODEC_ID_WS_VQA
	AV_CODEC_ID_MSRLE
	AV_CODEC_ID_MSVIDEO1
	AV_CODEC_ID_IDCIN
	AV_CODEC_ID_8BPS
	AV_CODEC_ID_SMC
	AV_CODEC_ID_FLIC
	AV_CODEC_ID_TRUEMOTION1
	AV_CODEC_ID_VMDVIDEO
	AV_CODEC_ID_MSZH
	AV_CODEC_ID_ZLIB
	AV_CODEC_ID_QTRLE
	AV_CODEC_ID_TSCC
	AV_CODEC_ID_ULTI
	AV_CODEC_ID_QDRAW
	AV_CODEC_ID_VIXL
	AV_CODEC_ID_QPEG
	AV_CODEC_ID_PNG
	AV_CODEC_ID_PPM
	AV_CODEC_ID_PBM
	AV_CODEC_ID_PGM
	AV_CODEC_ID_PGMYUV
	AV_CODEC_ID_PAM
	AV_CODEC_ID_FFVHUFF
	AV_CODEC_ID_RV30
	AV_CODEC_ID_RV40
	AV_CODEC_ID_VC1
	AV_CODEC_ID_WMV3
	AV_CODEC_ID_LOCO
	AV_CODEC_ID_WNV1
	AV_CODEC_ID_AASC
	AV_CODEC_ID_INDEO2
	AV_CODEC_ID_FRAPS
	AV_CODEC_ID_TRUEMOTION2
	AV_CODEC_ID_BMP
	AV_CODEC_ID_CSCD
	AV_CODEC_ID_MMVIDEO
	AV_CODEC_ID_ZMBV
	AV_CODEC_ID_AVS
	AV_CODEC_ID_SMACKVIDEO
	AV_CODEC_ID_NUV
	AV_CODEC_ID_KMVC
	AV_CODEC_ID_FLASHSV
	AV_CODEC_ID_CAVS
	AV_CODEC_ID_JPEG2000
	AV_CODEC_ID_VMNC
	AV_CODEC_ID_VP5
	AV_CODEC_ID_VP6
	AV_CODEC_ID_VP6F
	AV_CODEC_ID_TARGA
	AV_CODEC_ID_DSICINVIDEO
	AV_CODEC_ID_TIERTEXSEQVIDEO
	AV_CODEC_ID_TIFF
	AV_CODEC_ID_GIF
	AV_CODEC_ID_DXA
	AV_CODEC_ID_DNXHD
	AV_CODEC_ID_THP
	AV_CODEC_ID_SGI
	AV_CODEC_ID_C93
	AV_CODEC_ID_BETHSOFTVID
	AV_CODEC_ID_PTX
	AV_CODEC_ID_TXD
	AV_CODEC_ID_VP6A
	AV_CODEC_ID_AMV
	AV_CODEC_ID_VB
	AV_CODEC_ID_PCX
	AV_CODEC_ID_SUNRAST
	AV_CODEC_ID_INDEO4
	AV_CODEC_ID_INDEO5
	AV_CODEC_ID_MIMIC
	AV_CODEC_ID_RL2
	AV_CODEC_ID_ESCAPE124
	AV_CODEC_ID_DIRAC
	AV_CODEC_ID_BFI
	AV_CODEC_ID_CMV
	AV_CODEC_ID_MOTIONPIXELS
	AV_CODEC_ID_TGV
	AV_CODEC_ID_TGQ
	AV_CODEC_ID_TQI
	AV_CODEC_ID_AURA
	AV_CODEC_ID_AURA2
	AV_CODEC_ID_V210X
	AV_CODEC_ID_TMV
	AV_CODEC_ID_V210
	AV_CODEC_ID_DPX
	AV_CODEC_ID_MAD
	AV_CODEC_ID_FRWU
	AV_CODEC_ID_FLASHSV2
	AV_CODEC_ID_CDGRAPHICS
	AV_CODEC_ID_R210
	AV_CODEC_ID_ANM
	AV_CODEC_ID_BINKVIDEO
	AV_CODEC_ID_IFF_ILBM
	AV_CODEC_ID_KGV1
	AV_CODEC_ID_YOP
	AV_CODEC_ID_VP8
	AV_CODEC_ID_PICTOR
	AV_CODEC_ID_ANSI
	AV_CODEC_ID_A64_MULTI
	AV_CODEC_ID_A64_MULTI5
	AV_CODEC_ID_R10K
	AV_CODEC_ID_MXPEG
	AV_CODEC_ID_LAGARITH
	AV_CODEC_ID_PRORES
	AV_CODEC_ID_JV
	AV_CODEC_ID_DFA
	AV_CODEC_ID_WMV3IMAGE
	AV_CODEC_ID_VC1IMAGE
	AV_CODEC_ID_UTVIDEO
	AV_CODEC_ID_BMV_VIDEO
	AV_CODEC_ID_VBLE
	AV_CODEC_ID_DXTORY
	AV_CODEC_ID_V410
	AV_CODEC_ID_XWD
	AV_CODEC_ID_CDXL
	AV_CODEC_ID_XBM
	AV_CODEC_ID_ZEROCODEC
	AV_CODEC_ID_MSS1
	AV_CODEC_ID_MSA1
	AV_CODEC_ID_TSCC2
	AV_CODEC_ID_MTS2
	AV_CODEC_ID_CLLC
	AV_CODEC_ID_MSS2
	AV_CODEC_ID_VP9
	AV_CODEC_ID_AIC
	AV_CODEC_ID_ESCAPE130
	AV_CODEC_ID_G2M
	AV_CODEC_ID_WEBP
	AV_CODEC_ID_HNM4_VIDEO
	AV_CODEC_ID_HEVC
	AV_CODEC_ID_FIC
	AV_CODEC_ID_ALIAS_PIX
	AV_CODEC_ID_BRENDER_PIX
	AV_CODEC_ID_PAF_VIDEO
	AV_CODEC_ID_EXR
	AV_CODEC_ID_VP7
	AV_CODEC_ID_SANM
	AV_CODEC_ID_SGIRLE
	AV_CODEC_ID_MVC1
	AV_CODEC_ID_MVC2
	AV_CODEC_ID_HQX
	AV_CODEC_ID_TDSC
	AV_CODEC_ID_HQ_HQA
	AV_CODEC_ID_HAP
	AV_CODEC_ID_DDS
	AV_CODEC_ID_DXV
	AV_CODEC_ID_SCREENPRESSO
	AV_CODEC_ID_RSCC
	AV_CODEC_ID_AVS2
	AV_CODEC_ID_H265         = AV_CODEC_ID_HEVC
	AV_CODEC_ID_IFF_BYTERUN1 = AV_CODEC_ID_IFF_ILBM
)

const (
	AV_CODEC_ID_Y41P AVCodecId = iota + 0x8000
	AV_CODEC_ID_AVRP
	AV_CODEC_ID_012V
	AV_CODEC_ID_AVUI
	AV_CODEC_ID_AYUV
	AV_CODEC_ID_TARGA_Y216
	AV_CODEC_ID_V308
	AV_CODEC_ID_V408
	AV_CODEC_ID_YUV4
	AV_CODEC_ID_AVRN
	AV_CODEC_ID_CPIA
	AV_CODEC_ID_XFACE
	AV_CODEC_ID_SNOW
	AV_CODEC_ID_SMVJPEG
	AV_CODEC_ID_APNG
	AV_CODEC_ID_DAALA
	AV_CODEC_ID_CFHD
	AV_CODEC_ID_TRUEMOTION2RT
	AV_CODEC_ID_M101
	AV_CODEC_ID_MAGICYUV
	AV_CODEC_ID_SHEERVIDEO
	AV_CODEC_ID_YLC
	AV_CODEC_ID_PSD
	AV_CODEC_ID_PIXLET
	AV_CODEC_ID_SPEEDHQ
	AV_CODEC_ID_FMVC
	AV_CODEC_ID_SCPR
	AV_CODEC_ID_CLEARVIDEO
	AV_CODEC_ID_XPM
	AV_CODEC_ID_AV1
	AV_CODEC_ID_BITPACKED
	AV_CODEC_ID_MSCC
	AV_CODEC_ID_SRGC
	AV_CODEC_ID_SVG
	AV_CODEC_ID_GDV
	AV_CODEC_ID_FITS
	AV_CODEC_ID_IMM4
	AV_CODEC_ID_PROSUMER
	AV_CODEC_ID_MWSC
	AV_CODEC_ID_WCMV
	AV_CODEC_ID_RASC
)

const (
	// Audio Codecs
	AV_CODEC_ID_FIRST_AUDIO           = AV_CODEC_ID_PCM_S16LE
	AV_CODEC_ID_PCM_S16LE   AVCodecId = iota + 0x10000
	AV_CODEC_ID_PCM_S16BE
	AV_CODEC_ID_PCM_U16LE
	AV_CODEC_ID_PCM_U16BE
	AV_CODEC_ID_PCM_S8
	AV_CODEC_ID_PCM_U8
	AV_CODEC_ID_PCM_MULAW
	AV_CODEC_ID_PCM_ALAW
	AV_CODEC_ID_PCM_S32LE
	AV_CODEC_ID_PCM_S32BE
	AV_CODEC_ID_PCM_U32LE
	AV_CODEC_ID_PCM_U32BE
	AV_CODEC_ID_PCM_S24LE
	AV_CODEC_ID_PCM_S24BE
	AV_CODEC_ID_PCM_U24LE
	AV_CODEC_ID_PCM_U24BE
	AV_CODEC_ID_PCM_S24DAUD
	AV_CODEC_ID_PCM_ZORK
	AV_CODEC_ID_PCM_S16LE_PLANAR
	AV_CODEC_ID_PCM_DVD
	AV_CODEC_ID_PCM_F32BE
	AV_CODEC_ID_PCM_F32LE
	AV_CODEC_ID_PCM_F64BE
	AV_CODEC_ID_PCM_F64LE
	AV_CODEC_ID_PCM_BLURAY
	AV_CODEC_ID_PCM_LXF
	AV_CODEC_ID_S302M
	AV_CODEC_ID_PCM_S8_PLANAR
	AV_CODEC_ID_PCM_S24LE_PLANAR
	AV_CODEC_ID_PCM_S32LE_PLANAR
	AV_CODEC_ID_PCM_S16BE_PLANAR
)
const (
	AV_CODEC_ID_PCM_S64LE AVCodecId = iota + 0x10800
	AV_CODEC_ID_PCM_S64BE
	AV_CODEC_ID_PCM_F16LE
	AV_CODEC_ID_PCM_F24LE
	AV_CODEC_ID_PCM_VIDC
)

const (
	AV_CODEC_ID_ADPCM_IMA_QT AVCodecId = iota + 0x11000
	AV_CODEC_ID_ADPCM_IMA_WAV
	AV_CODEC_ID_ADPCM_IMA_DK3
	AV_CODEC_ID_ADPCM_IMA_DK4
	AV_CODEC_ID_ADPCM_IMA_WS
	AV_CODEC_ID_ADPCM_IMA_SMJPEG
	AV_CODEC_ID_ADPCM_MS
	AV_CODEC_ID_ADPCM_4XM
	AV_CODEC_ID_ADPCM_XA
	AV_CODEC_ID_ADPCM_ADX
	AV_CODEC_ID_ADPCM_EA
	AV_CODEC_ID_ADPCM_G726
	AV_CODEC_ID_ADPCM_CT
	AV_CODEC_ID_ADPCM_SWF
	AV_CODEC_ID_ADPCM_YAMAHA
	AV_CODEC_ID_ADPCM_SBPRO_4
	AV_CODEC_ID_ADPCM_SBPRO_3
	AV_CODEC_ID_ADPCM_SBPRO_2
	AV_CODEC_ID_ADPCM_THP
	AV_CODEC_ID_ADPCM_IMA_AMV
	AV_CODEC_ID_ADPCM_EA_R1
	AV_CODEC_ID_ADPCM_EA_R3
	AV_CODEC_ID_ADPCM_EA_R2
	AV_CODEC_ID_ADPCM_IMA_EA_SEAD
	AV_CODEC_ID_ADPCM_IMA_EA_EACS
	AV_CODEC_ID_ADPCM_EA_XAS
	AV_CODEC_ID_ADPCM_EA_MAXIS_XA
	AV_CODEC_ID_ADPCM_IMA_ISS
	AV_CODEC_ID_ADPCM_G722
	AV_CODEC_ID_ADPCM_IMA_APC
	AV_CODEC_ID_ADPCM_VIMA
)

const (
	AV_CODEC_ID_ADPCM_AFC AVCodecId = iota + 0x11800
	AV_CODEC_ID_ADPCM_IMA_OKI
	AV_CODEC_ID_ADPCM_DTK
	AV_CODEC_ID_ADPCM_IMA_RAD
	AV_CODEC_ID_ADPCM_G726LE
	AV_CODEC_ID_ADPCM_THP_LE
	AV_CODEC_ID_ADPCM_PSX
	AV_CODEC_ID_ADPCM_AICA
	AV_CODEC_ID_ADPCM_IMA_DAT4
	AV_CODEC_ID_ADPCM_MTAF
)

const (
	AV_CODEC_ID_AMR_NB AVCodecId = iota + 0x12000
	AV_CODEC_ID_AMR_WB
)

const (
	AV_CODEC_ID_RA_144 AVCodecId = iota + 0x13000
	AV_CODEC_ID_RA_288
)

const (
	AV_CODEC_ID_ROQ_DPCM AVCodecId = iota + 0x14000
	AV_CODEC_ID_INTERPLAY_DPCM
	AV_CODEC_ID_XAN_DPCM
	AV_CODEC_ID_SOL_DPCM
)

const (
	AV_CODEC_ID_SDX2_DPCM AVCodecId = iota + 0x14800
	AV_CODEC_ID_GREMLIN_DPCM
)
const (
	AV_CODEC_ID_MP2 AVCodecId = iota + 0x15000
	AV_CODEC_ID_MP3           ///< preferred ID for decoding MPEG audio layer 1, 2 or 3
	AV_CODEC_ID_AAC
	AV_CODEC_ID_AC3
	AV_CODEC_ID_DTS
	AV_CODEC_ID_VORBIS
	AV_CODEC_ID_DVAUDIO
	AV_CODEC_ID_WMAV1
	AV_CODEC_ID_WMAV2
	AV_CODEC_ID_MACE3
	AV_CODEC_ID_MACE6
	AV_CODEC_ID_VMDAUDIO
	AV_CODEC_ID_FLAC
	AV_CODEC_ID_MP3ADU
	AV_CODEC_ID_MP3ON4
	AV_CODEC_ID_SHORTEN
	AV_CODEC_ID_ALAC
	AV_CODEC_ID_WESTWOOD_SND1
	AV_CODEC_ID_GSM ///< as in Berlin toast format
	AV_CODEC_ID_QDM2
	AV_CODEC_ID_COOK
	AV_CODEC_ID_TRUESPEECH
	AV_CODEC_ID_TTA
	AV_CODEC_ID_SMACKAUDIO
	AV_CODEC_ID_QCELP
	AV_CODEC_ID_WAVPACK
	AV_CODEC_ID_DSICINAUDIO
	AV_CODEC_ID_IMC
	AV_CODEC_ID_MUSEPACK7
	AV_CODEC_ID_MLP
	AV_CODEC_ID_GSM_MS /* as found in WAV */
	AV_CODEC_ID_ATRAC3
	AV_CODEC_ID_APE
	AV_CODEC_ID_NELLYMOSER
	AV_CODEC_ID_MUSEPACK8
	AV_CODEC_ID_SPEEX
	AV_CODEC_ID_WMAVOICE
	AV_CODEC_ID_WMAPRO
	AV_CODEC_ID_WMALOSSLESS
	AV_CODEC_ID_ATRAC3P
	AV_CODEC_ID_EAC3
	AV_CODEC_ID_SIPR
	AV_CODEC_ID_MP1
	AV_CODEC_ID_TWINVQ
	AV_CODEC_ID_TRUEHD
	AV_CODEC_ID_MP4ALS
	AV_CODEC_ID_ATRAC1
	AV_CODEC_ID_BINKAUDIO_RDFT
	AV_CODEC_ID_BINKAUDIO_DCT
	AV_CODEC_ID_AAC_LATM
	AV_CODEC_ID_QDMC
	AV_CODEC_ID_CELT
	AV_CODEC_ID_G723_1
	AV_CODEC_ID_G729
	AV_CODEC_ID_8SVX_EXP
	AV_CODEC_ID_8SVX_FIB
	AV_CODEC_ID_BMV_AUDIO
	AV_CODEC_ID_RALF
	AV_CODEC_ID_IAC
	AV_CODEC_ID_ILBC
	AV_CODEC_ID_OPUS
	AV_CODEC_ID_COMFORT_NOISE
	AV_CODEC_ID_TAK
	AV_CODEC_ID_METASOUND
	AV_CODEC_ID_PAF_AUDIO
	AV_CODEC_ID_ON2AVC
	AV_CODEC_ID_DSS_SP
	AV_CODEC_ID_CODEC2
)

const (
	AV_CODEC_ID_FFWAVESYNTH AVCodecId = iota + 0x15800
	AV_CODEC_ID_SONIC
	AV_CODEC_ID_SONIC_LS
	AV_CODEC_ID_EVRC
	AV_CODEC_ID_SMV
	AV_CODEC_ID_DSD_LSBF
	AV_CODEC_ID_DSD_MSBF
	AV_CODEC_ID_DSD_LSBF_PLANAR
	AV_CODEC_ID_DSD_MSBF_PLANAR
	AV_CODEC_ID_4GV
	AV_CODEC_ID_INTERPLAY_ACM
	AV_CODEC_ID_XMA1
	AV_CODEC_ID_XMA2
	AV_CODEC_ID_DST
	AV_CODEC_ID_ATRAC3AL
	AV_CODEC_ID_ATRAC3PAL
	AV_CODEC_ID_DOLBY_E
	AV_CODEC_ID_APTX
	AV_CODEC_ID_APTX_HD
	AV_CODEC_ID_SBC
	AV_CODEC_ID_ATRAC9
)

const (
	AV_CODEC_ID_FIRST_SUBTITLE           = AV_CODEC_ID_DVD_SUBTITLE
	AV_CODEC_ID_DVD_SUBTITLE   AVCodecId = iota + 0x17000
	AV_CODEC_ID_DVB_SUBTITLE
	AV_CODEC_ID_TEXT ///< raw UTF-8 text
	AV_CODEC_ID_XSUB
	AV_CODEC_ID_SSA
	AV_CODEC_ID_MOV_TEXT
	AV_CODEC_ID_HDMV_PGS_SUBTITLE
	AV_CODEC_ID_DVB_TELETEXT
	AV_CODEC_ID_SRT
)

const (
	AV_CODEC_ID_MICRODVD AVCodecId = iota + 0x17800
	AV_CODEC_ID_EIA_608
	AV_CODEC_ID_JACOSUB
	AV_CODEC_ID_SAMI
	AV_CODEC_ID_REALTEXT
	AV_CODEC_ID_STL
	AV_CODEC_ID_SUBVIEWER1
	AV_CODEC_ID_SUBVIEWER
	AV_CODEC_ID_SUBRIP
	AV_CODEC_ID_WEBVTT
	AV_CODEC_ID_MPL2
	AV_CODEC_ID_VPLAYER
	AV_CODEC_ID_PJS
	AV_CODEC_ID_ASS
	AV_CODEC_ID_HDMV_TEXT_SUBTITLE
	AV_CODEC_ID_TTML
)

const (
	AV_CODEC_ID_FIRST_UNKNOWN           = AV_CODEC_ID_TTF
	AV_CODEC_ID_TTF           AVCodecId = iota + 0x18000
	AV_CODEC_ID_SCTE_35                 ///< Contain timestamp estimated through PCR of program stream.
)

const (
	AV_CODEC_ID_BINTEXT AVCodecId = iota + 0x18800
	AV_CODEC_ID_XBIN
	AV_CODEC_ID_IDF
	AV_CODEC_ID_OTF
	AV_CODEC_ID_SMPTE_KLV
	AV_CODEC_ID_DVD_NAV
	AV_CODEC_ID_TIMED_ID3
	AV_CODEC_ID_BIN_DATA

	AV_CODEC_ID_PROBE           AVCodecId = 0x19000 ///< codec_id is not known (like AV_CODEC_ID_NONE) but lavf should attempt to identify it
	AV_CODEC_ID_MPEG2TS         AVCodecId = 0x20000 // _FAKE_ codec to indicate a raw MPEG-2 TS
	AV_CODEC_ID_MPEG4SYSTEMS    AVCodecId = 0x20001 // _FAKE_ codec to indicate a MPEG-4 Systems
	AV_CODEC_ID_FFMETADATA      AVCodecId = 0x21000 // Dummy codec for streams containing only metadata information.
	AV_CODEC_ID_WRAPPED_AVFRAME AVCodecId = 0x21001 // Passthrough codec, AVFrames wrapped in AVPacket
)

const (
	AV_CODEC_ID_MIN = AV_CODEC_ID_MPEG1VIDEO
	AV_CODEC_ID_MAX = AV_CODEC_ID_WRAPPED_AVFRAME
)

const (
	AV_PKT_FLAG_NONE       AVPacketFlag = 0
	AV_PKT_FLAG_KEY        AVPacketFlag = 0x0001 // The packet contains a keyframe
	AV_PKT_FLAG_CORRUPT    AVPacketFlag = 0x0002 // The packet content is corrupted
	AV_PKT_FLAG_DISCARD    AVPacketFlag = 0x0004 // Flag is used to discard packets
	AV_PKT_FLAG_TRUSTED    AVPacketFlag = 0x0008 // The packet comes from a trusted source
	AV_PKT_FLAG_DISPOSABLE AVPacketFlag = 0x0010 // The packet contains frames that can be discarded
	AV_PKT_FLAG_MIN                     = AV_PKT_FLAG_KEY
	AV_PKT_FLAG_MAX                     = AV_PKT_FLAG_DISPOSABLE
)

const (
	AV_PIX_FMT_YUV420P        AVPixelFormat = iota // planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AV_PIX_FMT_YUYV422                             // packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AV_PIX_FMT_RGB24                               // packed RGB 8:8:8, 24bpp, RGBRGB...
	AV_PIX_FMT_BGR24                               // packed RGB 8:8:8, 24bpp, BGRBGR...
	AV_PIX_FMT_YUV422P                             // planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_YUV444P                             // planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AV_PIX_FMT_YUV410P                             // planar YUV 4:1:0, 9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AV_PIX_FMT_YUV411P                             // planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	AV_PIX_FMT_GRAY8                               // 8bpp.
	AV_PIX_FMT_MONOWHITE                           // 1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb.
	AV_PIX_FMT_MONOBLACK                           // 1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb.
	AV_PIX_FMT_PAL8                                // 8 bits with AV_PIX_FMT_RGB32 palette
	AV_PIX_FMT_YUVJ420P                            // planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV420P and setting color_range
	AV_PIX_FMT_YUVJ422P                            // planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV422P and setting color_range
	AV_PIX_FMT_YUVJ444P                            // planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV444P and setting color_range
	AV_PIX_FMT_UYVY422                             // packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	AV_PIX_FMT_UYYVYY411                           // packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	AV_PIX_FMT_BGR8                                // packed RGB 3:3:2, 8bpp, (msb)2B 3G 3R(lsb)
	AV_PIX_FMT_BGR4                                // packed RGB 1:2:1 bitstream, 4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_BGR4_BYTE                           // packed RGB 1:2:1, 8bpp, (msb)1B 2G 1R(lsb)
	AV_PIX_FMT_RGB8                                // packed RGB 3:3:2, 8bpp, (msb)2R 3G 3B(lsb)
	AV_PIX_FMT_RGB4                                // packed RGB 1:2:1 bitstream, 4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_RGB4_BYTE                           // packed RGB 1:2:1, 8bpp, (msb)1R 2G 1B(lsb)
	AV_PIX_FMT_NV12                                // planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AV_PIX_FMT_NV21                                // as above, but U and V bytes are swapped
	AV_PIX_FMT_ARGB                                // packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	AV_PIX_FMT_RGBA                                // packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	AV_PIX_FMT_ABGR                                // packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	AV_PIX_FMT_BGRA                                // packed BGRA 8:8:8:8, 32bpp, BGRABGRA...
	AV_PIX_FMT_GRAY16BE                            // 16bpp, big-endian.
	AV_PIX_FMT_GRAY16LE                            // 16bpp, little-endian.
	AV_PIX_FMT_YUV440P                             // planar YUV 4:4:0 (1 Cr & Cb sample per 1x2 Y samples)
	AV_PIX_FMT_YUVJ440P                            // planar YUV 4:4:0 full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV440P and setting color_range
	AV_PIX_FMT_YUVA420P                            // planar YUV 4:2:0, 20bpp, (1 Cr & Cb sample per 2x2 Y & A samples)
	AV_PIX_FMT_RGB48BE                             // packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_RGB48LE                             // packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as little-endian
	AV_PIX_FMT_RGB565BE                            // packed RGB 5:6:5, 16bpp, (msb) 5R 6G 5B(lsb), big-endian
	AV_PIX_FMT_RGB565LE                            // packed RGB 5:6:5, 16bpp, (msb) 5R 6G 5B(lsb), little-endian
	AV_PIX_FMT_RGB555BE                            // packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), big-endian , X=unused/undefined
	AV_PIX_FMT_RGB555LE                            // packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_BGR565BE                            // packed BGR 5:6:5, 16bpp, (msb) 5B 6G 5R(lsb), big-endian
	AV_PIX_FMT_BGR565LE                            // packed BGR 5:6:5, 16bpp, (msb) 5B 6G 5R(lsb), little-endian
	AV_PIX_FMT_BGR555BE                            // packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), big-endian , X=unused/undefined
	AV_PIX_FMT_BGR555LE                            // packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_VAAPI_MOCO                          // HW acceleration through VA API at motion compensation entry-point, Picture.data[3] contains a vaapi_render_state struct which contains macroblocks as well as various fields extracted from headers.
	AV_PIX_FMT_VAAPI_IDCT                          // HW acceleration through VA API at IDCT entry-point, Picture.data[3] contains a vaapi_render_state struct which contains fields extracted from headers.
	AV_PIX_FMT_VAAPI_VLD                           // HW decoding through VA API, Picture.data[3] contains a VASurfaceID.
	AV_PIX_FMT_VAAPI                               //
	AV_PIX_FMT_YUV420P16LE                         // planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P16BE                         // planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV422P16LE                         // planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P16BE                         // planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P16LE                         // planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P16BE                         // planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_DXVA2_VLD                           // HW decoding through DXVA2, Picture.data[3] contains a LPDIRECT3DSURFACE9 pointer.
	AV_PIX_FMT_RGB444LE                            // packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_RGB444BE                            // packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), big-endian, X=unused/undefined
	AV_PIX_FMT_BGR444LE                            // packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_BGR444BE                            // packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), big-endian, X=unused/undefined
	AV_PIX_FMT_YA8                                 // 8 bits gray, 8 bits alpha
	AV_PIX_FMT_Y400A                               // alias for AV_PIX_FMT_YA8
	AV_PIX_FMT_GRAY8A                              // alias for AV_PIX_FMT_YA8
	AV_PIX_FMT_BGR48BE                             // packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_BGR48LE                             // packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as little-endian
	AV_PIX_FMT_YUV420P9BE                          // The following 12 formats have the disadvantage of needing 1 format for each bit depth.
	AV_PIX_FMT_YUV420P9LE                          // planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P10BE                         // planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P10LE                         // planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P10BE                         // planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P10LE                         // planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P9BE                          // planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P9LE                          // planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P10BE                         // planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P10LE                         // planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P9BE                          // planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P9LE                          // planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_GBRP                                //
	AV_PIX_FMT_GBR24P                              // planar GBR 4:4:4 24bpp
	AV_PIX_FMT_GBRP9BE                             // planar GBR 4:4:4 27bpp, big-endian
	AV_PIX_FMT_GBRP9LE                             // planar GBR 4:4:4 27bpp, little-endian
	AV_PIX_FMT_GBRP10BE                            // planar GBR 4:4:4 30bpp, big-endian
	AV_PIX_FMT_GBRP10LE                            // planar GBR 4:4:4 30bpp, little-endian
	AV_PIX_FMT_GBRP16BE                            // planar GBR 4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRP16LE                            // planar GBR 4:4:4 48bpp, little-endian
	AV_PIX_FMT_YUVA422P                            // planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
	AV_PIX_FMT_YUVA444P                            // planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	AV_PIX_FMT_YUVA420P9BE                         // planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), big-endian
	AV_PIX_FMT_YUVA420P9LE                         // planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), little-endian
	AV_PIX_FMT_YUVA422P9BE                         // planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA422P9LE                         // planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA444P9BE                         // planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA444P9LE                         // planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA420P10BE                        // planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P10LE                        // planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P10BE                        // planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P10LE                        // planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P10BE                        // planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P10LE                        // planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA420P16BE                        // planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P16LE                        // planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P16BE                        // planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P16LE                        // planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P16BE                        // planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P16LE                        // planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AV_PIX_FMT_VDPAU                               // HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface.
	AV_PIX_FMT_XYZ12LE                             // packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as little-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_XYZ12BE                             // packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as big-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_NV16                                // interleaved chroma YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_NV20LE                              // interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_NV20BE                              // interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_RGBA64BE                            // packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_RGBA64LE                            // packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AV_PIX_FMT_BGRA64BE                            // packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_BGRA64LE                            // packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AV_PIX_FMT_YVYU422                             // packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb
	AV_PIX_FMT_YA16BE                              // 16 bits gray, 16 bits alpha (big-endian)
	AV_PIX_FMT_YA16LE                              // 16 bits gray, 16 bits alpha (little-endian)
	AV_PIX_FMT_GBRAP                               // planar GBRA 4:4:4:4 32bpp
	AV_PIX_FMT_GBRAP16BE                           // planar GBRA 4:4:4:4 64bpp, big-endian
	AV_PIX_FMT_GBRAP16LE                           // planar GBRA 4:4:4:4 64bpp, little-endian
	AV_PIX_FMT_QSV                                 // HW acceleration through QSV, data[3] contains a pointer to the mfxFrameSurface1 structure.
	AV_PIX_FMT_MMAL                                // HW acceleration though MMAL, data[3] contains a pointer to the MMAL_BUFFER_HEADER_T structure.
	AV_PIX_FMT_D3D11VA_VLD                         // HW decoding through Direct3D11 via old API, Picture.data[3] contains a ID3D11VideoDecoderOutputView pointer.
	AV_PIX_FMT_CUDA                                // HW acceleration through CUDA.
	AV_PIX_FMT_0RGB                                // packed RGB 8:8:8, 32bpp, XRGBXRGB... X=unused/undefined
	AV_PIX_FMT_RGB0                                // packed RGB 8:8:8, 32bpp, RGBXRGBX... X=unused/undefined
	AV_PIX_FMT_0BGR                                // packed BGR 8:8:8, 32bpp, XBGRXBGR... X=unused/undefined
	AV_PIX_FMT_BGR0                                // packed BGR 8:8:8, 32bpp, BGRXBGRX... X=unused/undefined
	AV_PIX_FMT_YUV420P12BE                         // planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P12LE                         // planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P14BE                         // planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P14LE                         // planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P12BE                         // planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P12LE                         // planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P14BE                         // planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P14LE                         // planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P12BE                         // planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P12LE                         // planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P14BE                         // planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P14LE                         // planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_GBRP12BE                            // planar GBR 4:4:4 36bpp, big-endian
	AV_PIX_FMT_GBRP12LE                            // planar GBR 4:4:4 36bpp, little-endian
	AV_PIX_FMT_GBRP14BE                            // planar GBR 4:4:4 42bpp, big-endian
	AV_PIX_FMT_GBRP14LE                            // planar GBR 4:4:4 42bpp, little-endian
	AV_PIX_FMT_YUVJ411P                            // planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples) full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV411P and setting color_range
	AV_PIX_FMT_BAYER_BGGR8                         // bayer, BGBG..(odd line), GRGR..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_RGGB8                         // bayer, RGRG..(odd line), GBGB..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_GBRG8                         // bayer, GBGB..(odd line), RGRG..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_GRBG8                         // bayer, GRGR..(odd line), BGBG..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_BGGR16LE                      // bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_BGGR16BE                      // bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_RGGB16LE                      // bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_RGGB16BE                      // bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_GBRG16LE                      // bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_GBRG16BE                      // bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_GRBG16LE                      // bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_GRBG16BE                      // bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_XVMC                                // XVideo Motion Acceleration via common packet passing.
	AV_PIX_FMT_YUV440P10LE                         // planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P10BE                         // planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_YUV440P12LE                         // planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P12BE                         // planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_AYUV64LE                            // packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_AYUV64BE                            // packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AV_PIX_FMT_VIDEOTOOLBOX                        // hardware decoding through Videotoolbox
	AV_PIX_FMT_P010LE                              // like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, little-endian
	AV_PIX_FMT_P010BE                              // like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, big-endian
	AV_PIX_FMT_GBRAP12BE                           // planar GBR 4:4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRAP12LE                           // planar GBR 4:4:4:4 48bpp, little-endian
	AV_PIX_FMT_GBRAP10BE                           // planar GBR 4:4:4:4 40bpp, big-endian
	AV_PIX_FMT_GBRAP10LE                           // planar GBR 4:4:4:4 40bpp, little-endian
	AV_PIX_FMT_MEDIACODEC                          // hardware decoding through MediaCodec
	AV_PIX_FMT_GRAY12BE                            // Y , 12bpp, big-endian.
	AV_PIX_FMT_GRAY12LE                            // Y , 12bpp, little-endian.
	AV_PIX_FMT_GRAY10BE                            // Y , 10bpp, big-endian.
	AV_PIX_FMT_GRAY10LE                            // Y , 10bpp, little-endian.
	AV_PIX_FMT_P016LE                              // like NV12, with 16bpp per component, little-endian
	AV_PIX_FMT_P016BE                              // like NV12, with 16bpp per component, big-endian
	AV_PIX_FMT_D3D11                               // Hardware surfaces for Direct3D11.
	AV_PIX_FMT_GRAY9BE                             // Y , 9bpp, big-endian.
	AV_PIX_FMT_GRAY9LE                             // Y , 9bpp, little-endian.
	AV_PIX_FMT_GBRPF32BE                           // IEEE-754 single precision planar GBR 4:4:4, 96bpp, big-endian.
	AV_PIX_FMT_GBRPF32LE                           // IEEE-754 single precision planar GBR 4:4:4, 96bpp, little-endian.
	AV_PIX_FMT_GBRAPF32BE                          // IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, big-endian.
	AV_PIX_FMT_GBRAPF32LE                          // IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, little-endian.
	AV_PIX_FMT_DRM_PRIME                           // DRM-managed buffers exposed through PRIME buffer sharing.
	AV_PIX_FMT_OPENCL                              // Hardware surfaces for OpenCL.
	AV_PIX_FMT_GRAY14BE                            // Y , 14bpp, big-endian.
	AV_PIX_FMT_GRAY14LE                            // Y , 14bpp, little-endian.
	AV_PIX_FMT_GRAYF32BE                           // IEEE-754 single precision Y, 32bpp, big-endian.
	AV_PIX_FMT_GRAYF32LE                           // IEEE-754 single precision Y, 32bpp, little-endian.
	AV_PIX_FMT_NONE           AVPixelFormat = -1
	AV_PIX_FMT_MIN                          = AV_PIX_FMT_YUV420P
	AV_PIX_FMT_MAX                          = AV_PIX_FMT_GRAYF32LE
)

const (
	AV_SAMPLE_FMT_NONE AVSampleFormat = iota
	AV_SAMPLE_FMT_U8                  //	unsigned 8 bits
	AV_SAMPLE_FMT_S16                 //	signed 16 bits
	AV_SAMPLE_FMT_S32                 //	signed 32 bits
	AV_SAMPLE_FMT_FLT                 //	float
	AV_SAMPLE_FMT_DBL                 //	double
	AV_SAMPLE_FMT_U8P                 //	unsigned 8 bits, planar
	AV_SAMPLE_FMT_S16P                //	signed 16 bits, planar
	AV_SAMPLE_FMT_S32P                //	signed 32 bits, planar
	AV_SAMPLE_FMT_FLTP                //	float, planar
	AV_SAMPLE_FMT_DBLP                //	double, planar
	AV_SAMPLE_FMT_S64                 //	signed 64 bits
	AV_SAMPLE_FMT_S64P                //	signed 64 bits, planar
)

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v AVCodecId) String() string {
	switch v {
	case AV_CODEC_ID_NONE:
		return "AV_CODEC_ID_NONE"
	case AV_CODEC_ID_MPEG1VIDEO:
		return "AV_CODEC_ID_MPEG1VIDEO"
	case AV_CODEC_ID_MPEG2VIDEO:
		return "AV_CODEC_ID_MPEG2VIDEO"
	case AV_CODEC_ID_H261:
		return "AV_CODEC_ID_H261"
	case AV_CODEC_ID_H263:
		return "AV_CODEC_ID_H263"
	case AV_CODEC_ID_RV10:
		return "AV_CODEC_ID_RV10"
	case AV_CODEC_ID_RV20:
		return "AV_CODEC_ID_RV20"
	case AV_CODEC_ID_MJPEG:
		return "AV_CODEC_ID_MJPEG"
	case AV_CODEC_ID_MJPEGB:
		return "AV_CODEC_ID_MJPEGB"
	case AV_CODEC_ID_LJPEG:
		return "AV_CODEC_ID_LJPEG"
	case AV_CODEC_ID_SP5X:
		return "AV_CODEC_ID_SP5X"
	case AV_CODEC_ID_JPEGLS:
		return "AV_CODEC_ID_JPEGLS"
	case AV_CODEC_ID_MPEG4:
		return "AV_CODEC_ID_MPEG4"
	case AV_CODEC_ID_RAWVIDEO:
		return "AV_CODEC_ID_RAWVIDEO"
	case AV_CODEC_ID_MSMPEG4V1:
		return "AV_CODEC_ID_MSMPEG4V1"
	case AV_CODEC_ID_MSMPEG4V2:
		return "AV_CODEC_ID_MSMPEG4V2"
	case AV_CODEC_ID_MSMPEG4V3:
		return "AV_CODEC_ID_MSMPEG4V3"
	case AV_CODEC_ID_WMV1:
		return "AV_CODEC_ID_WMV1"
	case AV_CODEC_ID_WMV2:
		return "AV_CODEC_ID_WMV2"
	case AV_CODEC_ID_H263P:
		return "AV_CODEC_ID_H263P"
	case AV_CODEC_ID_H263I:
		return "AV_CODEC_ID_H263I"
	case AV_CODEC_ID_FLV1:
		return "AV_CODEC_ID_FLV1"
	case AV_CODEC_ID_SVQ1:
		return "AV_CODEC_ID_SVQ1"
	case AV_CODEC_ID_SVQ3:
		return "AV_CODEC_ID_SVQ3"
	case AV_CODEC_ID_DVVIDEO:
		return "AV_CODEC_ID_DVVIDEO"
	case AV_CODEC_ID_HUFFYUV:
		return "AV_CODEC_ID_HUFFYUV"
	case AV_CODEC_ID_CYUV:
		return "AV_CODEC_ID_CYUV"
	case AV_CODEC_ID_H264:
		return "AV_CODEC_ID_H264"
	case AV_CODEC_ID_INDEO3:
		return "AV_CODEC_ID_INDEO3"
	case AV_CODEC_ID_VP3:
		return "AV_CODEC_ID_VP3"
	case AV_CODEC_ID_THEORA:
		return "AV_CODEC_ID_THEORA"
	case AV_CODEC_ID_ASV1:
		return "AV_CODEC_ID_ASV1"
	case AV_CODEC_ID_ASV2:
		return "AV_CODEC_ID_ASV2"
	case AV_CODEC_ID_FFV1:
		return "AV_CODEC_ID_FFV1"
	case AV_CODEC_ID_4XM:
		return "AV_CODEC_ID_4XM"
	case AV_CODEC_ID_VCR1:
		return "AV_CODEC_ID_VCR1"
	case AV_CODEC_ID_CLJR:
		return "AV_CODEC_ID_CLJR"
	case AV_CODEC_ID_MDEC:
		return "AV_CODEC_ID_MDEC"
	case AV_CODEC_ID_ROQ:
		return "AV_CODEC_ID_ROQ"
	case AV_CODEC_ID_INTERPLAY_VIDEO:
		return "AV_CODEC_ID_INTERPLAY_VIDEO"
	case AV_CODEC_ID_XAN_WC3:
		return "AV_CODEC_ID_XAN_WC3"
	case AV_CODEC_ID_XAN_WC4:
		return "AV_CODEC_ID_XAN_WC4"
	case AV_CODEC_ID_RPZA:
		return "AV_CODEC_ID_RPZA"
	case AV_CODEC_ID_CINEPAK:
		return "AV_CODEC_ID_CINEPAK"
	case AV_CODEC_ID_WS_VQA:
		return "AV_CODEC_ID_WS_VQA"
	case AV_CODEC_ID_MSRLE:
		return "AV_CODEC_ID_MSRLE"
	case AV_CODEC_ID_MSVIDEO1:
		return "AV_CODEC_ID_MSVIDEO1"
	case AV_CODEC_ID_IDCIN:
		return "AV_CODEC_ID_IDCIN"
	case AV_CODEC_ID_8BPS:
		return "AV_CODEC_ID_8BPS"
	case AV_CODEC_ID_SMC:
		return "AV_CODEC_ID_SMC"
	case AV_CODEC_ID_FLIC:
		return "AV_CODEC_ID_FLIC"
	case AV_CODEC_ID_TRUEMOTION1:
		return "AV_CODEC_ID_TRUEMOTION1"
	case AV_CODEC_ID_VMDVIDEO:
		return "AV_CODEC_ID_VMDVIDEO"
	case AV_CODEC_ID_MSZH:
		return "AV_CODEC_ID_MSZH"
	case AV_CODEC_ID_ZLIB:
		return "AV_CODEC_ID_ZLIB"
	case AV_CODEC_ID_QTRLE:
		return "AV_CODEC_ID_QTRLE"
	case AV_CODEC_ID_TSCC:
		return "AV_CODEC_ID_TSCC"
	case AV_CODEC_ID_ULTI:
		return "AV_CODEC_ID_ULTI"
	case AV_CODEC_ID_QDRAW:
		return "AV_CODEC_ID_QDRAW"
	case AV_CODEC_ID_VIXL:
		return "AV_CODEC_ID_VIXL"
	case AV_CODEC_ID_QPEG:
		return "AV_CODEC_ID_QPEG"
	case AV_CODEC_ID_PNG:
		return "AV_CODEC_ID_PNG"
	case AV_CODEC_ID_PPM:
		return "AV_CODEC_ID_PPM"
	case AV_CODEC_ID_PBM:
		return "AV_CODEC_ID_PBM"
	case AV_CODEC_ID_PGM:
		return "AV_CODEC_ID_PGM"
	case AV_CODEC_ID_PGMYUV:
		return "AV_CODEC_ID_PGMYUV"
	case AV_CODEC_ID_PAM:
		return "AV_CODEC_ID_PAM"
	case AV_CODEC_ID_FFVHUFF:
		return "AV_CODEC_ID_FFVHUFF"
	case AV_CODEC_ID_RV30:
		return "AV_CODEC_ID_RV30"
	case AV_CODEC_ID_RV40:
		return "AV_CODEC_ID_RV40"
	case AV_CODEC_ID_VC1:
		return "AV_CODEC_ID_VC1"
	case AV_CODEC_ID_WMV3:
		return "AV_CODEC_ID_WMV3"
	case AV_CODEC_ID_LOCO:
		return "AV_CODEC_ID_LOCO"
	case AV_CODEC_ID_WNV1:
		return "AV_CODEC_ID_WNV1"
	case AV_CODEC_ID_AASC:
		return "AV_CODEC_ID_AASC"
	case AV_CODEC_ID_INDEO2:
		return "AV_CODEC_ID_INDEO2"
	case AV_CODEC_ID_FRAPS:
		return "AV_CODEC_ID_FRAPS"
	case AV_CODEC_ID_TRUEMOTION2:
		return "AV_CODEC_ID_TRUEMOTION2"
	case AV_CODEC_ID_BMP:
		return "AV_CODEC_ID_BMP"
	case AV_CODEC_ID_CSCD:
		return "AV_CODEC_ID_CSCD"
	case AV_CODEC_ID_MMVIDEO:
		return "AV_CODEC_ID_MMVIDEO"
	case AV_CODEC_ID_ZMBV:
		return "AV_CODEC_ID_ZMBV"
	case AV_CODEC_ID_AVS:
		return "AV_CODEC_ID_AVS"
	case AV_CODEC_ID_SMACKVIDEO:
		return "AV_CODEC_ID_SMACKVIDEO"
	case AV_CODEC_ID_NUV:
		return "AV_CODEC_ID_NUV"
	case AV_CODEC_ID_KMVC:
		return "AV_CODEC_ID_KMVC"
	case AV_CODEC_ID_FLASHSV:
		return "AV_CODEC_ID_FLASHSV"
	case AV_CODEC_ID_CAVS:
		return "AV_CODEC_ID_CAVS"
	case AV_CODEC_ID_JPEG2000:
		return "AV_CODEC_ID_JPEG2000"
	case AV_CODEC_ID_VMNC:
		return "AV_CODEC_ID_VMNC"
	case AV_CODEC_ID_VP5:
		return "AV_CODEC_ID_VP5"
	case AV_CODEC_ID_VP6:
		return "AV_CODEC_ID_VP6"
	case AV_CODEC_ID_VP6F:
		return "AV_CODEC_ID_VP6F"
	case AV_CODEC_ID_TARGA:
		return "AV_CODEC_ID_TARGA"
	case AV_CODEC_ID_DSICINVIDEO:
		return "AV_CODEC_ID_DSICINVIDEO"
	case AV_CODEC_ID_TIERTEXSEQVIDEO:
		return "AV_CODEC_ID_TIERTEXSEQVIDEO"
	case AV_CODEC_ID_TIFF:
		return "AV_CODEC_ID_TIFF"
	case AV_CODEC_ID_GIF:
		return "AV_CODEC_ID_GIF"
	case AV_CODEC_ID_DXA:
		return "AV_CODEC_ID_DXA"
	case AV_CODEC_ID_DNXHD:
		return "AV_CODEC_ID_DNXHD"
	case AV_CODEC_ID_THP:
		return "AV_CODEC_ID_THP"
	case AV_CODEC_ID_SGI:
		return "AV_CODEC_ID_SGI"
	case AV_CODEC_ID_C93:
		return "AV_CODEC_ID_C93"
	case AV_CODEC_ID_BETHSOFTVID:
		return "AV_CODEC_ID_BETHSOFTVID"
	case AV_CODEC_ID_PTX:
		return "AV_CODEC_ID_PTX"
	case AV_CODEC_ID_TXD:
		return "AV_CODEC_ID_TXD"
	case AV_CODEC_ID_VP6A:
		return "AV_CODEC_ID_VP6A"
	case AV_CODEC_ID_AMV:
		return "AV_CODEC_ID_AMV"
	case AV_CODEC_ID_VB:
		return "AV_CODEC_ID_VB"
	case AV_CODEC_ID_PCX:
		return "AV_CODEC_ID_PCX"
	case AV_CODEC_ID_SUNRAST:
		return "AV_CODEC_ID_SUNRAST"
	case AV_CODEC_ID_INDEO4:
		return "AV_CODEC_ID_INDEO4"
	case AV_CODEC_ID_INDEO5:
		return "AV_CODEC_ID_INDEO5"
	case AV_CODEC_ID_MIMIC:
		return "AV_CODEC_ID_MIMIC"
	case AV_CODEC_ID_RL2:
		return "AV_CODEC_ID_RL2"
	case AV_CODEC_ID_ESCAPE124:
		return "AV_CODEC_ID_ESCAPE124"
	case AV_CODEC_ID_DIRAC:
		return "AV_CODEC_ID_DIRAC"
	case AV_CODEC_ID_BFI:
		return "AV_CODEC_ID_BFI"
	case AV_CODEC_ID_CMV:
		return "AV_CODEC_ID_CMV"
	case AV_CODEC_ID_MOTIONPIXELS:
		return "AV_CODEC_ID_MOTIONPIXELS"
	case AV_CODEC_ID_TGV:
		return "AV_CODEC_ID_TGV"
	case AV_CODEC_ID_TGQ:
		return "AV_CODEC_ID_TGQ"
	case AV_CODEC_ID_TQI:
		return "AV_CODEC_ID_TQI"
	case AV_CODEC_ID_AURA:
		return "AV_CODEC_ID_AURA"
	case AV_CODEC_ID_AURA2:
		return "AV_CODEC_ID_AURA2"
	case AV_CODEC_ID_V210X:
		return "AV_CODEC_ID_V210X"
	case AV_CODEC_ID_TMV:
		return "AV_CODEC_ID_TMV"
	case AV_CODEC_ID_V210:
		return "AV_CODEC_ID_V210"
	case AV_CODEC_ID_DPX:
		return "AV_CODEC_ID_DPX"
	case AV_CODEC_ID_MAD:
		return "AV_CODEC_ID_MAD"
	case AV_CODEC_ID_FRWU:
		return "AV_CODEC_ID_FRWU"
	case AV_CODEC_ID_FLASHSV2:
		return "AV_CODEC_ID_FLASHSV2"
	case AV_CODEC_ID_CDGRAPHICS:
		return "AV_CODEC_ID_CDGRAPHICS"
	case AV_CODEC_ID_R210:
		return "AV_CODEC_ID_R210"
	case AV_CODEC_ID_ANM:
		return "AV_CODEC_ID_ANM"
	case AV_CODEC_ID_BINKVIDEO:
		return "AV_CODEC_ID_BINKVIDEO"
	case AV_CODEC_ID_IFF_ILBM:
		return "AV_CODEC_ID_IFF_ILBM"
	case AV_CODEC_ID_KGV1:
		return "AV_CODEC_ID_KGV1"
	case AV_CODEC_ID_YOP:
		return "AV_CODEC_ID_YOP"
	case AV_CODEC_ID_VP8:
		return "AV_CODEC_ID_VP8"
	case AV_CODEC_ID_PICTOR:
		return "AV_CODEC_ID_PICTOR"
	case AV_CODEC_ID_ANSI:
		return "AV_CODEC_ID_ANSI"
	case AV_CODEC_ID_A64_MULTI:
		return "AV_CODEC_ID_A64_MULTI"
	case AV_CODEC_ID_A64_MULTI5:
		return "AV_CODEC_ID_A64_MULTI5"
	case AV_CODEC_ID_R10K:
		return "AV_CODEC_ID_R10K"
	case AV_CODEC_ID_MXPEG:
		return "AV_CODEC_ID_MXPEG"
	case AV_CODEC_ID_LAGARITH:
		return "AV_CODEC_ID_LAGARITH"
	case AV_CODEC_ID_PRORES:
		return "AV_CODEC_ID_PRORES"
	case AV_CODEC_ID_JV:
		return "AV_CODEC_ID_JV"
	case AV_CODEC_ID_DFA:
		return "AV_CODEC_ID_DFA"
	case AV_CODEC_ID_WMV3IMAGE:
		return "AV_CODEC_ID_WMV3IMAGE"
	case AV_CODEC_ID_VC1IMAGE:
		return "AV_CODEC_ID_VC1IMAGE"
	case AV_CODEC_ID_UTVIDEO:
		return "AV_CODEC_ID_UTVIDEO"
	case AV_CODEC_ID_BMV_VIDEO:
		return "AV_CODEC_ID_BMV_VIDEO"
	case AV_CODEC_ID_VBLE:
		return "AV_CODEC_ID_VBLE"
	case AV_CODEC_ID_DXTORY:
		return "AV_CODEC_ID_DXTORY"
	case AV_CODEC_ID_V410:
		return "AV_CODEC_ID_V410"
	case AV_CODEC_ID_XWD:
		return "AV_CODEC_ID_XWD"
	case AV_CODEC_ID_CDXL:
		return "AV_CODEC_ID_CDXL"
	case AV_CODEC_ID_XBM:
		return "AV_CODEC_ID_XBM"
	case AV_CODEC_ID_ZEROCODEC:
		return "AV_CODEC_ID_ZEROCODEC"
	case AV_CODEC_ID_MSS1:
		return "AV_CODEC_ID_MSS1"
	case AV_CODEC_ID_MSA1:
		return "AV_CODEC_ID_MSA1"
	case AV_CODEC_ID_TSCC2:
		return "AV_CODEC_ID_TSCC2"
	case AV_CODEC_ID_MTS2:
		return "AV_CODEC_ID_MTS2"
	case AV_CODEC_ID_CLLC:
		return "AV_CODEC_ID_CLLC"
	case AV_CODEC_ID_MSS2:
		return "AV_CODEC_ID_MSS2"
	case AV_CODEC_ID_VP9:
		return "AV_CODEC_ID_VP9"
	case AV_CODEC_ID_AIC:
		return "AV_CODEC_ID_AIC"
	case AV_CODEC_ID_ESCAPE130:
		return "AV_CODEC_ID_ESCAPE130"
	case AV_CODEC_ID_G2M:
		return "AV_CODEC_ID_G2M"
	case AV_CODEC_ID_WEBP:
		return "AV_CODEC_ID_WEBP"
	case AV_CODEC_ID_HNM4_VIDEO:
		return "AV_CODEC_ID_HNM4_VIDEO"
	case AV_CODEC_ID_HEVC:
		return "AV_CODEC_ID_HEVC"
	case AV_CODEC_ID_FIC:
		return "AV_CODEC_ID_FIC"
	case AV_CODEC_ID_ALIAS_PIX:
		return "AV_CODEC_ID_ALIAS_PIX"
	case AV_CODEC_ID_BRENDER_PIX:
		return "AV_CODEC_ID_BRENDER_PIX"
	case AV_CODEC_ID_PAF_VIDEO:
		return "AV_CODEC_ID_PAF_VIDEO"
	case AV_CODEC_ID_EXR:
		return "AV_CODEC_ID_EXR"
	case AV_CODEC_ID_VP7:
		return "AV_CODEC_ID_VP7"
	case AV_CODEC_ID_SANM:
		return "AV_CODEC_ID_SANM"
	case AV_CODEC_ID_SGIRLE:
		return "AV_CODEC_ID_SGIRLE"
	case AV_CODEC_ID_MVC1:
		return "AV_CODEC_ID_MVC1"
	case AV_CODEC_ID_MVC2:
		return "AV_CODEC_ID_MVC2"
	case AV_CODEC_ID_HQX:
		return "AV_CODEC_ID_HQX"
	case AV_CODEC_ID_TDSC:
		return "AV_CODEC_ID_TDSC"
	case AV_CODEC_ID_HQ_HQA:
		return "AV_CODEC_ID_HQ_HQA"
	case AV_CODEC_ID_HAP:
		return "AV_CODEC_ID_HAP"
	case AV_CODEC_ID_DDS:
		return "AV_CODEC_ID_DDS"
	case AV_CODEC_ID_DXV:
		return "AV_CODEC_ID_DXV"
	case AV_CODEC_ID_SCREENPRESSO:
		return "AV_CODEC_ID_SCREENPRESSO"
	case AV_CODEC_ID_RSCC:
		return "AV_CODEC_ID_RSCC"
	case AV_CODEC_ID_AVS2:
		return "AV_CODEC_ID_AVS2"
	case AV_CODEC_ID_Y41P:
		return "AV_CODEC_ID_Y41P"
	case AV_CODEC_ID_AVRP:
		return "AV_CODEC_ID_AVRP"
	case AV_CODEC_ID_012V:
		return "AV_CODEC_ID_012V"
	case AV_CODEC_ID_AVUI:
		return "AV_CODEC_ID_AVUI"
	case AV_CODEC_ID_AYUV:
		return "AV_CODEC_ID_AYUV"
	case AV_CODEC_ID_TARGA_Y216:
		return "AV_CODEC_ID_TARGA_Y216"
	case AV_CODEC_ID_V308:
		return "AV_CODEC_ID_V308"
	case AV_CODEC_ID_V408:
		return "AV_CODEC_ID_V408"
	case AV_CODEC_ID_YUV4:
		return "AV_CODEC_ID_YUV4"
	case AV_CODEC_ID_AVRN:
		return "AV_CODEC_ID_AVRN"
	case AV_CODEC_ID_CPIA:
		return "AV_CODEC_ID_CPIA"
	case AV_CODEC_ID_XFACE:
		return "AV_CODEC_ID_XFACE"
	case AV_CODEC_ID_SNOW:
		return "AV_CODEC_ID_SNOW"
	case AV_CODEC_ID_SMVJPEG:
		return "AV_CODEC_ID_SMVJPEG"
	case AV_CODEC_ID_APNG:
		return "AV_CODEC_ID_APNG"
	case AV_CODEC_ID_DAALA:
		return "AV_CODEC_ID_DAALA"
	case AV_CODEC_ID_CFHD:
		return "AV_CODEC_ID_CFHD"
	case AV_CODEC_ID_TRUEMOTION2RT:
		return "AV_CODEC_ID_TRUEMOTION2RT"
	case AV_CODEC_ID_M101:
		return "AV_CODEC_ID_M101"
	case AV_CODEC_ID_MAGICYUV:
		return "AV_CODEC_ID_MAGICYUV"
	case AV_CODEC_ID_SHEERVIDEO:
		return "AV_CODEC_ID_SHEERVIDEO"
	case AV_CODEC_ID_YLC:
		return "AV_CODEC_ID_YLC"
	case AV_CODEC_ID_PSD:
		return "AV_CODEC_ID_PSD"
	case AV_CODEC_ID_PIXLET:
		return "AV_CODEC_ID_PIXLET"
	case AV_CODEC_ID_SPEEDHQ:
		return "AV_CODEC_ID_SPEEDHQ"
	case AV_CODEC_ID_FMVC:
		return "AV_CODEC_ID_FMVC"
	case AV_CODEC_ID_SCPR:
		return "AV_CODEC_ID_SCPR"
	case AV_CODEC_ID_CLEARVIDEO:
		return "AV_CODEC_ID_CLEARVIDEO"
	case AV_CODEC_ID_XPM:
		return "AV_CODEC_ID_XPM"
	case AV_CODEC_ID_AV1:
		return "AV_CODEC_ID_AV1"
	case AV_CODEC_ID_BITPACKED:
		return "AV_CODEC_ID_BITPACKED"
	case AV_CODEC_ID_MSCC:
		return "AV_CODEC_ID_MSCC"
	case AV_CODEC_ID_SRGC:
		return "AV_CODEC_ID_SRGC"
	case AV_CODEC_ID_SVG:
		return "AV_CODEC_ID_SVG"
	case AV_CODEC_ID_GDV:
		return "AV_CODEC_ID_GDV"
	case AV_CODEC_ID_FITS:
		return "AV_CODEC_ID_FITS"
	case AV_CODEC_ID_IMM4:
		return "AV_CODEC_ID_IMM4"
	case AV_CODEC_ID_PROSUMER:
		return "AV_CODEC_ID_PROSUMER"
	case AV_CODEC_ID_MWSC:
		return "AV_CODEC_ID_MWSC"
	case AV_CODEC_ID_WCMV:
		return "AV_CODEC_ID_WCMV"
	case AV_CODEC_ID_RASC:
		return "AV_CODEC_ID_RASC"
	case AV_CODEC_ID_PCM_S16LE:
		return "AV_CODEC_ID_PCM_S16LE"
	case AV_CODEC_ID_PCM_S16BE:
		return "AV_CODEC_ID_PCM_S16BE"
	case AV_CODEC_ID_PCM_U16LE:
		return "AV_CODEC_ID_PCM_U16LE"
	case AV_CODEC_ID_PCM_U16BE:
		return "AV_CODEC_ID_PCM_U16BE"
	case AV_CODEC_ID_PCM_S8:
		return "AV_CODEC_ID_PCM_S8"
	case AV_CODEC_ID_PCM_U8:
		return "AV_CODEC_ID_PCM_U8"
	case AV_CODEC_ID_PCM_MULAW:
		return "AV_CODEC_ID_PCM_MULAW"
	case AV_CODEC_ID_PCM_ALAW:
		return "AV_CODEC_ID_PCM_ALAW"
	case AV_CODEC_ID_PCM_S32LE:
		return "AV_CODEC_ID_PCM_S32LE"
	case AV_CODEC_ID_PCM_S32BE:
		return "AV_CODEC_ID_PCM_S32BE"
	case AV_CODEC_ID_PCM_U32LE:
		return "AV_CODEC_ID_PCM_U32LE"
	case AV_CODEC_ID_PCM_U32BE:
		return "AV_CODEC_ID_PCM_U32BE"
	case AV_CODEC_ID_PCM_S24LE:
		return "AV_CODEC_ID_PCM_S24LE"
	case AV_CODEC_ID_PCM_S24BE:
		return "AV_CODEC_ID_PCM_S24BE"
	case AV_CODEC_ID_PCM_U24LE:
		return "AV_CODEC_ID_PCM_U24LE"
	case AV_CODEC_ID_PCM_U24BE:
		return "AV_CODEC_ID_PCM_U24BE"
	case AV_CODEC_ID_PCM_S24DAUD:
		return "AV_CODEC_ID_PCM_S24DAUD"
	case AV_CODEC_ID_PCM_ZORK:
		return "AV_CODEC_ID_PCM_ZORK"
	case AV_CODEC_ID_PCM_S16LE_PLANAR:
		return "AV_CODEC_ID_PCM_S16LE_PLANAR"
	case AV_CODEC_ID_PCM_DVD:
		return "AV_CODEC_ID_PCM_DVD"
	case AV_CODEC_ID_PCM_F32BE:
		return "AV_CODEC_ID_PCM_F32BE"
	case AV_CODEC_ID_PCM_F32LE:
		return "AV_CODEC_ID_PCM_F32LE"
	case AV_CODEC_ID_PCM_F64BE:
		return "AV_CODEC_ID_PCM_F64BE"
	case AV_CODEC_ID_PCM_F64LE:
		return "AV_CODEC_ID_PCM_F64LE"
	case AV_CODEC_ID_PCM_BLURAY:
		return "AV_CODEC_ID_PCM_BLURAY"
	case AV_CODEC_ID_PCM_LXF:
		return "AV_CODEC_ID_PCM_LXF"
	case AV_CODEC_ID_S302M:
		return "AV_CODEC_ID_S302M"
	case AV_CODEC_ID_PCM_S8_PLANAR:
		return "AV_CODEC_ID_PCM_S8_PLANAR"
	case AV_CODEC_ID_PCM_S24LE_PLANAR:
		return "AV_CODEC_ID_PCM_S24LE_PLANAR"
	case AV_CODEC_ID_PCM_S32LE_PLANAR:
		return "AV_CODEC_ID_PCM_S32LE_PLANAR"
	case AV_CODEC_ID_PCM_S16BE_PLANAR:
		return "AV_CODEC_ID_PCM_S16BE_PLANAR"
	case AV_CODEC_ID_PCM_S64LE:
		return "AV_CODEC_ID_PCM_S64LE"
	case AV_CODEC_ID_PCM_S64BE:
		return "AV_CODEC_ID_PCM_S64BE"
	case AV_CODEC_ID_PCM_F16LE:
		return "AV_CODEC_ID_PCM_F16LE"
	case AV_CODEC_ID_PCM_F24LE:
		return "AV_CODEC_ID_PCM_F24LE"
	case AV_CODEC_ID_PCM_VIDC:
		return "AV_CODEC_ID_PCM_VIDC"
	case AV_CODEC_ID_ADPCM_IMA_QT:
		return "AV_CODEC_ID_ADPCM_IMA_QT"
	case AV_CODEC_ID_ADPCM_IMA_WAV:
		return "AV_CODEC_ID_ADPCM_IMA_WAV"
	case AV_CODEC_ID_ADPCM_IMA_DK3:
		return "AV_CODEC_ID_ADPCM_IMA_DK3"
	case AV_CODEC_ID_ADPCM_IMA_DK4:
		return "AV_CODEC_ID_ADPCM_IMA_DK4"
	case AV_CODEC_ID_ADPCM_IMA_WS:
		return "AV_CODEC_ID_ADPCM_IMA_WS"
	case AV_CODEC_ID_ADPCM_IMA_SMJPEG:
		return "AV_CODEC_ID_ADPCM_IMA_SMJPEG"
	case AV_CODEC_ID_ADPCM_MS:
		return "AV_CODEC_ID_ADPCM_MS"
	case AV_CODEC_ID_ADPCM_4XM:
		return "AV_CODEC_ID_ADPCM_4XM"
	case AV_CODEC_ID_ADPCM_XA:
		return "AV_CODEC_ID_ADPCM_XA"
	case AV_CODEC_ID_ADPCM_ADX:
		return "AV_CODEC_ID_ADPCM_ADX"
	case AV_CODEC_ID_ADPCM_EA:
		return "AV_CODEC_ID_ADPCM_EA"
	case AV_CODEC_ID_ADPCM_G726:
		return "AV_CODEC_ID_ADPCM_G726"
	case AV_CODEC_ID_ADPCM_CT:
		return "AV_CODEC_ID_ADPCM_CT"
	case AV_CODEC_ID_ADPCM_SWF:
		return "AV_CODEC_ID_ADPCM_SWF"
	case AV_CODEC_ID_ADPCM_YAMAHA:
		return "AV_CODEC_ID_ADPCM_YAMAHA"
	case AV_CODEC_ID_ADPCM_SBPRO_4:
		return "AV_CODEC_ID_ADPCM_SBPRO_4"
	case AV_CODEC_ID_ADPCM_SBPRO_3:
		return "AV_CODEC_ID_ADPCM_SBPRO_3"
	case AV_CODEC_ID_ADPCM_SBPRO_2:
		return "AV_CODEC_ID_ADPCM_SBPRO_2"
	case AV_CODEC_ID_ADPCM_THP:
		return "AV_CODEC_ID_ADPCM_THP"
	case AV_CODEC_ID_ADPCM_IMA_AMV:
		return "AV_CODEC_ID_ADPCM_IMA_AMV"
	case AV_CODEC_ID_ADPCM_EA_R1:
		return "AV_CODEC_ID_ADPCM_EA_R1"
	case AV_CODEC_ID_ADPCM_EA_R3:
		return "AV_CODEC_ID_ADPCM_EA_R3"
	case AV_CODEC_ID_ADPCM_EA_R2:
		return "AV_CODEC_ID_ADPCM_EA_R2"
	case AV_CODEC_ID_ADPCM_IMA_EA_SEAD:
		return "AV_CODEC_ID_ADPCM_IMA_EA_SEAD"
	case AV_CODEC_ID_ADPCM_IMA_EA_EACS:
		return "AV_CODEC_ID_ADPCM_IMA_EA_EACS"
	case AV_CODEC_ID_ADPCM_EA_XAS:
		return "AV_CODEC_ID_ADPCM_EA_XAS"
	case AV_CODEC_ID_ADPCM_EA_MAXIS_XA:
		return "AV_CODEC_ID_ADPCM_EA_MAXIS_XA"
	case AV_CODEC_ID_ADPCM_IMA_ISS:
		return "AV_CODEC_ID_ADPCM_IMA_ISS"
	case AV_CODEC_ID_ADPCM_G722:
		return "AV_CODEC_ID_ADPCM_G722"
	case AV_CODEC_ID_ADPCM_IMA_APC:
		return "AV_CODEC_ID_ADPCM_IMA_APC"
	case AV_CODEC_ID_ADPCM_VIMA:
		return "AV_CODEC_ID_ADPCM_VIMA"
	case AV_CODEC_ID_ADPCM_AFC:
		return "AV_CODEC_ID_ADPCM_AFC"
	case AV_CODEC_ID_ADPCM_IMA_OKI:
		return "AV_CODEC_ID_ADPCM_IMA_OKI"
	case AV_CODEC_ID_ADPCM_DTK:
		return "AV_CODEC_ID_ADPCM_DTK"
	case AV_CODEC_ID_ADPCM_IMA_RAD:
		return "AV_CODEC_ID_ADPCM_IMA_RAD"
	case AV_CODEC_ID_ADPCM_G726LE:
		return "AV_CODEC_ID_ADPCM_G726LE"
	case AV_CODEC_ID_ADPCM_THP_LE:
		return "AV_CODEC_ID_ADPCM_THP_LE"
	case AV_CODEC_ID_ADPCM_PSX:
		return "AV_CODEC_ID_ADPCM_PSX"
	case AV_CODEC_ID_ADPCM_AICA:
		return "AV_CODEC_ID_ADPCM_AICA"
	case AV_CODEC_ID_ADPCM_IMA_DAT4:
		return "AV_CODEC_ID_ADPCM_IMA_DAT4"
	case AV_CODEC_ID_ADPCM_MTAF:
		return "AV_CODEC_ID_ADPCM_MTAF"
	case AV_CODEC_ID_AMR_NB:
		return "AV_CODEC_ID_AMR_NB"
	case AV_CODEC_ID_AMR_WB:
		return "AV_CODEC_ID_AMR_WB"
	case AV_CODEC_ID_RA_144:
		return "AV_CODEC_ID_RA_144"
	case AV_CODEC_ID_RA_288:
		return "AV_CODEC_ID_RA_288"
	case AV_CODEC_ID_ROQ_DPCM:
		return "AV_CODEC_ID_ROQ_DPCM"
	case AV_CODEC_ID_INTERPLAY_DPCM:
		return "AV_CODEC_ID_INTERPLAY_DPCM"
	case AV_CODEC_ID_XAN_DPCM:
		return "AV_CODEC_ID_XAN_DPCM"
	case AV_CODEC_ID_SOL_DPCM:
		return "AV_CODEC_ID_SOL_DPCM"
	case AV_CODEC_ID_SDX2_DPCM:
		return "AV_CODEC_ID_SDX2_DPCM"
	case AV_CODEC_ID_GREMLIN_DPCM:
		return "AV_CODEC_ID_GREMLIN_DPCM"
	case AV_CODEC_ID_MP2:
		return "AV_CODEC_ID_MP2"
	case AV_CODEC_ID_MP3:
		return "AV_CODEC_ID_MP3"
	case AV_CODEC_ID_AAC:
		return "AV_CODEC_ID_AAC"
	case AV_CODEC_ID_AC3:
		return "AV_CODEC_ID_AC3"
	case AV_CODEC_ID_DTS:
		return "AV_CODEC_ID_DTS"
	case AV_CODEC_ID_VORBIS:
		return "AV_CODEC_ID_VORBIS"
	case AV_CODEC_ID_DVAUDIO:
		return "AV_CODEC_ID_DVAUDIO"
	case AV_CODEC_ID_WMAV1:
		return "AV_CODEC_ID_WMAV1"
	case AV_CODEC_ID_WMAV2:
		return "AV_CODEC_ID_WMAV2"
	case AV_CODEC_ID_MACE3:
		return "AV_CODEC_ID_MACE3"
	case AV_CODEC_ID_MACE6:
		return "AV_CODEC_ID_MACE6"
	case AV_CODEC_ID_VMDAUDIO:
		return "AV_CODEC_ID_VMDAUDIO"
	case AV_CODEC_ID_FLAC:
		return "AV_CODEC_ID_FLAC"
	case AV_CODEC_ID_MP3ADU:
		return "AV_CODEC_ID_MP3ADU"
	case AV_CODEC_ID_MP3ON4:
		return "AV_CODEC_ID_MP3ON4"
	case AV_CODEC_ID_SHORTEN:
		return "AV_CODEC_ID_SHORTEN"
	case AV_CODEC_ID_ALAC:
		return "AV_CODEC_ID_ALAC"
	case AV_CODEC_ID_WESTWOOD_SND1:
		return "AV_CODEC_ID_WESTWOOD_SND1"
	case AV_CODEC_ID_GSM:
		return "AV_CODEC_ID_GSM"
	case AV_CODEC_ID_QDM2:
		return "AV_CODEC_ID_QDM2"
	case AV_CODEC_ID_COOK:
		return "AV_CODEC_ID_COOK"
	case AV_CODEC_ID_TRUESPEECH:
		return "AV_CODEC_ID_TRUESPEECH"
	case AV_CODEC_ID_TTA:
		return "AV_CODEC_ID_TTA"
	case AV_CODEC_ID_SMACKAUDIO:
		return "AV_CODEC_ID_SMACKAUDIO"
	case AV_CODEC_ID_QCELP:
		return "AV_CODEC_ID_QCELP"
	case AV_CODEC_ID_WAVPACK:
		return "AV_CODEC_ID_WAVPACK"
	case AV_CODEC_ID_DSICINAUDIO:
		return "AV_CODEC_ID_DSICINAUDIO"
	case AV_CODEC_ID_IMC:
		return "AV_CODEC_ID_IMC"
	case AV_CODEC_ID_MUSEPACK7:
		return "AV_CODEC_ID_MUSEPACK7"
	case AV_CODEC_ID_MLP:
		return "AV_CODEC_ID_MLP"
	case AV_CODEC_ID_GSM_MS:
		return "AV_CODEC_ID_GSM_MS"
	case AV_CODEC_ID_ATRAC3:
		return "AV_CODEC_ID_ATRAC3"
	case AV_CODEC_ID_APE:
		return "AV_CODEC_ID_APE"
	case AV_CODEC_ID_NELLYMOSER:
		return "AV_CODEC_ID_NELLYMOSER"
	case AV_CODEC_ID_MUSEPACK8:
		return "AV_CODEC_ID_MUSEPACK8"
	case AV_CODEC_ID_SPEEX:
		return "AV_CODEC_ID_SPEEX"
	case AV_CODEC_ID_WMAVOICE:
		return "AV_CODEC_ID_WMAVOICE"
	case AV_CODEC_ID_WMAPRO:
		return "AV_CODEC_ID_WMAPRO"
	case AV_CODEC_ID_WMALOSSLESS:
		return "AV_CODEC_ID_WMALOSSLESS"
	case AV_CODEC_ID_ATRAC3P:
		return "AV_CODEC_ID_ATRAC3P"
	case AV_CODEC_ID_EAC3:
		return "AV_CODEC_ID_EAC3"
	case AV_CODEC_ID_SIPR:
		return "AV_CODEC_ID_SIPR"
	case AV_CODEC_ID_MP1:
		return "AV_CODEC_ID_MP1"
	case AV_CODEC_ID_TWINVQ:
		return "AV_CODEC_ID_TWINVQ"
	case AV_CODEC_ID_TRUEHD:
		return "AV_CODEC_ID_TRUEHD"
	case AV_CODEC_ID_MP4ALS:
		return "AV_CODEC_ID_MP4ALS"
	case AV_CODEC_ID_ATRAC1:
		return "AV_CODEC_ID_ATRAC1"
	case AV_CODEC_ID_BINKAUDIO_RDFT:
		return "AV_CODEC_ID_BINKAUDIO_RDFT"
	case AV_CODEC_ID_BINKAUDIO_DCT:
		return "AV_CODEC_ID_BINKAUDIO_DCT"
	case AV_CODEC_ID_AAC_LATM:
		return "AV_CODEC_ID_AAC_LATM"
	case AV_CODEC_ID_QDMC:
		return "AV_CODEC_ID_QDMC"
	case AV_CODEC_ID_CELT:
		return "AV_CODEC_ID_CELT"
	case AV_CODEC_ID_G723_1:
		return "AV_CODEC_ID_G723_1"
	case AV_CODEC_ID_G729:
		return "AV_CODEC_ID_G729"
	case AV_CODEC_ID_8SVX_EXP:
		return "AV_CODEC_ID_8SVX_EXP"
	case AV_CODEC_ID_8SVX_FIB:
		return "AV_CODEC_ID_8SVX_FIB"
	case AV_CODEC_ID_BMV_AUDIO:
		return "AV_CODEC_ID_BMV_AUDIO"
	case AV_CODEC_ID_RALF:
		return "AV_CODEC_ID_RALF"
	case AV_CODEC_ID_IAC:
		return "AV_CODEC_ID_IAC"
	case AV_CODEC_ID_ILBC:
		return "AV_CODEC_ID_ILBC"
	case AV_CODEC_ID_OPUS:
		return "AV_CODEC_ID_OPUS"
	case AV_CODEC_ID_COMFORT_NOISE:
		return "AV_CODEC_ID_COMFORT_NOISE"
	case AV_CODEC_ID_TAK:
		return "AV_CODEC_ID_TAK"
	case AV_CODEC_ID_METASOUND:
		return "AV_CODEC_ID_METASOUND"
	case AV_CODEC_ID_PAF_AUDIO:
		return "AV_CODEC_ID_PAF_AUDIO"
	case AV_CODEC_ID_ON2AVC:
		return "AV_CODEC_ID_ON2AVC"
	case AV_CODEC_ID_DSS_SP:
		return "AV_CODEC_ID_DSS_SP"
	case AV_CODEC_ID_CODEC2:
		return "AV_CODEC_ID_CODEC2"
	case AV_CODEC_ID_FFWAVESYNTH:
		return "AV_CODEC_ID_FFWAVESYNTH"
	case AV_CODEC_ID_SONIC:
		return "AV_CODEC_ID_SONIC"
	case AV_CODEC_ID_SONIC_LS:
		return "AV_CODEC_ID_SONIC_LS"
	case AV_CODEC_ID_EVRC:
		return "AV_CODEC_ID_EVRC"
	case AV_CODEC_ID_SMV:
		return "AV_CODEC_ID_SMV"
	case AV_CODEC_ID_DSD_LSBF:
		return "AV_CODEC_ID_DSD_LSBF"
	case AV_CODEC_ID_DSD_MSBF:
		return "AV_CODEC_ID_DSD_MSBF"
	case AV_CODEC_ID_DSD_LSBF_PLANAR:
		return "AV_CODEC_ID_DSD_LSBF_PLANAR"
	case AV_CODEC_ID_DSD_MSBF_PLANAR:
		return "AV_CODEC_ID_DSD_MSBF_PLANAR"
	case AV_CODEC_ID_4GV:
		return "AV_CODEC_ID_4GV"
	case AV_CODEC_ID_INTERPLAY_ACM:
		return "AV_CODEC_ID_INTERPLAY_ACM"
	case AV_CODEC_ID_XMA1:
		return "AV_CODEC_ID_XMA1"
	case AV_CODEC_ID_XMA2:
		return "AV_CODEC_ID_XMA2"
	case AV_CODEC_ID_DST:
		return "AV_CODEC_ID_DST"
	case AV_CODEC_ID_ATRAC3AL:
		return "AV_CODEC_ID_ATRAC3AL"
	case AV_CODEC_ID_ATRAC3PAL:
		return "AV_CODEC_ID_ATRAC3PAL"
	case AV_CODEC_ID_DOLBY_E:
		return "AV_CODEC_ID_DOLBY_E"
	case AV_CODEC_ID_APTX:
		return "AV_CODEC_ID_APTX"
	case AV_CODEC_ID_APTX_HD:
		return "AV_CODEC_ID_APTX_HD"
	case AV_CODEC_ID_SBC:
		return "AV_CODEC_ID_SBC"
	case AV_CODEC_ID_ATRAC9:
		return "AV_CODEC_ID_ATRAC9"
	case AV_CODEC_ID_DVD_SUBTITLE:
		return "AV_CODEC_ID_DVD_SUBTITLE"
	case AV_CODEC_ID_DVB_SUBTITLE:
		return "AV_CODEC_ID_DVB_SUBTITLE"
	case AV_CODEC_ID_TEXT:
		return "AV_CODEC_ID_TEXT"
	case AV_CODEC_ID_XSUB:
		return "AV_CODEC_ID_XSUB"
	case AV_CODEC_ID_SSA:
		return "AV_CODEC_ID_SSA"
	case AV_CODEC_ID_MOV_TEXT:
		return "AV_CODEC_ID_MOV_TEXT"
	case AV_CODEC_ID_HDMV_PGS_SUBTITLE:
		return "AV_CODEC_ID_HDMV_PGS_SUBTITLE"
	case AV_CODEC_ID_DVB_TELETEXT:
		return "AV_CODEC_ID_DVB_TELETEXT"
	case AV_CODEC_ID_SRT:
		return "AV_CODEC_ID_SRT"
	case AV_CODEC_ID_MICRODVD:
		return "AV_CODEC_ID_MICRODVD"
	case AV_CODEC_ID_EIA_608:
		return "AV_CODEC_ID_EIA_608"
	case AV_CODEC_ID_JACOSUB:
		return "AV_CODEC_ID_JACOSUB"
	case AV_CODEC_ID_SAMI:
		return "AV_CODEC_ID_SAMI"
	case AV_CODEC_ID_REALTEXT:
		return "AV_CODEC_ID_REALTEXT"
	case AV_CODEC_ID_STL:
		return "AV_CODEC_ID_STL"
	case AV_CODEC_ID_SUBVIEWER1:
		return "AV_CODEC_ID_SUBVIEWER1"
	case AV_CODEC_ID_SUBVIEWER:
		return "AV_CODEC_ID_SUBVIEWER"
	case AV_CODEC_ID_SUBRIP:
		return "AV_CODEC_ID_SUBRIP"
	case AV_CODEC_ID_WEBVTT:
		return "AV_CODEC_ID_WEBVTT"
	case AV_CODEC_ID_MPL2:
		return "AV_CODEC_ID_MPL2"
	case AV_CODEC_ID_VPLAYER:
		return "AV_CODEC_ID_VPLAYER"
	case AV_CODEC_ID_PJS:
		return "AV_CODEC_ID_PJS"
	case AV_CODEC_ID_ASS:
		return "AV_CODEC_ID_ASS"
	case AV_CODEC_ID_HDMV_TEXT_SUBTITLE:
		return "AV_CODEC_ID_HDMV_TEXT_SUBTITLE"
	case AV_CODEC_ID_TTML:
		return "AV_CODEC_ID_TTML"
	case AV_CODEC_ID_TTF:
		return "AV_CODEC_ID_TTF"
	case AV_CODEC_ID_SCTE_35:
		return "AV_CODEC_ID_SCTE_35"
	case AV_CODEC_ID_BINTEXT:
		return "AV_CODEC_ID_BINTEXT"
	case AV_CODEC_ID_XBIN:
		return "AV_CODEC_ID_XBIN"
	case AV_CODEC_ID_IDF:
		return "AV_CODEC_ID_IDF"
	case AV_CODEC_ID_OTF:
		return "AV_CODEC_ID_OTF"
	case AV_CODEC_ID_SMPTE_KLV:
		return "AV_CODEC_ID_SMPTE_KLV"
	case AV_CODEC_ID_DVD_NAV:
		return "AV_CODEC_ID_DVD_NAV"
	case AV_CODEC_ID_TIMED_ID3:
		return "AV_CODEC_ID_TIMED_ID3"
	case AV_CODEC_ID_BIN_DATA:
		return "AV_CODEC_ID_BIN_DATA"
	case AV_CODEC_ID_PROBE:
		return "AV_CODEC_ID_PROBE"
	case AV_CODEC_ID_MPEG2TS:
		return "AV_CODEC_ID_MPEG2TS"
	case AV_CODEC_ID_MPEG4SYSTEMS:
		return "AV_CODEC_ID_MPEG4SYSTEMS"
	case AV_CODEC_ID_FFMETADATA:
		return "AV_CODEC_ID_FFMETADATA"
	case AV_CODEC_ID_WRAPPED_AVFRAME:
		return "AV_CODEC_ID_WRAPPED_AVFRAME"
	default:
		return "[?? Invalid AVCodecId value]"
	}
}

func (v AVMediaType) String() string {
	switch v {
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
	case AVMEDIA_TYPE_UNKNOWN:
		return "AVMEDIA_TYPE_UNKNOWN"
	default:
		return "[?? Unknown AVMediaType value]"
	}
}

func (v AVCodecCap) String() string {
	if v == AV_CODEC_CAP_NONE {
		return v.FlagString()
	}
	str := ""
	for f := AV_CODEC_CAP_MIN; f != AV_CODEC_CAP_MAX; f <<= 1 {
		if f&v == f {
			str += "|" + f.FlagString()
		}
	}
	return strings.TrimPrefix(str, "|")
}

func (v AVCodecCap) FlagString() string {
	switch v {
	case AV_CODEC_CAP_NONE:
		return "AV_CODEC_CAP_NONE"
	case AV_CODEC_CAP_DRAW_HORIZ_BAND:
		return "AV_CODEC_CAP_DRAW_HORIZ_BAND"
	case AV_CODEC_CAP_DR1:
		return "AV_CODEC_CAP_DR1"
	case AV_CODEC_CAP_TRUNCATED:
		return "AV_CODEC_CAP_TRUNCATED"
	case AV_CODEC_CAP_DELAY:
		return "AV_CODEC_CAP_DELAY"
	case AV_CODEC_CAP_SMALL_LAST_FRAME:
		return "AV_CODEC_CAP_SMALL_LAST_FRAME"
	case AV_CODEC_CAP_SUBFRAMES:
		return "AV_CODEC_CAP_SUBFRAMES"
	case AV_CODEC_CAP_EXPERIMENTAL:
		return "AV_CODEC_CAP_EXPERIMENTAL"
	case AV_CODEC_CAP_CHANNEL_CONF:
		return "AV_CODEC_CAP_CHANNEL_CONF"
	case AV_CODEC_CAP_FRAME_THREADS:
		return "AV_CODEC_CAP_FRAME_THREADS"
	case AV_CODEC_CAP_SLICE_THREADS:
		return "AV_CODEC_CAP_SLICE_THREADS"
	case AV_CODEC_CAP_PARAM_CHANGE:
		return "AV_CODEC_CAP_PARAM_CHANGE"
	case AV_CODEC_CAP_AUTO_THREADS:
		return "AV_CODEC_CAP_AUTO_THREADS"
	case AV_CODEC_CAP_VARIABLE_FRAME_SIZE:
		return "AV_CODEC_CAP_VARIABLE_FRAME_SIZE"
	case AV_CODEC_CAP_AVOID_PROBING:
		return "AV_CODEC_CAP_AVOID_PROBING"
	case AV_CODEC_CAP_INTRA_ONLY:
		return "AV_CODEC_CAP_INTRA_ONLY"
	case AV_CODEC_CAP_LOSSLESS:
		return "AV_CODEC_CAP_LOSSLESS"
	case AV_CODEC_CAP_HARDWARE:
		return "AV_CODEC_CAP_HARDWARE"
	case AV_CODEC_CAP_HYBRID:
		return "AV_CODEC_CAP_HYBRID"
	default:
		return "[?? Invalid AVCodecCap value]"
	}
}

func (v AVDisposition) String() string {
	if v == AV_DISPOSITION_NONE {
		return v.FlagString()
	}
	str := ""
	for f := AV_DISPOSITION_MIN; f != AV_DISPOSITION_MAX; f <<= 1 {
		if f&v == f {
			str += "|" + f.FlagString()
		}
	}
	return strings.TrimPrefix(str, "|")
}

func (v AVDisposition) FlagString() string {
	switch v {
	case AV_DISPOSITION_NONE:
		return "AV_DISPOSITION_NONE"
	case AV_DISPOSITION_DEFAULT:
		return "AV_DISPOSITION_DEFAULT"
	case AV_DISPOSITION_DUB:
		return "AV_DISPOSITION_DUB"
	case AV_DISPOSITION_ORIGINAL:
		return "AV_DISPOSITION_ORIGINAL"
	case AV_DISPOSITION_COMMENT:
		return "AV_DISPOSITION_COMMENT"
	case AV_DISPOSITION_LYRICS:
		return "AV_DISPOSITION_LYRICS"
	case AV_DISPOSITION_KARAOKE:
		return "AV_DISPOSITION_KARAOKE"
	case AV_DISPOSITION_FORCED:
		return "AV_DISPOSITION_FORCED"
	case AV_DISPOSITION_HEARING_IMPAIRED:
		return "AV_DISPOSITION_HEARING_IMPAIRED"
	case AV_DISPOSITION_VISUAL_IMPAIRED:
		return "AV_DISPOSITION_VISUAL_IMPAIRED"
	case AV_DISPOSITION_CLEAN_EFFECTS:
		return "AV_DISPOSITION_CLEAN_EFFECTS"
	case AV_DISPOSITION_ATTACHED_PIC:
		return "AV_DISPOSITION_ATTACHED_PIC"
	case AV_DISPOSITION_TIMED_THUMBNAILS:
		return "AV_DISPOSITION_TIMED_THUMBNAILS"
	case AV_DISPOSITION_CAPTIONS:
		return "AV_DISPOSITION_CAPTIONS"
	case AV_DISPOSITION_DESCRIPTIONS:
		return "AV_DISPOSITION_DESCRIPTIONS"
	case AV_DISPOSITION_METADATA:
		return "AV_DISPOSITION_METADATA"
	case AV_DISPOSITION_DEPENDENT:
		return "AV_DISPOSITION_DEPENDENT"
	case AV_DISPOSITION_STILL_IMAGE:
		return "AV_DISPOSITION_STILL_IMAGE"
	default:
		return "[?? Invalid AVDisposition value]"
	}
}

func (v AVFormatFlag) String() string {
	if v == AVFMT_NONE {
		return v.FlagString()
	}
	str := ""
	for f := AVFMT_MIN; f != AVFMT_MAX; f <<= 1 {
		if f&v == f {
			str += "|" + f.FlagString()
		}
	}
	return strings.TrimPrefix(str, "|")
}

func (v AVFormatFlag) FlagString() string {
	switch v {
	case AVFMT_NONE:
		return "AVFMT_NONE"
	case AVFMT_NOFILE:
		return "AVFMT_NOFILE"
	case AVFMT_NEEDNUMBER:
		return "AVFMT_NEEDNUMBER"
	case AVFMT_SHOW_IDS:
		return "AVFMT_SHOW_IDS"
	case AVFMT_GLOBALHEADER:
		return "AVFMT_GLOBALHEADER"
	case AVFMT_NOTIMESTAMPS:
		return "AVFMT_NOTIMESTAMPS"
	case AVFMT_GENERIC_INDEX:
		return "AVFMT_GENERIC_INDEX"
	case AVFMT_TS_DISCONT:
		return "AVFMT_TS_DISCONT"
	case AVFMT_VARIABLE_FPS:
		return "AVFMT_VARIABLE_FPS"
	case AVFMT_NODIMENSIONS:
		return "AVFMT_NODIMENSIONS"
	case AVFMT_NOSTREAMS:
		return "AVFMT_NOSTREAMS"
	case AVFMT_NOBINSEARCH:
		return "AVFMT_NOBINSEARCH"
	case AVFMT_NOGENSEARCH:
		return "AVFMT_NOGENSEARCH"
	case AVFMT_NO_BYTE_SEEK:
		return "AVFMT_NO_BYTE_SEEK"
	case AVFMT_ALLOW_FLUSH:
		return "AVFMT_ALLOW_FLUSH"
	case AVFMT_TS_NONSTRICT:
		return "AVFMT_TS_NONSTRICT"
	case AVFMT_TS_NEGATIVE:
		return "AVFMT_TS_NEGATIVE"
	case AVFMT_SEEK_TO_PTS:
		return "AVFMT_SEEK_TO_PTS"
	default:
		return "[?? Invalid AVFormatFlag value]"
	}
}

func (v AVLogLevel) String() string {
	switch v {
	case AV_LOG_QUIET:
		return "AV_LOG_QUIET"
	case AV_LOG_PANIC:
		return "AV_LOG_PANIC"
	case AV_LOG_FATAL:
		return "AV_LOG_FATAL"
	case AV_LOG_ERROR:
		return "AV_LOG_ERROR"
	case AV_LOG_WARNING:
		return "AV_LOG_WARNING"
	case AV_LOG_INFO:
		return "AV_LOG_INFO"
	case AV_LOG_VERBOSE:
		return "AV_LOG_VERBOSE"
	case AV_LOG_DEBUG:
		return "AV_LOG_DEBUG"
	case AV_LOG_TRACE:
		return "AV_LOG_TRACE"
	default:
		return "[?? Invalid AVLogLevel value]"
	}
}

func (f AVPacketFlag) String() string {
	str := ""
	if f == AV_PKT_FLAG_NONE {
		return f.FlagString()
	}
	for v := AV_PKT_FLAG_MIN; v <= AV_PKT_FLAG_MAX; v <<= 1 {
		if f&v == v {
			str += v.FlagString() + "|"
		}
	}
	return strings.TrimSuffix(str, "|")
}

func (v AVPacketFlag) FlagString() string {
	switch v {
	case AV_PKT_FLAG_NONE:
		return "AV_PKT_FLAG_NONE"
	case AV_PKT_FLAG_KEY:
		return "AV_PKT_FLAG_KEY"
	case AV_PKT_FLAG_CORRUPT:
		return "AV_PKT_FLAG_CORRUPT"
	case AV_PKT_FLAG_DISCARD:
		return "AV_PKT_FLAG_DISCARD"
	case AV_PKT_FLAG_TRUSTED:
		return "AV_PKT_FLAG_TRUSTED"
	case AV_PKT_FLAG_DISPOSABLE:
		return "AV_PKT_FLAG_DISPOSABLE"
	default:
		return "[?? Invalid AVPacketFlag]"
	}
}

func (p AVPictureType) String() string {
	switch p {
	case AV_PICTURE_TYPE_NONE:
		return "AV_PICTURE_TYPE_NONE"
	case AV_PICTURE_TYPE_I:
		return "AV_PICTURE_TYPE_I"
	case AV_PICTURE_TYPE_P:
		return "AV_PICTURE_TYPE_P"
	case AV_PICTURE_TYPE_B:
		return "AV_PICTURE_TYPE_B"
	case AV_PICTURE_TYPE_S:
		return "AV_PICTURE_TYPE_S"
	case AV_PICTURE_TYPE_SI:
		return "AV_PICTURE_TYPE_SI"
	case AV_PICTURE_TYPE_SP:
		return "AV_PICTURE_TYPE_SP"
	case AV_PICTURE_TYPE_BI:
		return "AV_PICTURE_TYPE_BI"
	default:
		return "[?? Invalid AVPictureType value]"
	}
}

func (c AVChannelLayout) String() string {
	if c == AV_CH_LAYOUT_NONE {
		return c.FlagString()
	}
	str := ""
	for v := AV_CH_LAYOUT_MIN; v <= AV_CH_LAYOUT_MAX; v <<= 1 {
		if c&v == v {
			str += v.FlagString() + "|"
		}
	}
	return strings.TrimSuffix(str, "|")
}

func (c AVChannelLayout) FlagString() string {
	switch c {
	case AV_CH_FRONT_LEFT:
		return "AV_CH_FRONT_LEFT"
	case AV_CH_FRONT_RIGHT:
		return "AV_CH_FRONT_RIGHT"
	case AV_CH_FRONT_CENTER:
		return "AV_CH_FRONT_CENTER"
	case AV_CH_LOW_FREQUENCY:
		return "AV_CH_LOW_FREQUENCY"
	case AV_CH_BACK_LEFT:
		return "AV_CH_BACK_LEFT"
	case AV_CH_BACK_RIGHT:
		return "AV_CH_BACK_RIGHT"
	case AV_CH_FRONT_LEFT_OF_CENTER:
		return "AV_CH_FRONT_LEFT_OF_CENTER"
	case AV_CH_FRONT_RIGHT_OF_CENTER:
		return "AV_CH_FRONT_RIGHT_OF_CENTER"
	case AV_CH_BACK_CENTER:
		return "AV_CH_BACK_CENTER"
	case AV_CH_SIDE_LEFT:
		return "AV_CH_SIDE_LEFT"
	case AV_CH_SIDE_RIGHT:
		return "AV_CH_SIDE_RIGHT"
	case AV_CH_TOP_CENTER:
		return "AV_CH_TOP_CENTER"
	case AV_CH_TOP_FRONT_LEFT:
		return "AV_CH_TOP_FRONT_LEFT"
	case AV_CH_TOP_FRONT_CENTER:
		return "AV_CH_TOP_FRONT_CENTER"
	case AV_CH_TOP_FRONT_RIGHT:
		return "AV_CH_TOP_FRONT_RIGHT"
	case AV_CH_TOP_BACK_LEFT:
		return "AV_CH_TOP_BACK_LEFT"
	case AV_CH_TOP_BACK_CENTER:
		return "AV_CH_TOP_BACK_CENTER"
	case AV_CH_TOP_BACK_RIGHT:
		return "AV_CH_TOP_BACK_RIGHT"
	case AV_CH_STEREO_LEFT:
		return "AV_CH_STEREO_LEFT"
	case AV_CH_STEREO_RIGHT:
		return "AV_CH_STEREO_RIGHT"
	default:
		return "[?? Invalid AVChannelLayout value]"
	}
}

func (f AVPixelFormat) String() string {
	switch f {
	case AV_PIX_FMT_NONE:
		return "AV_PIX_FMT_NONE"
	case AV_PIX_FMT_YUV420P:
		return "AV_PIX_FMT_YUV420P"
	case AV_PIX_FMT_YUYV422:
		return "AV_PIX_FMT_YUYV422"
	case AV_PIX_FMT_RGB24:
		return "AV_PIX_FMT_RGB24"
	case AV_PIX_FMT_BGR24:
		return "AV_PIX_FMT_BGR24"
	case AV_PIX_FMT_YUV422P:
		return "AV_PIX_FMT_YUV422P"
	case AV_PIX_FMT_YUV444P:
		return "AV_PIX_FMT_YUV444P"
	case AV_PIX_FMT_YUV410P:
		return "AV_PIX_FMT_YUV410P"
	case AV_PIX_FMT_YUV411P:
		return "AV_PIX_FMT_YUV411P"
	case AV_PIX_FMT_GRAY8:
		return "AV_PIX_FMT_GRAY8"
	case AV_PIX_FMT_MONOWHITE:
		return "AV_PIX_FMT_MONOWHITE"
	case AV_PIX_FMT_MONOBLACK:
		return "AV_PIX_FMT_MONOBLACK"
	case AV_PIX_FMT_PAL8:
		return "AV_PIX_FMT_PAL8"
	case AV_PIX_FMT_YUVJ420P:
		return "AV_PIX_FMT_YUVJ420P"
	case AV_PIX_FMT_YUVJ422P:
		return "AV_PIX_FMT_YUVJ422P"
	case AV_PIX_FMT_YUVJ444P:
		return "AV_PIX_FMT_YUVJ444P"
	case AV_PIX_FMT_UYVY422:
		return "AV_PIX_FMT_UYVY422"
	case AV_PIX_FMT_UYYVYY411:
		return "AV_PIX_FMT_UYYVYY411"
	case AV_PIX_FMT_BGR8:
		return "AV_PIX_FMT_BGR8"
	case AV_PIX_FMT_BGR4:
		return "AV_PIX_FMT_BGR4"
	case AV_PIX_FMT_BGR4_BYTE:
		return "AV_PIX_FMT_BGR4_BYTE"
	case AV_PIX_FMT_RGB8:
		return "AV_PIX_FMT_RGB8"
	case AV_PIX_FMT_RGB4:
		return "AV_PIX_FMT_RGB4"
	case AV_PIX_FMT_RGB4_BYTE:
		return "AV_PIX_FMT_RGB4_BYTE"
	case AV_PIX_FMT_NV12:
		return "AV_PIX_FMT_NV12"
	case AV_PIX_FMT_NV21:
		return "AV_PIX_FMT_NV21"
	case AV_PIX_FMT_ARGB:
		return "AV_PIX_FMT_ARGB"
	case AV_PIX_FMT_RGBA:
		return "AV_PIX_FMT_RGBA"
	case AV_PIX_FMT_ABGR:
		return "AV_PIX_FMT_ABGR"
	case AV_PIX_FMT_BGRA:
		return "AV_PIX_FMT_BGRA"
	case AV_PIX_FMT_GRAY16BE:
		return "AV_PIX_FMT_GRAY16BE"
	case AV_PIX_FMT_GRAY16LE:
		return "AV_PIX_FMT_GRAY16LE"
	case AV_PIX_FMT_YUV440P:
		return "AV_PIX_FMT_YUV440P"
	case AV_PIX_FMT_YUVJ440P:
		return "AV_PIX_FMT_YUVJ440P"
	case AV_PIX_FMT_YUVA420P:
		return "AV_PIX_FMT_YUVA420P"
	case AV_PIX_FMT_RGB48BE:
		return "AV_PIX_FMT_RGB48BE"
	case AV_PIX_FMT_RGB48LE:
		return "AV_PIX_FMT_RGB48LE"
	case AV_PIX_FMT_RGB565BE:
		return "AV_PIX_FMT_RGB565BE"
	case AV_PIX_FMT_RGB565LE:
		return "AV_PIX_FMT_RGB565LE"
	case AV_PIX_FMT_RGB555BE:
		return "AV_PIX_FMT_RGB555BE"
	case AV_PIX_FMT_RGB555LE:
		return "AV_PIX_FMT_RGB555LE"
	case AV_PIX_FMT_BGR565BE:
		return "AV_PIX_FMT_BGR565BE"
	case AV_PIX_FMT_BGR565LE:
		return "AV_PIX_FMT_BGR565LE"
	case AV_PIX_FMT_BGR555BE:
		return "AV_PIX_FMT_BGR555BE"
	case AV_PIX_FMT_BGR555LE:
		return "AV_PIX_FMT_BGR555LE"
	case AV_PIX_FMT_VAAPI_MOCO:
		return "AV_PIX_FMT_VAAPI_MOCO"
	case AV_PIX_FMT_VAAPI_IDCT:
		return "AV_PIX_FMT_VAAPI_IDCT"
	case AV_PIX_FMT_VAAPI_VLD:
		return "AV_PIX_FMT_VAAPI_VLD"
	case AV_PIX_FMT_VAAPI:
		return "AV_PIX_FMT_VAAPI"
	case AV_PIX_FMT_YUV420P16LE:
		return "AV_PIX_FMT_YUV420P16LE"
	case AV_PIX_FMT_YUV420P16BE:
		return "AV_PIX_FMT_YUV420P16BE"
	case AV_PIX_FMT_YUV422P16LE:
		return "AV_PIX_FMT_YUV422P16LE"
	case AV_PIX_FMT_YUV422P16BE:
		return "AV_PIX_FMT_YUV422P16BE"
	case AV_PIX_FMT_YUV444P16LE:
		return "AV_PIX_FMT_YUV444P16LE"
	case AV_PIX_FMT_YUV444P16BE:
		return "AV_PIX_FMT_YUV444P16BE"
	case AV_PIX_FMT_DXVA2_VLD:
		return "AV_PIX_FMT_DXVA2_VLD"
	case AV_PIX_FMT_RGB444LE:
		return "AV_PIX_FMT_RGB444LE"
	case AV_PIX_FMT_RGB444BE:
		return "AV_PIX_FMT_RGB444BE"
	case AV_PIX_FMT_BGR444LE:
		return "AV_PIX_FMT_BGR444LE"
	case AV_PIX_FMT_BGR444BE:
		return "AV_PIX_FMT_BGR444BE"
	case AV_PIX_FMT_YA8:
		return "AV_PIX_FMT_YA8"
	case AV_PIX_FMT_Y400A:
		return "AV_PIX_FMT_Y400A"
	case AV_PIX_FMT_GRAY8A:
		return "AV_PIX_FMT_GRAY8A"
	case AV_PIX_FMT_BGR48BE:
		return "AV_PIX_FMT_BGR48BE"
	case AV_PIX_FMT_BGR48LE:
		return "AV_PIX_FMT_BGR48LE"
	case AV_PIX_FMT_YUV420P9BE:
		return "AV_PIX_FMT_YUV420P9BE"
	case AV_PIX_FMT_YUV420P9LE:
		return "AV_PIX_FMT_YUV420P9LE"
	case AV_PIX_FMT_YUV420P10BE:
		return "AV_PIX_FMT_YUV420P10BE"
	case AV_PIX_FMT_YUV420P10LE:
		return "AV_PIX_FMT_YUV420P10LE"
	case AV_PIX_FMT_YUV422P10BE:
		return "AV_PIX_FMT_YUV422P10BE"
	case AV_PIX_FMT_YUV422P10LE:
		return "AV_PIX_FMT_YUV422P10LE"
	case AV_PIX_FMT_YUV444P9BE:
		return "AV_PIX_FMT_YUV444P9BE"
	case AV_PIX_FMT_YUV444P9LE:
		return "AV_PIX_FMT_YUV444P9LE"
	case AV_PIX_FMT_YUV444P10BE:
		return "AV_PIX_FMT_YUV444P10BE"
	case AV_PIX_FMT_YUV444P10LE:
		return "AV_PIX_FMT_YUV444P10LE"
	case AV_PIX_FMT_YUV422P9BE:
		return "AV_PIX_FMT_YUV422P9BE"
	case AV_PIX_FMT_YUV422P9LE:
		return "AV_PIX_FMT_YUV422P9LE"
	case AV_PIX_FMT_GBRP:
		return "AV_PIX_FMT_GBRP"
	case AV_PIX_FMT_GBR24P:
		return "AV_PIX_FMT_GBR24P"
	case AV_PIX_FMT_GBRP9BE:
		return "AV_PIX_FMT_GBRP9BE"
	case AV_PIX_FMT_GBRP9LE:
		return "AV_PIX_FMT_GBRP9LE"
	case AV_PIX_FMT_GBRP10BE:
		return "AV_PIX_FMT_GBRP10BE"
	case AV_PIX_FMT_GBRP10LE:
		return "AV_PIX_FMT_GBRP10LE"
	case AV_PIX_FMT_GBRP16BE:
		return "AV_PIX_FMT_GBRP16BE"
	case AV_PIX_FMT_GBRP16LE:
		return "AV_PIX_FMT_GBRP16LE"
	case AV_PIX_FMT_YUVA422P:
		return "AV_PIX_FMT_YUVA422P"
	case AV_PIX_FMT_YUVA444P:
		return "AV_PIX_FMT_YUVA444P"
	case AV_PIX_FMT_YUVA420P9BE:
		return "AV_PIX_FMT_YUVA420P9BE"
	case AV_PIX_FMT_YUVA420P9LE:
		return "AV_PIX_FMT_YUVA420P9LE"
	case AV_PIX_FMT_YUVA422P9BE:
		return "AV_PIX_FMT_YUVA422P9BE"
	case AV_PIX_FMT_YUVA422P9LE:
		return "AV_PIX_FMT_YUVA422P9LE"
	case AV_PIX_FMT_YUVA444P9BE:
		return "AV_PIX_FMT_YUVA444P9BE"
	case AV_PIX_FMT_YUVA444P9LE:
		return "AV_PIX_FMT_YUVA444P9LE"
	case AV_PIX_FMT_YUVA420P10BE:
		return "AV_PIX_FMT_YUVA420P10BE"
	case AV_PIX_FMT_YUVA420P10LE:
		return "AV_PIX_FMT_YUVA420P10LE"
	case AV_PIX_FMT_YUVA422P10BE:
		return "AV_PIX_FMT_YUVA422P10BE"
	case AV_PIX_FMT_YUVA422P10LE:
		return "AV_PIX_FMT_YUVA422P10LE"
	case AV_PIX_FMT_YUVA444P10BE:
		return "AV_PIX_FMT_YUVA444P10BE"
	case AV_PIX_FMT_YUVA444P10LE:
		return "AV_PIX_FMT_YUVA444P10LE"
	case AV_PIX_FMT_YUVA420P16BE:
		return "AV_PIX_FMT_YUVA420P16BE"
	case AV_PIX_FMT_YUVA420P16LE:
		return "AV_PIX_FMT_YUVA420P16LE"
	case AV_PIX_FMT_YUVA422P16BE:
		return "AV_PIX_FMT_YUVA422P16BE"
	case AV_PIX_FMT_YUVA422P16LE:
		return "AV_PIX_FMT_YUVA422P16LE"
	case AV_PIX_FMT_YUVA444P16BE:
		return "AV_PIX_FMT_YUVA444P16BE"
	case AV_PIX_FMT_YUVA444P16LE:
		return "AV_PIX_FMT_YUVA444P16LE"
	case AV_PIX_FMT_VDPAU:
		return "AV_PIX_FMT_VDPAU"
	case AV_PIX_FMT_XYZ12LE:
		return "AV_PIX_FMT_XYZ12LE"
	case AV_PIX_FMT_XYZ12BE:
		return "AV_PIX_FMT_XYZ12BE"
	case AV_PIX_FMT_NV16:
		return "AV_PIX_FMT_NV16"
	case AV_PIX_FMT_NV20LE:
		return "AV_PIX_FMT_NV20LE"
	case AV_PIX_FMT_NV20BE:
		return "AV_PIX_FMT_NV20BE"
	case AV_PIX_FMT_RGBA64BE:
		return "AV_PIX_FMT_RGBA64BE"
	case AV_PIX_FMT_RGBA64LE:
		return "AV_PIX_FMT_RGBA64LE"
	case AV_PIX_FMT_BGRA64BE:
		return "AV_PIX_FMT_BGRA64BE"
	case AV_PIX_FMT_BGRA64LE:
		return "AV_PIX_FMT_BGRA64LE"
	case AV_PIX_FMT_YVYU422:
		return "AV_PIX_FMT_YVYU422"
	case AV_PIX_FMT_YA16BE:
		return "AV_PIX_FMT_YA16BE"
	case AV_PIX_FMT_YA16LE:
		return "AV_PIX_FMT_YA16LE"
	case AV_PIX_FMT_GBRAP:
		return "AV_PIX_FMT_GBRAP"
	case AV_PIX_FMT_GBRAP16BE:
		return "AV_PIX_FMT_GBRAP16BE"
	case AV_PIX_FMT_GBRAP16LE:
		return "AV_PIX_FMT_GBRAP16LE"
	case AV_PIX_FMT_QSV:
		return "AV_PIX_FMT_QSV"
	case AV_PIX_FMT_MMAL:
		return "AV_PIX_FMT_MMAL"
	case AV_PIX_FMT_D3D11VA_VLD:
		return "AV_PIX_FMT_D3D11VA_VLD"
	case AV_PIX_FMT_CUDA:
		return "AV_PIX_FMT_CUDA"
	case AV_PIX_FMT_0RGB:
		return "AV_PIX_FMT_0RGB"
	case AV_PIX_FMT_RGB0:
		return "AV_PIX_FMT_RGB0"
	case AV_PIX_FMT_0BGR:
		return "AV_PIX_FMT_0BGR"
	case AV_PIX_FMT_BGR0:
		return "AV_PIX_FMT_BGR0"
	case AV_PIX_FMT_YUV420P12BE:
		return "AV_PIX_FMT_YUV420P12BE"
	case AV_PIX_FMT_YUV420P12LE:
		return "AV_PIX_FMT_YUV420P12LE"
	case AV_PIX_FMT_YUV420P14BE:
		return "AV_PIX_FMT_YUV420P14BE"
	case AV_PIX_FMT_YUV420P14LE:
		return "AV_PIX_FMT_YUV420P14LE"
	case AV_PIX_FMT_YUV422P12BE:
		return "AV_PIX_FMT_YUV422P12BE"
	case AV_PIX_FMT_YUV422P12LE:
		return "AV_PIX_FMT_YUV422P12LE"
	case AV_PIX_FMT_YUV422P14BE:
		return "AV_PIX_FMT_YUV422P14BE"
	case AV_PIX_FMT_YUV422P14LE:
		return "AV_PIX_FMT_YUV422P14LE"
	case AV_PIX_FMT_YUV444P12BE:
		return "AV_PIX_FMT_YUV444P12BE"
	case AV_PIX_FMT_YUV444P12LE:
		return "AV_PIX_FMT_YUV444P12LE"
	case AV_PIX_FMT_YUV444P14BE:
		return "AV_PIX_FMT_YUV444P14BE"
	case AV_PIX_FMT_YUV444P14LE:
		return "AV_PIX_FMT_YUV444P14LE"
	case AV_PIX_FMT_GBRP12BE:
		return "AV_PIX_FMT_GBRP12BE"
	case AV_PIX_FMT_GBRP12LE:
		return "AV_PIX_FMT_GBRP12LE"
	case AV_PIX_FMT_GBRP14BE:
		return "AV_PIX_FMT_GBRP14BE"
	case AV_PIX_FMT_GBRP14LE:
		return "AV_PIX_FMT_GBRP14LE"
	case AV_PIX_FMT_YUVJ411P:
		return "AV_PIX_FMT_YUVJ411P"
	case AV_PIX_FMT_BAYER_BGGR8:
		return "AV_PIX_FMT_BAYER_BGGR8"
	case AV_PIX_FMT_BAYER_RGGB8:
		return "AV_PIX_FMT_BAYER_RGGB8"
	case AV_PIX_FMT_BAYER_GBRG8:
		return "AV_PIX_FMT_BAYER_GBRG8"
	case AV_PIX_FMT_BAYER_GRBG8:
		return "AV_PIX_FMT_BAYER_GRBG8"
	case AV_PIX_FMT_BAYER_BGGR16LE:
		return "AV_PIX_FMT_BAYER_BGGR16LE"
	case AV_PIX_FMT_BAYER_BGGR16BE:
		return "AV_PIX_FMT_BAYER_BGGR16BE"
	case AV_PIX_FMT_BAYER_RGGB16LE:
		return "AV_PIX_FMT_BAYER_RGGB16LE"
	case AV_PIX_FMT_BAYER_RGGB16BE:
		return "AV_PIX_FMT_BAYER_RGGB16BE"
	case AV_PIX_FMT_BAYER_GBRG16LE:
		return "AV_PIX_FMT_BAYER_GBRG16LE"
	case AV_PIX_FMT_BAYER_GBRG16BE:
		return "AV_PIX_FMT_BAYER_GBRG16BE"
	case AV_PIX_FMT_BAYER_GRBG16LE:
		return "AV_PIX_FMT_BAYER_GRBG16LE"
	case AV_PIX_FMT_BAYER_GRBG16BE:
		return "AV_PIX_FMT_BAYER_GRBG16BE"
	case AV_PIX_FMT_XVMC:
		return "AV_PIX_FMT_XVMC"
	case AV_PIX_FMT_YUV440P10LE:
		return "AV_PIX_FMT_YUV440P10LE"
	case AV_PIX_FMT_YUV440P10BE:
		return "AV_PIX_FMT_YUV440P10BE"
	case AV_PIX_FMT_YUV440P12LE:
		return "AV_PIX_FMT_YUV440P12LE"
	case AV_PIX_FMT_YUV440P12BE:
		return "AV_PIX_FMT_YUV440P12BE"
	case AV_PIX_FMT_AYUV64LE:
		return "AV_PIX_FMT_AYUV64LE"
	case AV_PIX_FMT_AYUV64BE:
		return "AV_PIX_FMT_AYUV64BE"
	case AV_PIX_FMT_VIDEOTOOLBOX:
		return "AV_PIX_FMT_VIDEOTOOLBOX"
	case AV_PIX_FMT_P010LE:
		return "AV_PIX_FMT_P010LE"
	case AV_PIX_FMT_P010BE:
		return "AV_PIX_FMT_P010BE"
	case AV_PIX_FMT_GBRAP12BE:
		return "AV_PIX_FMT_GBRAP12BE"
	case AV_PIX_FMT_GBRAP12LE:
		return "AV_PIX_FMT_GBRAP12LE"
	case AV_PIX_FMT_GBRAP10BE:
		return "AV_PIX_FMT_GBRAP10BE"
	case AV_PIX_FMT_GBRAP10LE:
		return "AV_PIX_FMT_GBRAP10LE"
	case AV_PIX_FMT_MEDIACODEC:
		return "AV_PIX_FMT_MEDIACODEC"
	case AV_PIX_FMT_GRAY12BE:
		return "AV_PIX_FMT_GRAY12BE"
	case AV_PIX_FMT_GRAY12LE:
		return "AV_PIX_FMT_GRAY12LE"
	case AV_PIX_FMT_GRAY10BE:
		return "AV_PIX_FMT_GRAY10BE"
	case AV_PIX_FMT_GRAY10LE:
		return "AV_PIX_FMT_GRAY10LE"
	case AV_PIX_FMT_P016LE:
		return "AV_PIX_FMT_P016LE"
	case AV_PIX_FMT_P016BE:
		return "AV_PIX_FMT_P016BE"
	case AV_PIX_FMT_D3D11:
		return "AV_PIX_FMT_D3D11"
	case AV_PIX_FMT_GRAY9BE:
		return "AV_PIX_FMT_GRAY9BE"
	case AV_PIX_FMT_GRAY9LE:
		return "AV_PIX_FMT_GRAY9LE"
	case AV_PIX_FMT_GBRPF32BE:
		return "AV_PIX_FMT_GBRPF32BE"
	case AV_PIX_FMT_GBRPF32LE:
		return "AV_PIX_FMT_GBRPF32LE"
	case AV_PIX_FMT_GBRAPF32BE:
		return "AV_PIX_FMT_GBRAPF32BE"
	case AV_PIX_FMT_GBRAPF32LE:
		return "AV_PIX_FMT_GBRAPF32LE"
	case AV_PIX_FMT_DRM_PRIME:
		return "AV_PIX_FMT_DRM_PRIME"
	case AV_PIX_FMT_OPENCL:
		return "AV_PIX_FMT_OPENCL"
	case AV_PIX_FMT_GRAY14BE:
		return "AV_PIX_FMT_GRAY14BE"
	case AV_PIX_FMT_GRAY14LE:
		return "AV_PIX_FMT_GRAY14LE"
	case AV_PIX_FMT_GRAYF32BE:
		return "AV_PIX_FMT_GRAYF32BE"
	case AV_PIX_FMT_GRAYF32LE:
		return "AV_PIX_FMT_GRAYF32LE"
	default:
		return "[?? Invalid AVPixelFormat value]"
	}
}

func (f AVSampleFormat) String() string {
	switch f {
	case AV_SAMPLE_FMT_NONE:
		return "AV_SAMPLE_FMT_NONE"
	case AV_SAMPLE_FMT_U8:
		return "AV_SAMPLE_FMT_U8"
	case AV_SAMPLE_FMT_S16:
		return "AV_SAMPLE_FMT_S16"
	case AV_SAMPLE_FMT_S32:
		return "AV_SAMPLE_FMT_S32"
	case AV_SAMPLE_FMT_FLT:
		return "AV_SAMPLE_FMT_FLT"
	case AV_SAMPLE_FMT_DBL:
		return "AV_SAMPLE_FMT_DBL"
	case AV_SAMPLE_FMT_U8P:
		return "AV_SAMPLE_FMT_U8P"
	case AV_SAMPLE_FMT_S16P:
		return "AV_SAMPLE_FMT_S16P"
	case AV_SAMPLE_FMT_S32P:
		return "AV_SAMPLE_FMT_S32P"
	case AV_SAMPLE_FMT_FLTP:
		return "AV_SAMPLE_FMT_FLTP"
	case AV_SAMPLE_FMT_DBLP:
		return "AV_SAMPLE_FMT_DBLP"
	case AV_SAMPLE_FMT_S64:
		return "AV_SAMPLE_FMT_S64"
	case AV_SAMPLE_FMT_S64P:
		return "AV_SAMPLE_FMT_S64P"
	default:
		return "[?? Invalid AVSampleFormat value]"
	}
}
