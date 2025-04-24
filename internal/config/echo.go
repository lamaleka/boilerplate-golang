package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewEcho(logger *logrus.Logger) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = NewErrorHandler()
	e.Use(middleware.Logger())
	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogURI:    true,
	// 	LogStatus: true,

	// 	LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
	// 		logger.WithFields(logrus.Fields{
	// 			"URI":    values.URI,
	// 			"status": values.Status,
	// 		}).Info("request")

	// 		return nil
	// 	},
	// }))
	e.Use(middleware.Recover())
	return e
}
func NewErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.JSON(code, map[string]interface{}{
			"error":   true,
			"code":    code,
			"message": err.Error(),
		})
	}
}
