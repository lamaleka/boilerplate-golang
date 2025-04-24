package app

import (
	"github.com/lamaleka/boilerplate-golang/internal/config"
	"github.com/lamaleka/boilerplate-golang/internal/delivery/http/v1"
	"github.com/lamaleka/boilerplate-golang/internal/delivery/http/v1/guard"
	"github.com/lamaleka/boilerplate-golang/internal/delivery/http/v1/routes"
	"github.com/lamaleka/boilerplate-golang/internal/entity"
	"github.com/lamaleka/boilerplate-golang/internal/repository"
	"github.com/lamaleka/boilerplate-golang/internal/usecase"
	"github.com/lamaleka/boilerplate-golang/pkg/webdav"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	App       *echo.Echo
	Log       *logrus.Logger
	Validator *config.CValidator
	Viper     *entity.ConfViper
}

func Bootstrap(config *BootstrapConfig) {
	config.App.Validator = config.Validator

	webdavUsecase := webdav.NewWebdavUseCase(config.Viper.Api.Webdav)

	mediaUsecase := usecase.NewMediaUseCase(webdavUsecase)
	mediaHandler := http.NewMediaHandler(mediaUsecase)

	msEmployeeRepository := repository.NewMsEmployeeRepository(config.DB)
	msEmployeeUseCase := usecase.NewMsEmployeeUseCase(msEmployeeRepository)
	msEmployeeHandler := http.NewMsEmployeeHandler(msEmployeeUseCase)

	msRoleRepository := repository.NewMsRoleRepository(config.DB)
	usecase.NewMsRoleUseCase(msRoleRepository)

	msUserRepository := repository.NewMsUserRepository(config.DB, msEmployeeRepository)
	msUserUseCase := usecase.NewMsUserUseCase(msUserRepository, msEmployeeRepository, msRoleRepository)
	msUserHandler := http.NewMsUserHandler(msUserUseCase)

	guard := guard.NewGuard(config.Viper.Jwt.Access, config.Viper.Jwt.Refresh, msUserRepository)

	dropdownUsecase := usecase.NewDropdownUseCase(msEmployeeRepository)
	dropdownHandler := http.NewDropdownHandler(dropdownUsecase)

	authUseCase := usecase.NewAuthUseCase(msUserRepository, config.Viper.Jwt)
	authHandler := http.NewAuthHandler(authUseCase, config.Viper.Api.Sso)

	routeConfig := routes.RouteConfig{
		App:               config.App,
		AuthHandler:       authHandler,
		Guard:             guard,
		DropdownHandler:   dropdownHandler,
		MediaHandler:      mediaHandler,
		MsUserHandler:     msUserHandler,
		MsEmployeeHandler: msEmployeeHandler,
	}
	routeConfig.Setup()
}
