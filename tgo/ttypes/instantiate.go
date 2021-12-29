package ttypes

import (
	"go/types"
	"testing"
)

func Instantiate(t testing.TB, ctxt *types.Context, orig types.Type, targs []types.Type, validate bool) types.Type {
	t.Helper()
	r0, err := types.Instantiate(ctxt, orig, targs, validate)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
