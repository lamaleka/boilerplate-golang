package entity

import (
	"time"

	"github.com/lamaleka/boilerplate-golang/common/enum"

	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

type DocumentStatus = enum.DocumentStatus
type DocumentType = enum.DocumentType
type EmployeeType = enum.EmployeeType

type UniqueID = mssql.UniqueIdentifier
type Time = time.Time
type DeletedAt = gorm.DeletedAt
type DB = gorm.DB
