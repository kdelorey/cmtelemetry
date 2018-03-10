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

	// OverallDistance is the total distance travelled.
	OverallDistance
)
