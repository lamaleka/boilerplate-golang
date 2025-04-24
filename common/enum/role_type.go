package enum

import "encoding/json"

type RoleType int

const (
	Admin RoleType = iota + 1
	Officer
	Buyer
)

var roleTypeLabels = map[RoleType]string{
	Admin:   "Administrator",
	Officer: "Officer",
	Buyer:   "Buyer",
}

func (s RoleType) Label() string {
	switch s {
	case Admin:
		return "Administrator"
	case Officer:
		return "Officer"
	case Buyer:
		return "Buyer"
	default:
		return "Unknown"
	}
}
func (s RoleType) Value() int {
	switch s {
	case Admin:
		return 1
	case Officer:
		return 2
	case Buyer:
		return 3
	default:
		return -1
	}
}

func (s RoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value": int(s),
		"label": s.Label(),
	})
}
