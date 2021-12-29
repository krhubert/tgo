package tjson

import (
	"encoding/json"
	"testing"
)

func Unmarshal(t testing.TB, data []byte, v any) {
	t.Helper()
	err := json.Unmarshal(data, v)
	if err != nil {
		t.Fatal(err)
	}
	return
}
