package main

import (
	"testing"

	"github.com/krhubert/tgo/tos"
)

func TestOpen(t *testing.T) {
	tos.Open(t, "file").Close()
}
