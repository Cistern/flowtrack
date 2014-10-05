package flowtrack

import (
	"math/rand"
	"net"
	"testing"
)

func TestPrintTopTalkers(t *testing.T) {
	for i := 0; i < 500000; i++ {
		source := net.IPv4(192, 168, 1, byte(rand.Intn(10)))
		sourcePort := rand.Intn(1000)

		destination := net.IPv4(193, 168, 1, 1)
		destPort := rand.Intn(500)

		Process(source, destination, sourcePort, destPort, rand.Intn(1500))
	}

	t.Log(topFlowsPackets)

	PrintTopN(10)
}
