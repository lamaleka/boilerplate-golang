package http

import (
	"fmt"
	"net/http"

	"github.com/lamaleka/boilerplate-golang/common/errs"
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/labstack/echo/v4"
)

type AuthUseCase interface {
	Login(payload *LoginRequest) (*SessionData, error)
	VerifySSO(username string) (*SessionData, string, error)
	Check(tokenString string) (*SessionData, error)
}

type AuthHandler struct {
	UseCase    AuthUseCase
	ConfSsoApi *ConfApiSso
}

func NewAuthHandler(usecase AuthUseCase, confSsoApi *ConfApiSso) *AuthHandler {
	return &AuthHandler{
		UseCase:    usecase,
		ConfSsoApi: confSsoApi,
	}
}

func (h *AuthHandler) VerifySSO(c echo.Context) error {
	accessToken, errGetAccessToken := c.Cookie("access_token")
	if errGetAccessToken != nil {
		return rest.BadRequest(c, errGetAccessToken.Error())
	}
	refreshToken, errGetRefreshToken := c.Cookie("refresh_token")
	if errGetRefreshToken != nil {
		return rest.BadRequest(c, errGetRefreshToken.Error())
	}
	claims, err := utils.DecodeJWT(accessToken.Value)
	if err != nil {
		return rest.BadRequest(c, errs.ErrInvalidToken.Error())
	}
	payload := &CheckAuthSSORequest{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
		Host:         c.Request().Host,
		UserAgent:    c.Request().UserAgent(),
	}
	if err := h.CheckSSO(payload); err != nil {
		return rest.UnAuth(c)
	}
	sessionData, appAccessToken, err := h.UseCase.VerifySSO(claims.PreferredUsername)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	sessionData.AccessToken = accessToken.Value
	sessionData.RefreshToken = refreshToken.Value
	c.SetCookie(&http.Cookie{
		Name:     "app_access_token",
		Value:    appAccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
	return rest.Okz(c, sessionData)
}

func (h *AuthHandler) Check(c echo.Context) error {
	accessToken, errGetAccessToken := c.Cookie("access_token")
	if errGetAccessToken != nil {
		return rest.BadRequest(c, errGetAccessToken.Error())
	}
	refreshToken, errGetRefreshToken := c.Cookie("refresh_token")
	if errGetRefreshToken != nil {
		return rest.BadRequest(c, errGetRefreshToken.Error())
	}
	tokenString := accessToken.Value
	appAccessToken, errGetSSO := c.Cookie("app_access_token")
	if errGetSSO == nil {
		payload := &CheckAuthSSORequest{
			AccessToken:  accessToken.Value,
			RefreshToken: refreshToken.Value,
			Host:         c.Request().Host,
			UserAgent:    c.Request().UserAgent(),
		}
		if err := h.CheckSSO(payload); err != nil {
			return rest.UnAuth(c)
		}
		tokenString = appAccessToken.Value
	}
	sessionData, err := h.UseCase.Check(tokenString)
	if err != nil {
		return rest.UnAuth(c)
	}
	sessionData.AccessToken = accessToken.Value
	sessionData.RefreshToken = refreshToken.Value
	return rest.Okz(c, sessionData)
}
func (h *AuthHandler) CheckSSO(payload *CheckAuthSSORequest) error {
	url := fmt.Sprintf("%s?%s", h.ConfSsoApi.Url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", fmt.Sprintf("access_token=%s;refresh_token=%s", payload.AccessToken, payload.RefreshToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Host", payload.Host)
	req.Header.Set("User-Agent", payload.UserAgent)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return errs.ErrUnauthorized
	}
	return nil
}
func (h *AuthHandler) Login(c echo.Context) error {
	payload := new(LoginRequest)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	sessionData, err := h.UseCase.Login(payload)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    sessionData.AccessToken,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		// MaxAge:   3600 * 24 * 1,
	})
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    sessionData.RefreshToken,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		// MaxAge:   3600 * 24 * 3,
	})
	return rest.Okz(c, sessionData)
}
