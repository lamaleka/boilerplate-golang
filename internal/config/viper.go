package config

import (
	"fmt"

	"github.com/lamaleka/boilerplate-golang/common/utils"
	"github.com/lamaleka/boilerplate-golang/internal/entity"

	"github.com/spf13/viper"
)

func NewViper() *entity.ConfViper {
	viper := viper.New()
	deploymentMode := utils.DeploymentMode()
	fmt.Println("\033[34mUsing:", deploymentMode.Label(), "\033[0m")
	viper.SetConfigName(deploymentMode.File())
	viper.SetConfigType("json")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viperConfig := &entity.ConfViper{
		App: &entity.ConfApp{
			AppName: viper.GetString("app.name"),
		},
		Web: &entity.ConfWeb{
			Prefork: viper.GetBool("web.prefork"),
			Port:    viper.GetInt("web.port"),
		},
		Log: &entity.ConfLog{
			Level: viper.GetInt("log.level"),
		},
		Jwt: &entity.ConfJwtConfig{
			Access: &entity.ConfJwt{
				Secret:  viper.GetString("jwt.access.secret"),
				Expired: viper.GetInt("jwt.access.expired"),
				MaxAge:  viper.GetInt("jwt.access.maxAge"),
			},
			Refresh: &entity.ConfJwt{
				Secret:  viper.GetString("jwt.refresh.secret"),
				Expired: viper.GetInt("jwt.refresh.expired"),
				MaxAge:  viper.GetInt("jwt.refresh.maxAge"),
			},
		},
		Db: &entity.ConfDbConfig{
			App: &entity.ConfDb{
				Host:     viper.GetString("database.app.host"),
				Port:     viper.GetInt("database.app.port"),
				Username: viper.GetString("database.app.username"),
				Password: viper.GetString("database.app.password"),
				Name:     viper.GetString("database.app.name"),
				Pool: &entity.ConfDbPool{
					Idle:     viper.GetInt("database.app.pool.idle"),
					Max:      viper.GetInt("database.app.pool.max"),
					Lifetime: viper.GetInt("database.app.pool.lifetime"),
				},
			},
			Bi: &entity.ConfDb{
				Host:     viper.GetString("database.bi.host"),
				Port:     viper.GetInt("database.bi.port"),
				Username: viper.GetString("database.bi.username"),
				Password: viper.GetString("database.bi.password"),
				Name:     viper.GetString("database.bi.name"),
				Pool: &entity.ConfDbPool{
					Idle:     viper.GetInt("database.bi.pool.idle"),
					Max:      viper.GetInt("database.bi.pool.max"),
					Lifetime: viper.GetInt("database.bi.pool.lifetime"),
				},
			},
		},
		Api: &entity.ConfApiConfig{
			Sso: &entity.ConfApiSso{
				Url:    viper.GetString("api.sso.url"),
				Secret: viper.GetString("api.employee.secret"),
			},
			Webdav: &entity.ConfApiWebdav{
				Url:    viper.GetString("api.webdav.url"),
				User:   viper.GetString("api.webdav.user"),
				Secret: viper.GetString("api.webdav.secret"),
				Path:   viper.GetString("api.webdav.path"),
			},
		},
	}

	return viperConfig
}
