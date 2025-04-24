package routes

import (
	"fmt"

	"github.com/lamaleka/boilerplate-golang/internal/delivery/http/v1"
	"github.com/lamaleka/boilerplate-golang/internal/delivery/http/v1/guard"

	"time"

	"github.com/labstack/echo/v4"
)

var startTime time.Time

type RouteConfig struct {
	App               *echo.Echo
	Guard             *guard.Guard
	AuthHandler       *http.AuthHandler
	MediaHandler      *http.MediaHandler
	DropdownHandler   *http.DropdownHandler
	MsUserHandler     *http.MsUserHandler
	MsEmployeeHandler *http.MsEmployeeHandler
}

func (c *RouteConfig) Setup() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := time.Now().In(loc)
	c.App.GET("/", func(c echo.Context) error {
		uptime := time.Since(startTime).Abs().Hours()
		return c.JSON(200, map[string]string{
			"app_version": "0.0.1+1",
			"started_at":  startTime.Format(time.RFC822),
			"uptime":      fmt.Sprintf("%.2f", uptime),
		})
	})

	v1 := c.App.Group("/api/v1")

	guestRoute := v1.Group("")
	c.SetupGuestRoute(guestRoute)

	commonGuard := c.Guard.NewPermissionGuard("admin_access", "officer_access", "vendor_access")
	commonRoute := v1.Group("/common")
	commonRoute.Use(commonGuard)
	c.SetupCommonRoute(guestRoute)

	adminGuard := c.Guard.NewPermissionGuard("admin_access")
	adminRoute := v1.Group("/admin")
	adminRoute.Use(adminGuard)
	c.SetupAdminRoute(adminRoute)

	officerGuard := c.Guard.NewPermissionGuard("officer_access")
	officerRoute := v1.Group("/officer")
	officerRoute.Use(officerGuard)
	c.SetupOfficerRoute(officerRoute)

	buyerGuard := c.Guard.NewPermissionGuard("buyer_access")
	buyerRoute := v1.Group("/buyer")
	buyerRoute.Use(buyerGuard)
	c.SetupBuyerRoute(buyerRoute)

}

func (c *RouteConfig) SetupGuestRoute(group *echo.Group) {
	// Auth
	auth := group.Group("/auth")
	auth.POST("/login", c.AuthHandler.Login)
	auth.POST("/verify-sso", c.AuthHandler.VerifySSO)
	auth.GET("/check", c.AuthHandler.Check)
}

func (c *RouteConfig) SetupCommonRoute(group *echo.Group) {
	// Media
	media := group.Group("/media")
	media.GET("/:FileName", c.MediaHandler.View)
}

func (c *RouteConfig) SetupAdminRoute(group *echo.Group) {
	// Dropdown
	dropdown := group.Group("/dropdown")
	dropdown.GET("/employee", c.DropdownHandler.GetAllEmployee)
	dropdown.GET("/employee/unregistered", c.DropdownHandler.GetAllEmployeeUnregistered)

	// Master Data
	masters := group.Group("/masters")

	// Master Data User
	msUser := masters.Group("/user")
	msUser.GET("", c.MsUserHandler.GetAll)
	msUser.POST("", c.MsUserHandler.Create)
	msUser.GET("/:ID", c.MsUserHandler.Detail)
	msUser.PUT("/:ID/reset-password", c.MsUserHandler.ResetPassword)
	msUser.PUT("/:ID", c.MsUserHandler.Update)
	msUser.DELETE("/:ID", c.MsUserHandler.Delete)

	// Master Data Karyawan
	msEmployee := masters.Group("/employee")
	msEmployee.GET("", c.MsEmployeeHandler.GetAll)
	msEmployee.POST("", c.MsEmployeeHandler.Create)
	msEmployee.GET("/:ID", c.MsEmployeeHandler.Detail)
	msEmployee.PUT("/:ID", c.MsEmployeeHandler.Update)
	msEmployee.PATCH("/:ID/status", c.MsEmployeeHandler.UpdateStatus)

	// Trx Document
	group.Group("/document")
}

func (c *RouteConfig) SetupOfficerRoute(group *echo.Group) {
	group.Group("/document")

}
func (c *RouteConfig) SetupBuyerRoute(group *echo.Group) {
	group.Group("/document")
}
