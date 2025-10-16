package enum

type TypeUser uint8

const (
	SuperAdmin TypeUser = iota
	Admin
	User
)

func (t TypeUser) String() string {
	switch t {
	case SuperAdmin:
		return "SUPERADMIN"
	case Admin:
		return "ADMIN"
	case User:
		return "USER"
	default:
		return "UNKNOWN"
	}
}
