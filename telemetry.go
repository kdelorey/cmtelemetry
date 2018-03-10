package dirt

import (
	"encoding/binary"
	"math"
	"net"
)

// Telemetry is the basic
type Telemetry struct {
	udp    *net.UDPConn
	buffer []byte
}

// StartDefaultTelemetry will open the default udp port and begin processing
// telemetry from Codemasters' games.
func StartDefaultTelemetry() (*Telemetry, error) {
	return StartTelemetry(":20777")
}

// StartTelemetry will open a udp port up at the specified location and begin
// processing telemetry from Codemasters' games.
func StartTelemetry(address string) (telemetry *Telemetry, err error) {
	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		return
	}

	udp, err := net.ListenUDP("udp", addr)

	if err != nil {
		return
	}

	telemetry = &Telemetry{udp: udp, buffer: make([]byte, 264)}

	//go telemetry.telemetryRoutine()

	return
}

// Close will stop the connection.
func (telemetry *Telemetry) Close() {
	telemetry.Close()
}

// GetFieldValue todo
func (telemetry *Telemetry) GetFieldValue(field TelemetryField) float32 {
	offset := int32(field) * 4

	bits := binary.LittleEndian.Uint32(telemetry.buffer[offset : offset+4])

	return math.Float32frombits(bits)
}

func (telemetry *Telemetry) telemetryRoutine() {
	for {
		telemetry.udp.ReadFromUDP(telemetry.buffer)
	}
}
