package domain

import (
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	base := Base{
		ID:        "123",
		CreatedBy: "SYSTEM",
		UpdatedBy: "SYSTEM",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
		Metadata:  map[string]string{},
	}

	// Test ID field
	if base.ID != "123" {
		t.Errorf("unexpected ID value: got %s, want %s", base.ID, "123")
	}

	// Test CreatedBy field
	if base.CreatedBy != "SYSTEM" {
		t.Errorf("unexpected CreatedBy value: got %s, want %s", base.CreatedBy, "SYSTEM")
	}

	// Test UpdatedBy field
	if base.UpdatedBy != "SYSTEM" {
		t.Errorf("unexpected UpdatedBy value: got %s, want %s", base.UpdatedBy, "SYSTEM")
	}

	// Test CreatedAt field
	if base.CreatedAt.IsZero() {
		t.Error("unexpected zero value for CreatedAt")
	}

	// Test UpdatedAt field
	if base.UpdatedAt.IsZero() {
		t.Error("unexpected zero value for UpdatedAt")
	}

	// Test DeletedAt field
	if base.DeletedAt.IsZero() {
		t.Error("unexpected zero value for DeletedAt")
	}

	// Test Metadata field
	if base.Metadata == nil {
		t.Error("unexpected nil value for Metadata")
	}
}
