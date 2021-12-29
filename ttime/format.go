package ttime

import (
	"testing"
	"time"
)

func Parse(t testing.TB, layout, value string) time.Time {
	t.Helper()
	r0, err := time.Parse(layout, value)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseInLocation(t testing.TB, layout, value string, loc *time.Location) time.Time {
	t.Helper()
	r0, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseDuration(t testing.TB, s string) time.Duration {
	t.Helper()
	r0, err := time.ParseDuration(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
