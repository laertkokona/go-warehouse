package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetRoleName tests the GetRoleName function
func TestGetRoleName(t *testing.T) {
	role := GetRoleName(1)
	assert.Equal(t, role, "User", "Expected role to be User, got", role)
	role = GetRoleName(2)
	assert.Equal(t, role, "Admin", "Expected role to be Admin, got", role)
	role = GetRoleName(3)
	assert.Equal(t, role, "SysAdmin", "Expected role to be SysAdmin, got", role)
	role = GetRoleName(4)
	assert.Equal(t, role, "No Role", "Expected role to be No Role, got", role)
}
