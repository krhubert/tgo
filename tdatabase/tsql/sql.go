package tsql

import (
	"database/sql"
	"testing"
)

func Open(t testing.TB, driverName, dataSourceName string) *sql.DB {
	t.Helper()
	r0, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
