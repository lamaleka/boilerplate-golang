package model

import (
	"fmt"

	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"strings"
)

type MsUserRequest struct {
	Password   string `json:"password" validate:"required,min=8"`
	EmployeeID string `json:"employee_id" validate:"required_if=RoleType 1 RoleType 2"`
	// RoleType   RoleType `json:"role_type" validate:"required,oneof=1 2 3"`
} //@name MsUserRequest

type MsUserParams struct {
	Username string `query:"username"`
	*rest.DefaultParams
}

func (payload MsUserRequest) ToMsUserEntity(employee *MsEmployee, roles []MsRole) *MsUser {
	return &MsUser{
		Username:   employee.Badge,
		EmployeeID: employee.ID,
		Employee:   employee,
		Roles:      roles,
		Password:   payload.Password,
		IsActive:   employee.IsActive,
	}
}

func (params MsUserParams) MapSearch(query *DB) {
	var conditions []string
	var args []interface{}
	addStringQuery := utils.MapStringSearchQuery(&conditions, &args)
	addStringQuery("username", params.Username)
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
