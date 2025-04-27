package main

import (
	"fmt"
	"net"
)

func findAvailablePort(startPort int) int {
	port := startPort
	for {
		addr := fmt.Sprintf(":%d", port)
		l, err := net.Listen("tcp", addr)
		if err == nil {
			l.Close()
			return port
		}
		port++
	}
}
