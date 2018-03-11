package cmtelemetry

type mode0Accessor struct {
	buffer *[]byte
}

func createMode0Accessor() (a TelemetryAccessor, b []byte) {
	return
}

func (a mode0Accessor) GetFieldValue(field TelemetryField) float32 {
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
	return 0.0
}
