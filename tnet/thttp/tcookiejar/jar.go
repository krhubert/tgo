package tcookiejar

import (
	"net/http/cookiejar"
	"testing"
)

func New(t testing.TB, o *cookiejar.Options) *cookiejar.Jar {
	t.Helper()
	r0, err := cookiejar.New(o)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
