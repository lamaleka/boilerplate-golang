package usecase

import (
	"time"

	"github.com/lamaleka/boilerplate-golang/common/utils"
	"github.com/lamaleka/boilerplate-golang/internal/entity"
	"github.com/lamaleka/boilerplate-golang/internal/model"
)

type Time = time.Time
type Month = time.Month
type JWTClaims = utils.JWTClaims

type LoginRequest = model.LoginRequest
type CheckAuthResponse = model.CheckAuthResponse
type SessionData = model.SessionData
type MsUserParams = model.MsUserParams
type MsUserRequest = model.MsUserRequest
type MsRoleParams = model.MsRoleParams
type MsRolePayload = model.MsRoleRequest
type MsEmployeeParams = model.MsEmployeeParams
type MsEmployeeRequest = model.MsEmployeeRequest

type ConfJwtConfig = entity.ConfJwtConfig
type MsEmployee = entity.MsEmployee
type MsRole = entity.MsRole
type MsUser = entity.MsUser

type MediaUploadRequest = model.MediaUploadRequest
type MediaUploadResponse = model.MediaUploadResponse
type MediaViewResponse = model.MediaViewResponse

type MsEmployeeRepository interface {
	GetAll(params *MsEmployeeParams) (*[]MsEmployee, error)
	GetAllUnregistered(params *MsEmployeeParams) (*[]MsEmployee, error)
	FindByBadge(badge string) (*MsEmployee, error)
	FirstOrCreate(data *MsEmployee) (*MsEmployee, error)
	FindByUUID(uuid string) (*MsEmployee, error)
	FindByUUIDs(uuids []string) ([]MsEmployee, error)
	Create(data *MsEmployee) (*MsEmployee, error)
	Update(oldData *MsEmployee, newData *MsEmployee) (*MsEmployee, error)
	UpdateStatus(uuid string) error
	Delete(uuid string) error
}

type MsRoleRepository interface {
	GetAll(params *MsRoleParams) (*[]MsRole, error)
	FindByUUID(uuid string) (*MsRole, error)
	FindByUUIDs(uuids []string) ([]MsRole, error)
	FindByIDs(ids []int) ([]MsRole, error)
	Create(data *MsRole) (*MsRole, error)
	Update(oldData *MsRole, newData *MsRole) (*MsRole, error)
	Delete(uuid string) error
}

type MsUserRepository interface {
	GetAll(params *MsUserParams) (*[]MsUser, error)
	FindByUsername(username string) (*MsUser, error)
	FindByUUID(uuid string) (*MsUser, error)
	FindByUUIDs(uuids []string) ([]MsUser, error)
	ResetPassword(data *MsUser, newPassword string) (*MsUser, error)
	Create(data *MsUser) (*MsUser, error)
	Update(oldData *MsUser, newData *MsUser) (*MsUser, error)
	Delete(uuid string) error
}

type ApiWebdavUsecase interface {
	Upload(payload *MediaUploadRequest) (*MediaUploadResponse, error)
	View(fileName string) (*MediaViewResponse, error)
}
