package cmtelemetry

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

	// WheelPatchSpeedBackLeft is the speed of the contact patch of the back left
	// wheel in meters per second (m/s).
	WheelPatchSpeedBackLeft

	// WheelPatchSpeedBackRight is the speed of the contact patch of the back
	// right wheel in meters per second (m/s).
	WheelPatchSpeedBackRight

	// WheelPatchSpeedFrontLeft is the speed of the contact patch of the front
	// left wheel in meters per second (m/s).
	WheelPatchSpeedFrontLeft

	// WheelPatchSpeedFrontRight is the speed of the contact patch of the front
	// right wheel in meters per second (m/s).
	WheelPatchSpeedFrontRight

	// ThrottleInput is the level of throttle input applied through the controller.
	// A normalized value between 0 and 1.
	ThrottleInput

	// SteeringInput the level of steering input applied through the controller. A
	// value between -1 and 1.
	SteeringInput

	// BrakeInput is the level of brake input applied through the controller. A
	// normalized value between 0 and 1.
	BrakeInput

	// ClutchInput is the level of clutch input applied through the controller. A
	// normalized value between 0 and 1.
	ClutchInput

	// Gear is the current gear number. 0 for neutral, -1 for reverse.
	Gear

	// GForceLateral is the G-Force in the lateral direction.
	GForceLateral

	// GForceLongitudinal is the G-Force in the longitudinal direction.
	GForceLongitudinal

	// Lap is the current lap number.
	Lap

	// EngineRate is the engine rotations in radians per second.
	EngineRate

	// NativeSliSupport (undocumented, the following is speculation) is native
	// support for shift light indicators.
	NativeSliSupport

	// RacePosition is the current position within the race.
	RacePosition

	// KersLevel (undocumented, the following is speculation) is the level of
	// of the kinetic energy recovery system.
	//
	// Not supported in DiRT 4
	KersLevel

	// KersLevelMax (undocumented, the following is speculation) is the maximum
	// level of of the kinetic energy recovery system.
	//
	// Not supported in DiRT 4
	KersLevelMax

	// DRS (undocumented, the following is speculation) involves the drag reduction
	// system
	//
	// Not supported in DiRT 4
	DRS

	// TractionControl (undocumented, the following is speculation) is whether
	// traction control is engaged.
	//
	// Not supported in DiRT 4
	TractionControl

	// ABS (undocumented, the following is speculation) is whether ABS is engaged.
	//
	// Not supported in DiRT 4
	ABS

	// FuelInTank (undocumented, the following is speculation) is the amount of
	// fuel in the tank.
	//
	// Not supported in DiRT 4
	FuelInTank

	// FuelCapacity (undocumented, the following is speculation) is maximum
	// capacity of fuel.
	//
	// Not supported in DiRT 4
	FuelCapacity

	// InPits (undocumented, the following is speculation) is whether or not the
	// car is in the pit area.
	//
	// Not supported in DiRT 4
	InPits

	// RaceSector is the number of splits in the lap.
	RaceSector

	// SectorTime1 is the time of the first split.
	SectorTime1

	// SectorTime2 is the time of the second split.
	SectorTime2

	// BrakeTempBackLeft is the temperature of the back left brake in °C.
	BrakeTempBackLeft

	// BrakeTempBackRight is the temperature of the back right brake in °C.
	BrakeTempBackRight

	// BrakeTempFrontLeft is the temperature of the front left brake in °C.
	BrakeTempFrontLeft

	// BrakeTempFrontRight is the temperature of the front left brake in °C.
	BrakeTempFrontRight

	// TirePressureBackLeft (undocumented, the following is speculation) is the
	// current pressure of the back left tire.
	//
	// Not supported in DiRT 4
	TirePressureBackLeft

	// TirePressureBackRight (undocumented, the following is speculation) is the
	// current pressure of the back right tire.
	//
	// Not supported in DiRT 4
	TirePressureBackRight

	// TirePressureFrontLeft (undocumented, the following is speculation) is the
	// current pressure of the front left tire.
	//
	// Not supported in DiRT 4
	TirePressureFrontLeft

	// TirePressureFrontRight (undocumented, the following is speculation) is the
	// current pressure of the front right tire.
	//
	// Not supported in DiRT 4
	TirePressureFrontRight

	// LapsCompleted (undocumented, the following is speculation) is the number
	// of laps completed.
	LapsCompleted

	// TotalLaps is the total number of laps in the race.
	TotalLaps

	// TrackLength is the total length of the track in meters.
	TrackLength

	// LastLapTime is the time of the last lap, in seconds.
	LastLapTime

	// MaxRpm is the maximum engine rate, in radians per second.
	MaxRpm

	// IdleRpm is the idle rate of the engine in radians per second.
	IdleRpm

	// MaxGears is the number of forward gears available in the vehicle.
	MaxGears

	// AngularVelocityX is the X-component of the world-space angular velocity of the vehicle (in radians/s).
	AngularVelocityX

	// AngularVelocityY is the Y-component of the world-space angular velocity of the vehicle (in radians/s).
	AngularVelocityY

	// AngularVelocityZ is the Z-component of the world-space angular velocity of the vehicle (in radians/s).
	AngularVelocityZ

	// Yaw is the yaw of the vehicle (in radians).
	Yaw

	// Pitch is the pitch of the vehicle (in radians).
	Pitch

	// Roll is the roll of the vehicle (in radians).
	Roll

	// AccelerationX is the X-component of the world-space accelleration of the vehicle (in m/s^2).
	AccelerationX

	// AccelerationY is the Y-component of the world-space accelleration of the vehicle (in m/s^2).
	AccelerationY

	// AccelerationZ is the Z-component of the world-space accelleration of the vehicle (in m/s^2).
	AccelerationZ
)
