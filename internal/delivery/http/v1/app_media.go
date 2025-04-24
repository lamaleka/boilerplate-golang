package http

import (
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/rest"

	"github.com/labstack/echo/v4"
)

type MediaUsecase interface {
	View(fileName string) (*MediaViewResponse, error)
}

type MediaHandler struct {
	Usecase MediaUsecase
}

func NewMediaHandler(
	Usecase MediaUsecase,
) *MediaHandler {
	return &MediaHandler{
		Usecase: Usecase,
	}
}

func (h *MediaHandler) View(c echo.Context) error {
	id := c.Param("FileName")
	res, err := h.Usecase.View(id)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	c.Response().Header().Set("Content-Type", res.ContentType)
	c.Response().Header().Set("Content-Disposition", strings.Replace(res.ContentDisposition, "attachment", "inline", 1))
	_, errWriteResponse := c.Response().Write(res.FileBytes)
	return errWriteResponse
}
