package ttime

import (
	"testing"
	"time"
)

func LoadLocationFromTZData(t testing.TB, name string, data []byte) *time.Location {
	t.Helper()
	r0, err := time.LoadLocationFromTZData(name, data)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
