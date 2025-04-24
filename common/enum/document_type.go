package enum

import (
	"encoding/json"

	"github.com/lamaleka/boilerplate-golang/common/errs"
)

type DocumentType int

const (
	LC DocumentType = iota + 1
	SKBDN
)

func (s DocumentType) Label() string {
	switch s {
	case LC:
		return "LC"
	case SKBDN:
		return "SKBDN"
	default:
		return "Unknown"
	}
}
func (s DocumentType) FromInt(value int) (DocumentType, error) {
	switch value {
	case 1:
		return LC, nil
	case 2:
		return SKBDN, nil
	default:
		return -1, errs.ErrInvalidDocumentType
	}
}

func (s DocumentType) FromString(value string) (DocumentType, error) {
	switch value {
	case "1":
		return LC, nil
	case "2":
		return SKBDN, nil
	default:
		return -1, errs.ErrInvalidDocumentType
	}
}

func (s DocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value": int(s),
		"label": s.Label(),
	})
}
