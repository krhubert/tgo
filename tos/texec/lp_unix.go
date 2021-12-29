package texec

import (
	"os/exec"
	"testing"
)

func LookPath(t testing.TB, file string) string {
	t.Helper()
	r0, err := exec.LookPath(file)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
