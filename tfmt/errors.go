package tfmt

import (
	"fmt"
	"testing"
)

func Errorf(t testing.TB, format string, a ...any) {
	t.Helper()
	err := fmt.Errorf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return
}
