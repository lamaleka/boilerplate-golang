package enum

import "encoding/json"

type EmployeeType int

const (
	TKO EmployeeType = iota + 1
	TKNO
	TKNOB
)

func (s EmployeeType) Label() string {
	switch s {
	case TKO:
		return "TKO"
	case TKNO:
		return "TKNO"
	case TKNOB:
		return "TKNO-Borongan"
	default:
		return "Unknown"
	}
}

func (s EmployeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value": int(s),
		"label": s.Label(),
	})
}
