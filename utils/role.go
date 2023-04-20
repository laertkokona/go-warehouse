package utils

// Role constants
const (
	User     = 1
	Admin    = 2
	SysAdmin = 3
)

// GetRoleName returns the role name based on the role id
func GetRoleName(role int) string {
	switch role {
	case User:
		return "User"
	case Admin:
		return "Admin"
	case SysAdmin:
		return "SysAdmin"
	default:
		return "No Role"
	}
}
