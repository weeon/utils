package grpc

import (
	"net"
	"os"
)

func ListenerFromEnv() (net.Listener, error) {
	addr := os.Getenv("SRV_ADDR")
	if len(addr) == 0 {
		addr = ":8282"
	}
	return net.Listen("tcp", addr)
}
