package tbig

import (
	"math/big"
	"testing"
)

func ParseFloat(t testing.TB, s string, base int, prec uint, mode big.RoundingMode) (f *big.Float, b int) {
	t.Helper()
	r0, r1, err := big.ParseFloat(s, base, prec, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
