package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	// Check we can just run our server
	go StartServer()
}
