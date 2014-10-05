package flowtrack

import (
	"fmt"
	"net"
)

type flowkey struct {
	SourceAddr      string
	SourcePort      int
	DestinationAddr string
	DestinationPort int
}

type addrPortKey struct {
	Addr string
	Port int
}

func generateFlowKey(source, destination net.IP, sourcePort, destPort int) flowkey {
	return flowkey{
		SourceAddr:      string(source),
		SourcePort:      sourcePort,
		DestinationAddr: string(destination),
		DestinationPort: destPort,
	}
}

func generateAddrPortKey(addr net.IP, port int) addrPortKey {
	return addrPortKey{
		Addr: string(addr),
		Port: port,
	}
}

func (k flowkey) String() string {
	return fmt.Sprintf("%s:%d -> %s:%d", net.IP(k.SourceAddr), k.SourcePort, net.IP(k.DestinationAddr), k.DestinationPort)
}

func (k addrPortKey) String() string {
	return fmt.Sprintf("%s:%d", net.IP(k.Addr), k.Port)
}
