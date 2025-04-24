package repository

import (
	"time"

	"github.com/lamaleka/boilerplate-golang/internal/entity"
	"github.com/lamaleka/boilerplate-golang/internal/model"

	"gorm.io/gorm"
)

type DB = gorm.DB
type Month = time.Month

type MsEmployee = entity.MsEmployee
type MsRole = entity.MsRole
type MsUser = entity.MsUser

type MsEmployeeParams = model.MsEmployeeParams
type MsRoleParams = model.MsRoleParams
type MsUserParams = model.MsUserParams
