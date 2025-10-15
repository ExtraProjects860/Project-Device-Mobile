package enum

type SqliteType uint8

const (
	File SqliteType = iota
	Memory
)

func (t SqliteType) String() string {
	switch t {
	case File:
		return "file"
	case Memory:
		return ":memory"
	default:
		return "unknown"
	}
}
