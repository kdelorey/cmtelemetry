package cmtelemetry

import (
	"encoding/binary"
	"math"
	"net"
)

// Telemetry is the basic
type Telemetry struct {
	udp    *net.UDPConn
	buffer []byte
	quit   chan struct{}
}

// StartDefaultTelemetry will open the default udp port and begin processing
// telemetry from Codemasters' games.
func StartDefaultTelemetry() (*Telemetry, error) {
	return StartTelemetry(":20777")
}

// StartTelemetry will open a udp port up at the specified location and begin
// processing telemetry from Codemasters' games.
func StartTelemetry(address string) (t *Telemetry, err error) {
	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		return
	}

	udp, err := net.ListenUDP("udp", addr)

	if err != nil {
		return
	}

	t = &Telemetry{
		udp:    udp,
		buffer: make([]byte, 264),
		quit:   make(chan struct{}),
	}

	go t.telemetryRoutine()

	return
}

// Close will stop the connection.
func (t *Telemetry) Close() {
	close(t.quit)
	t.udp.Close()
}

// GetFieldValue will retrun the value of the specified telemetry field.
func (t *Telemetry) GetFieldValue(field TelemetryField) float32 {
	offset := int32(field) * 4

	bits := binary.LittleEndian.Uint32(t.buffer[offset : offset+4])

	return math.Float32frombits(bits)
}

func (t *Telemetry) telemetryRoutine() {
	for {
		select {
		case <-t.quit:
			return

		default:
			t.udp.ReadFromUDP(t.buffer)
			break
		}
	}
}
