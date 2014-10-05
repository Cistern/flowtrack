package flowtrack

import (
	"net"
	"testing"
)

func TestFlowKeyGen(t *testing.T) {
	source := net.IPv4(192, 168, 1, 1)
	sourcePort := 80

	destination := net.IPv4(192, 168, 1, 5)
	destPort := 23464

	t.Log(generateFlowKey(source, destination, sourcePort, destPort))
}
