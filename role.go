package main

import "fmt"

// Role is a type that represents a role in a system
// it has a single field, name, which is a string
type Role struct {
	name string
}

// Name returns the name of the role
func (r Role) Name() string {
	return r.name
}

// NewRole creates a new role with the given name
func NewRole(name string) Role {
	return Role{name: name}
}

var (
	Admin = Role{"ADMIN"}
	User  = Role{"USER"}
)

// Roles is a map of role names to roles
var roles = map[string]Role{
	Admin.name: Admin,
	User.name:  User,
}

// ParseRole parses a role from a string
// returns an error if the role does not exist
func ParseRole(name string) (Role, error) {
	role, exists := roles[name]
	if !exists {
		return Role{}, fmt.Errorf("role %q does not exist", name)
	}

	return role, nil
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
// This allows us to unmarshal a role from a JSON string.
func (r *Role) UnmarshalText(name []byte) error {
	role, err := ParseRole(string(name))
	if err != nil {
		return err
	}

	*r = role

	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
// This allows us to marshal a role to a JSON string.
func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.name), nil
}

// Equal returns true if the roles are equal
func (r Role) Equal(other Role) bool {
	return r.name == other.name
}
