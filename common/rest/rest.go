package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	Error      bool        `json:"error,omitempty"`
	Code       int         `json:"code" validate:"required"`
	Message    string      `json:"message" validate:"required"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
} //@name Rest

type DefaultParams struct {
	Offset  int
	Keyword string
	OrderBy string `query:"order_by"`
	Order   string `query:"order" validate:"omitempty,oneof=asc desc"`
	*Pagination
}
type Pagination struct {
	PageNumber int   `json:"page_number"`
	PageSize   int   `json:"page_size"`
	TotalData  int64 `json:"total_data"`
}

func Ok(c echo.Context) error {
	return c.JSON(http.StatusOK, Rest{
		Code:    http.StatusOK,
		Message: "OK",
	})
}

func Blob(c echo.Context, data []byte) error {
	return c.Blob(http.StatusOK, "application/octet-stream", data)
}
func PDF(c echo.Context, data []byte) error {
	return c.Blob(http.StatusOK, "application/pdf", data)
}
func MediaData(c echo.Context, contentType string, data []byte) error {
	return c.Blob(http.StatusOK, contentType, data)
}

func Deleted(c echo.Context) error {
	return c.JSON(http.StatusOK, Rest{
		Code:    http.StatusOK,
		Message: "Berhasil Dihapus",
	})
}
func Okz(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Rest{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    data,
	})
}
func Paginate(c echo.Context, data interface{}, params *DefaultParams) error {
	return c.JSON(http.StatusOK, Rest{
		Code:    http.StatusOK,
		Message: "OK",
		Pagination: &Pagination{
			PageSize:   params.PageSize,
			PageNumber: params.PageNumber,
			TotalData:  params.TotalData,
		},
		Data: data,
	})
}
func BadRequest(c echo.Context, err string) error {
	return c.JSON(http.StatusBadRequest, Rest{
		Error:   true,
		Code:    http.StatusBadRequest,
		Message: err,
	})
}
func Conflict(c echo.Context, err string) error {
	return c.JSON(http.StatusConflict, Rest{
		Error:   true,
		Code:    http.StatusConflict,
		Message: err,
	})
}
func ServerError(c echo.Context, err string) error {
	return c.JSON(http.StatusInternalServerError, Rest{
		Error:   true,
		Code:    http.StatusInternalServerError,
		Message: err,
	})
}
func NotFound(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, Rest{
		Error:   true,
		Code:    http.StatusNotFound,
		Message: "Url " + c.Path() + " tidak ditemukan",
	})
}
func UnAuth(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, Rest{
		Error:   true,
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	})
}
func ReachLimit(c echo.Context) error {
	return c.JSON(http.StatusTooManyRequests, Rest{
		Error:   true,
		Code:    http.StatusTooManyRequests,
		Message: "limit request reached, try again later",
	})
}
func Forbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, Rest{
		Error:   true,
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	})
}
