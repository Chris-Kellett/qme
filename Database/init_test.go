package database

import (
	"testing"
)

func Test_DatabasePing(t *testing.T) {
	Init()
	if Db == nil {
		t.Fatalf("DB was nil")
	}
	err := Db.Ping()
	if err != nil {
		t.Fatalf("DB Ping Failed: %s", err)
	}
}
