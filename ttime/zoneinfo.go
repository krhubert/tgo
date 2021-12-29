package ttime

import (
	"testing"
	"time"
)

func LoadLocation(t testing.TB, name string) *time.Location {
	t.Helper()
	r0, err := time.LoadLocation(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
