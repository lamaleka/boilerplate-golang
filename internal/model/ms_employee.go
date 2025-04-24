package model

import (
	"fmt"
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/enum"
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"gorm.io/gorm"
)

type MsEmployeeRequest struct {
	Nama      string  `json:"nama"`
	Badge     string  `json:"badge"`
	DeptID    *string `json:"dept_id,omitempty"`
	DeptTitle string  `json:"dept_title"`
	Email     *string `json:"email,omitempty" validate:"email"`
	PosID     *string `json:"pos_id,omitempty"`
	PosTitle  string  `json:"pos_title"`
} //@name MsEmployeeRequest

type MsEmployeeParams struct {
	Nama      string `query:"nama"`
	Badge     string `query:"badge"`
	DeptTitle string `query:"dept_title"`
	Email     string `query:"email"`
	PosTitle  string `query:"pos_title"`
	IsActive  *int   `query:"is_active"`
	*rest.DefaultParams
}

func (payload MsEmployeeRequest) ToMsEmployeeEntity() *MsEmployee {
	return &MsEmployee{
		Nama:         payload.Nama,
		Badge:        payload.Badge,
		DeptID:       payload.DeptID,
		DeptTitle:    payload.DeptTitle,
		Email:        payload.Email,
		PosID:        payload.PosID,
		PosTitle:     payload.PosTitle,
		EmployeeType: enum.TKO,
		IsActive:     true,
	}
}

func (params MsEmployeeParams) MapSearch(query *gorm.DB) {
	var conditions []string
	var args []interface{}
	addStringQuery := utils.MapStringSearchQuery(&conditions, &args)
	addStringQuery("nama", params.Nama)
	addStringQuery("badge", params.Badge)
	addStringQuery("dept_title", params.DeptTitle)
	addStringQuery("email", params.Email)
	addStringQuery("pos_title", params.PosTitle)
	if len(conditions) > 0 {
		query.Where(strings.Join(conditions, " AND "), args...)
	}
	conditions = conditions[:0]
	args = args[:0]
	addNumberQuery := utils.MapNumericSearchQuery(&conditions, &args)
	addNumberQuery("is_active", params.IsActive)
	if len(conditions) > 0 {
		query.Where(strings.Join(conditions, " AND "), args...)
	}
	query.Count(&params.TotalData)
	if params.OrderBy != "" && params.Order != "" {
		query.Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order))
	} else {
		query.Order("created_at desc")
	}
}
