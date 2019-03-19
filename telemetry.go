package cmtelemetry

import (
	"net"
)

// Telemetry is the basic
type Telemetry struct {
	dataAccessor TelemetryAccessor

	udp    *net.UDPConn
	buffer []byte
	quit   chan struct{}
	rec    chan TelemetryAccessor
}

// TelemetryAccessor is an interface that is responsible for accessing the current
// telemetry values.
type TelemetryAccessor interface {
	GetFieldValue(field TelemetryField) (float32, error)
}

// GatherDefaultTelemetry will open the default udp port and begin processing
// telemetry from Codemasters' games.
func GatherDefaultTelemetry(frames chan TelemetryAccessor) (*Telemetry, error) {
	return GatherMode3Telemetry(":20777", frames)
}

// GatherMode0Telemetry will open a udp port up at the specified location and begin
// processing telemetry from Codemasters' games.
func GatherMode0Telemetry(address string, frames chan TelemetryAccessor) (t *Telemetry, err error) {
	acc, buff := createMode0Accessor()

	return gatherTelemetry(
		address,
		frames,
		acc,
		buff)
}

// GatherMode3Telemetry will open a udp port up at the specified location and begin
// processing telemetry from Codemasters' games.
func GatherMode3Telemetry(address string, frames chan TelemetryAccessor) (t *Telemetry, err error) {
	acc, buff := createMode3Accessor()

	return gatherTelemetry(
		address,
		frames,
		acc,
		buff)
}

func gatherTelemetry(address string, frames chan TelemetryAccessor, accessor TelemetryAccessor, buffer []byte) (t *Telemetry, err error) {
	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		return
	}

	udp, err := net.ListenUDP("udp", addr)

	if err != nil {
		return
	}

	t = &Telemetry{
		dataAccessor: accessor,
		udp:          udp,
		buffer:       buffer,
		quit:         make(chan struct{}),
		rec:          frames,
	}

	go t.telemetryRoutine()

	return
}

// Close will stop the connection.
func (t *Telemetry) Close() {
	close(t.rec)
	close(t.quit)
	t.udp.Close()
}

func (t *Telemetry) telemetryRoutine() {
	for {
		select {
		case <-t.quit:
			return

		default:
			t.udp.ReadFromUDP(t.buffer)
			t.rec <- t.dataAccessor
			break
		}
	}
}
