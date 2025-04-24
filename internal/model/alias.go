package model

import (
	"time"

	"github.com/lamaleka/boilerplate-golang/common/enum"
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/internal/entity"

	"gorm.io/gorm"
)

type DocumentType = enum.DocumentType
type DocumentStatus = enum.DocumentStatus

type DB = gorm.DB
type DefaultParams = rest.DefaultParams
type EmployeeType = enum.EmployeeType
type MsEmployee = entity.MsEmployee

type MsRole = entity.MsRole
type MsUser = entity.MsUser
type RoleType = enum.RoleType
type Time = time.Time
