package ffmpeg

import (
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////
// CGO

import "C"

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS

func boolToInt(v bool) C.int {
	if v {
		return C.int(1)
	}
	return 0
}

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

func cByteSlice(p unsafe.Pointer, sz C.int) []byte {
	if p == nil {
		return nil
	}
	return (*[1 << 30]byte)(p)[:int(sz)]
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

func cFloat64Slice(p unsafe.Pointer, sz C.int) []float64 {
	if p == nil {
		return nil
	}
	return (*[1 << 30]float64)(p)[:int(sz)]
}

func cAVStreamSlice(p unsafe.Pointer, sz C.int) []*AVStream {
	if p == nil {
		return nil
	}
	return (*[1 << 30]*AVStream)(p)[:int(sz)]
}
