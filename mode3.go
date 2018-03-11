package cmtelemetry

import (
	"encoding/binary"
	"math"
)

type mode3Accessor struct {
	buffer *[]byte
}

func createMode3Accessor() (a mode3Accessor, b []byte) {
	b = make([]byte, 264)
	a = mode3Accessor{buffer: &b}
	return
}

func (a mode3Accessor) GetFieldValue(field TelemetryField) float32 {
	offset := int32(field) * 4

	bits := binary.LittleEndian.Uint32((*a.buffer)[offset : offset+4])

	return math.Float32frombits(bits)
}
