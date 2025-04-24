package usecase

import (
	"fmt"
	"time"

	"github.com/lamaleka/boilerplate-golang/common/errs"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AppAuthUseCase struct {
	MsUserRepository MsUserRepository
	Jwt              *ConfJwtConfig
}

func NewAuthUseCase(msUserRepository MsUserRepository, jwt *ConfJwtConfig) *AppAuthUseCase {
	return &AppAuthUseCase{
		MsUserRepository: msUserRepository,
		Jwt:              jwt,
	}
}

func (s *AppAuthUseCase) Check(tokenString string) (*SessionData, error) {
	token, err := utils.ClaimJWT(tokenString, s.Jwt.Access.Secret)
	if err != nil {
		return nil, errs.ErrInvalidToken
	}
	claims, ok := token.Claims.(*utils.JWTClaims)
	if !ok {
		return nil, errs.ErrInvalidToken
	}
	user, errFindUser := s.MsUserRepository.FindByUsername(claims.Username)
	if errFindUser != nil {
		return nil, errFindUser
	}
	sessionData := &SessionData{
		IsAuthenticated: true,
		Username:        user.Username,
		Roles:           user.GetRoleSlugs(),
		CreatedAt:       user.CreatedAt,
		Permissions:     user.GetPermissions(),
	}
	sessionData.PreferredName = &user.Employee.Nama
	sessionData.Occupation = user.Employee.PosTitle
	return sessionData, nil
}
func (s *AppAuthUseCase) VerifySSO(username string) (*SessionData, string, error) {
	username = fmt.Sprintf("6%s", username)
	user, errFindUser := s.MsUserRepository.FindByUsername(username)
	if errFindUser != nil {
		return nil, "", errFindUser
	}
	sessionData := &SessionData{
		IsAuthenticated: true,
		Username:        user.Username,
		Roles:           user.GetRoleSlugs(),
		CreatedAt:       user.CreatedAt,
		Permissions:     user.GetPermissions(),
	}
	sessionData.PreferredName = &user.Employee.Nama
	sessionData.Occupation = user.Employee.PosTitle
	accessTokenClaims := &JWTClaims{
		Username:    user.Username,
		Roles:       user.GetRoleSlugs(),
		Permissions: user.GetPermissions(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "lorem-app",
			Audience:  jwt.ClaimStrings{"sso"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	accessToken, errGenerateAccessToken := utils.GenerateJWT(accessTokenClaims, s.Jwt.Access.Secret)
	if errGenerateAccessToken != nil {
		return nil, "", errGenerateAccessToken
	}
	return sessionData, accessToken, nil
}

func (s *AppAuthUseCase) Login(payload *LoginRequest) (*SessionData, error) {
	user, errFindUser := s.MsUserRepository.FindByUsername(payload.Username)
	if errFindUser != nil {
		return nil, errFindUser
	}
	errCompareHashAndPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if errCompareHashAndPassword != nil {
		return nil, errs.ErrInvalidUsernameOrPassword
	}
	subject := user.Username
	accessTokenClaims := &JWTClaims{
		Username:    user.Username,
		Roles:       user.GetRoleSlugs(),
		Permissions: user.GetPermissions(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "lorem-app",
			Audience:  jwt.ClaimStrings{"sso"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	accessToken, errGenerateAccessToken := utils.GenerateJWT(accessTokenClaims, s.Jwt.Access.Secret)
	if errGenerateAccessToken != nil {
		return nil, errGenerateAccessToken
	}
	refreshTokenClaims := &JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}
	refreshToken, errGenerateRefreshToken := utils.GenerateJWT(refreshTokenClaims, s.Jwt.Refresh.Secret)
	if errGenerateRefreshToken != nil {
		return nil, errGenerateRefreshToken
	}
	sessionData := &SessionData{
		IsAuthenticated: true,
		Username:        user.Username,
		Roles:           user.GetRoleSlugs(),
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		CreatedAt:       user.CreatedAt,
		Permissions:     user.GetPermissions(),
	}
	sessionData.PreferredName = &user.Employee.Nama
	sessionData.Occupation = user.Employee.PosTitle
	return sessionData, nil
}
