package tuser

import (
	"os/user"
	"testing"
)

func Current(t testing.TB) *user.User {
	t.Helper()
	r0, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Lookup(t testing.TB, username string) *user.User {
	t.Helper()
	r0, err := user.Lookup(username)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupId(t testing.TB, uid string) *user.User {
	t.Helper()
	r0, err := user.LookupId(uid)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupGroup(t testing.TB, name string) *user.Group {
	t.Helper()
	r0, err := user.LookupGroup(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupGroupId(t testing.TB, gid string) *user.Group {
	t.Helper()
	r0, err := user.LookupGroupId(gid)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
