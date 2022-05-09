package database

import "testing"

func TestGet(t *testing.T) {
	db := Get()
	if db == nil {
		t.Fatal("db is nil")
	}
}
