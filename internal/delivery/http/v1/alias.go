package http

import (
	"time"

	"github.com/lamaleka/boilerplate-golang/internal/entity"
	"github.com/lamaleka/boilerplate-golang/internal/model"
)

type Time = time.Time
type LoginRequest = model.LoginRequest
type CheckAuthSSORequest = model.CheckAuthSSORequest
type SessionData = model.SessionData
type MsEmployeeParams = model.MsEmployeeParams
type MsEmployeeRequest = model.MsEmployeeRequest
type MsUserParams = model.MsUserParams
type MsUserPayload = model.MsUserRequest
type MediaUploadRequest = model.MediaUploadRequest
type MediaUploadResponse = model.MediaUploadResponse
type MediaViewResponse = model.MediaViewResponse

type ConfJwtConfig = entity.ConfJwtConfig
type ConfApiSso = entity.ConfApiSso
type MsEmployee = entity.MsEmployee
type MsUser = entity.MsUser
