package enum

import (
	"fmt"
	"strings"
)

type ItemsOrder uint8

const (
	ASC ItemsOrder = iota
	DESC
)

func ParseItemsOrder(value string) (ItemsOrder, error) {
	switch strings.ToUpper(value) {
	case "ASC":
		return ASC, nil
	case "DESC":
		return DESC, nil
	default:
		return 0, fmt.Errorf("invalid items order: %s", value)
	}
}

func (i ItemsOrder) String() string {
	switch i {
	case ASC:
		return "ASC"
	case DESC:
		return "DESC"
	default:
		return "UNKNOWN"
	}
}
