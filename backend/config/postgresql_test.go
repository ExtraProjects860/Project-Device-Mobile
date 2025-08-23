package config

import (
	"testing"
)

func TestConectionPostgreSQL(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatalf("ExpectedFailed to initialize database connection: %v", err)
	}

	db, err := GetDB().DB()
	if err != nil {
		t.Fatalf("Expected valid sql.DB got error! Error: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("Expected to ping database successfully, got %v", err)
	}
}
