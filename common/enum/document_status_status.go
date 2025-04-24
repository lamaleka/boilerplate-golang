package enum

import (
	"encoding/json"

	"github.com/lamaleka/boilerplate-golang/common/errs"
)

type DocumentStatus int

const (
	Submitted DocumentStatus = iota + 1
	OnReview
	Rejected
	Approved
	OnLoading
	OnPayment
	Completed
)

func (s DocumentStatus) Label() string {
	switch s {
	case Submitted:
		return "Submitted"
	case OnReview:
		return "On Review"
	case Rejected:
		return "Rejected"
	case Approved:
		return "Approved"
	case OnLoading:
		return "On Loading"
	case OnPayment:
		return "On Payment"
	case Completed:
		return "Completed"
	default:
		return "Unknown Status"
	}
}

func (s DocumentStatus) FromInt(value int) (DocumentStatus, error) {
	switch value {
	case 1:
		return Submitted, nil
	case 2:
		return OnReview, nil
	case 3:
		return Rejected, nil
	case 4:
		return Approved, nil
	case 5:
		return OnLoading, nil
	case 6:
		return OnPayment, nil
	case 7:
		return Completed, nil
	default:
		return -1, errs.ErrInvalidDocumentStatus
	}
}

func (s DocumentStatus) FromString(value string) (DocumentStatus, error) {
	switch value {
	case "1":
		return Submitted, nil
	case "2":
		return OnReview, nil
	case "3":
		return Rejected, nil
	case "4":
		return Approved, nil
	case "5":
		return OnLoading, nil
	case "6":
		return OnPayment, nil
	case "7":
		return Completed, nil
	default:
		return -1, errs.ErrInvalidDocumentStatus
	}
}

func (s DocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value": int(s),
		"label": s.Label(),
	})
}
