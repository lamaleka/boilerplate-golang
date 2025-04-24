package guard

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (m *Guard) NewPermissionGuard(requiredPermissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var tokenString string
			appAccessToken, errGetAppAccessToken := c.Cookie("app_access_token")
			if errGetAppAccessToken != nil {
				accessToken, errGetAccessToken := c.Cookie("access_token")
				if errGetAccessToken != nil {
					return rest.UnAuth(c)
				}
				tokenString = accessToken.Value
			} else {
				tokenString = appAccessToken.Value
			}
			var claims *utils.JWTClaims
			token, err := utils.ClaimJWT(tokenString, m.JWTAccess.Secret)
			if err != nil {
				refreshToken, errGetRefreshToken := c.Cookie("refresh_token")
				if errGetRefreshToken != nil {
					return rest.UnAuth(c)
				}
				newClaims, newAccessToken := m.GenerateFromRefreshToken(refreshToken.Value)
				if newClaims == nil {
					return rest.UnAuth(c)
				}
				c.SetCookie(&http.Cookie{
					Name:     "access_token",
					Value:    newAccessToken,
					Path:     "/",
					HttpOnly: false,
					Secure:   true,
					SameSite: http.SameSiteLaxMode,
				})
				claims = newClaims
			} else {
				newClaims, ok := token.Claims.(*utils.JWTClaims)
				if !ok {
					return rest.UnAuth(c)
				}
				claims = newClaims
			}
			c.Set("Subject", "")
			permissionMap := make(map[string]bool)
			for _, permission := range claims.Permissions {
				permissionMap[permission] = true
			}
			if claims.SSO == 0 {
				c.Set("Subject", claims.Subject)
			}
			for _, requiredPermission := range requiredPermissions {
				if permissionMap[requiredPermission] {
					return next(c)
				}
			}
			return rest.Forbidden(c)
		}
	}
}
func (m *Guard) GenerateFromRefreshToken(refreshToken string) (*utils.JWTClaims, string) {
	token, err := utils.ClaimJWT(refreshToken, m.JwtRefresh.Secret)
	if err != nil {
		return nil, ""
	}
	claims, ok := token.Claims.(*utils.JWTClaims)
	if !ok {
		return nil, ""
	}
	user, errFindUser := m.MsUserRepository.FindByUsername(claims.Username)
	if errFindUser != nil {
		return nil, ""
	}
	subject := user.Username
	newAccessTokenClaims := &utils.JWTClaims{
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
	newAccessToken, errGenerateAccessToken := utils.GenerateJWT(newAccessTokenClaims, m.JWTAccess.Secret)
	if errGenerateAccessToken != nil {
		return nil, ""
	}
	fmt.Println("Generated New Token")
	return newAccessTokenClaims, newAccessToken

}
