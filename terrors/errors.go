package terrors

import (
	"errors"
	"testing"
)

func New(t testing.TB, text string) {
	t.Helper()
	err := errors.New(text)
	if err != nil {
		t.Fatal(err)
	}
	return
}
