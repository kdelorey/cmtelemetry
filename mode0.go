package cmtelemetry

import (
	"errors"
)

type mode0Accessor struct {
	buffer       *[]byte
	fieldOffsets map[TelemetryField]uint32
}

func createMode0Accessor() (a TelemetryAccessor, b []byte) {
	panic("not supported yet")
}

func (a mode0Accessor) GetFieldValue(field TelemetryField) (float32, error) {
	_, ok := a.fieldOffsets[field]

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
	//	- position_y,			int32, 	scint325535
	//	- position_z,			int32, 	scale 65535
	//	- fourcc text="ToCA"
	return 0, errors.New("not supported yet")
}
