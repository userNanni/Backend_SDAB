package main

import (
	"fmt"
	"net"
	"testing"
)

func TestFindAvailablePort(t *testing.T) {
	startPort := 3000
	port := findAvailablePort(startPort)

	if port < startPort {
		t.Errorf("Expected port to be greater than or equal to %d, but got %d", startPort, port)
	}

	// Verify that the returned port is actually available
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Errorf("Expected port %d to be available, but got error: %v", port, err)
	} else {
		l.Close()
	}
}
