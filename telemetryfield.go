package dirt

// TelemetryField represents the id of a telemetry field.
type TelemetryField int16

const (
	// TotalTime represents the total time.
	TotalTime TelemetryField = iota

	// LapTime represents the time of the current lap.
	LapTime

	// LapDistance is the total distance traveled in the current lap.
	LapDistance

	// TotalDistance is the total distance travelled.
	TotalDistance

	// PositionX is the vehicle's x component of position in world space (in meters)
	PositionX

	// PositionY is the vehicle's y component of position in world space (in meters).
	PositionY

	// PositionZ is the vehicle's z component of position in world space (in meters).
	PositionZ

	// Speed is the magnitude of the vehicle's linear velocity in meters per
	// second (m/s)
	Speed

	// VelocityX is the x component of the linear velocity of the vehicle in
	// world space (in meters).
	VelocityX

	// VelocityY is the y component of the linear velocity of the vehicle in
	// world space (in meters).
	VelocityY

	// VelocityZ is the z component of the linear velocity of the vehicle in
	// world space (in meters).
	VelocityZ

	// LeftDirectionX is the x component of the left direction of the vehicle
	// in world space (in meters).
	LeftDirectionX

	// LeftDirectionY is the y component of the left direction of the vehicle
	// in world space (in meters).
	LeftDirectionY

	// LeftDirectionZ is the z component of the left direction of the vehicle
	// in world space (in meters).
	LeftDirectionZ

	// ForwardDirectionX is the x component of the forward direction of the vehicle
	// in world space (in meters).
	ForwardDirectionX

	// ForwardDirectionY is the y component of the forward direction of the vehicle
	// in world space (in meters).
	ForwardDirectionY

	// ForwardDirectionZ is the z component of the forward direction of the vehicle
	// in world space (in meters).
	ForwardDirectionZ

	// SuspensionPositionBackLeft is the vertical position of the wheel from the
	// rest (in vehicle space), measured in meters.
	SuspensionPositionBackLeft

	// SuspensionPositionBackRight is the vertical position of the wheel from the
	// rest (in vehicle space), measured in meters.
	SuspensionPositionBackRight

	// SuspensionPositionFrontLeft is the vertical position of the wheel from the
	// rest (in vehicle space), measured in meters.
	SuspensionPositionFrontLeft

	// SuspensionPositionFrontRight is the vertical position of the wheel from the
	// rest (in vehicle space), measured in meters.
	SuspensionPositionFrontRight

	// SuspensionVelocityBackLeft is the vertical velocity of the wheel from the
	// rest (in vehicle space), measured in meteres per second (m/s).
	SuspensionVelocityBackLeft

	// SuspensionVelocityBackRight is the vertical velocity of the wheel from the
	// rest (in vehicle space), measured in meteres per second (m/s).
	SuspensionVelocityBackRight

	// SuspensionVelocityFrontLeft is the vertical velocity of the wheel from the
	// rest (in vehicle space), measured in meteres per second (m/s).
	SuspensionVelocityFrontLeft

	// SuspensionVelocityFrontRight is the vertical velocity of the wheel from the
	// rest (in vehicle space), measured in meteres per second (m/s).
	SuspensionVelocityFrontRight
)
