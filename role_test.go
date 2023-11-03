package main

import "testing"

func TestRole(t *testing.T) {
	role := Admin

	if role.Name() != "ADMIN" {
		t.Errorf("expected role name to be ADMIN, got %q", role.Name())
	}
}

func TestParseRole(t *testing.T) {
	t.Run("valid role ADMIN", func(t *testing.T) {
		role, err := ParseRole("ADMIN")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if role.Name() != "ADMIN" {
			t.Errorf("expected role name to be ADMIN, got %q", role.Name())
		}
	})

	t.Run("valid role USER", func(t *testing.T) {
		role, err := ParseRole("USER")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if role.Name() != "USER" {
			t.Errorf("expected role name to be USER, got %q", role.Name())
		}
	})

	t.Run("invalid role", func(t *testing.T) {
		_, err := ParseRole("INVALID")

		expected := "role \"INVALID\" does not exist"
		if err.Error() != expected {
			t.Errorf("expected error message %q, got %q", expected, err.Error())
		}
	})
}