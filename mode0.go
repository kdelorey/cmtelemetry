package cmtelemetry

import (
	"encoding/binary"
	"errors"
	"math"
)

type mode0Accessor struct {
	buffer       *[]byte
	fieldOffsets map[TelemetryField]uint32
}

func createMode0Accessor() (a TelemetryAccessor, b []byte) {

	offsets := make(map[TelemetryField]uint32)

	offsets[TotalTime] = 0
	offsets[AngularVelocityX] = 4
	offsets[AngularVelocityY] = 8
	offsets[AngularVelocityZ] = 12
	offsets[Yaw] = 16
	offsets[Pitch] = 20
	offsets[Roll] = 24
	offsets[AccelerationX] = 28
	offsets[AccelerationY] = 32
	offsets[AccelerationZ] = 36
	offsets[VelocityX] = 40
	offsets[VelocityY] = 44
	offsets[VelocityZ] = 48
	offsets[PositionX] = 52
	offsets[PositionY] = 56
	offsets[PositionZ] = 60

	// Buffer is an extra 4 bytes long to suppor the fourcc field
	buffer := make([]byte, 68)

	acc := mode0Accessor{
		buffer:       &buffer,
		fieldOffsets: offsets,
	}

	return acc, buffer
}

func (a mode0Accessor) GetFieldValue(field TelemetryField) (float32, error) {
	offset, ok := a.fieldOffsets[field]

	if !ok {
		return 0.0, errors.New("invalid field for mode 0")
	}

	// Format:
	// 	- total_time, 			uint32, scale 1000.0
	//	- angular_velocity_x, 	float, 	scale -1
	//	- angular_velocity_y, 	float, 	scale 1
	//	- angular_velocity_z, 	float, 	scale 1
	//	- yaw, 					float, 	scale 1
	//	- pitch, 				float, 	scale 1
	//	- role, 				float, 	scale 1
	//	- yaw, 					float, 	scale 1
	//	- acceleration_x,		float, 	scale -1
	//	- acceleration_y,		float, 	scale 1
	//	- acceleration_z,		float, 	scale 1
	//	- velocity_x,			float, 	scale -1
	//	- velocity_y,			float, 	scale 1
	//	- velocity_z,			float, 	scale 1
	//	- position_x,			int32, 	scale -65535
	//	- position_y,			int32, 	scale 65535
	//	- position_z,			int32, 	scale 65535
	//	- fourcc text="ToCA"

	switch field {
	case TotalTime:
	case PositionX:
	case PositionY:
	case PositionZ:
		return 0, errors.New("too lazy to support ints")
	}

	bits := binary.LittleEndian.Uint32((*a.buffer)[offset : offset+4])
	return math.Float32frombits(bits), nil
}
