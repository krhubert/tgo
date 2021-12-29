package tsyslog

import (
	"log"
	"log/syslog"
	"testing"
)

func New(t testing.TB, priority syslog.Priority, tag string) *syslog.Writer {
	t.Helper()
	r0, err := syslog.New(priority, tag)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Dial(t testing.TB, network, raddr string, priority syslog.Priority, tag string) *syslog.Writer {
	t.Helper()
	r0, err := syslog.Dial(network, raddr, priority, tag)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewLogger(t testing.TB, p syslog.Priority, logFlag int) *log.Logger {
	t.Helper()
	r0, err := syslog.NewLogger(p, logFlag)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
