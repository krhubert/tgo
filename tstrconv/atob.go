package tstrconv

import (
	"strconv"
	"testing"
)

func ParseBool(t testing.TB, str string) bool {
	t.Helper()
	r0, err := strconv.ParseBool(str)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
