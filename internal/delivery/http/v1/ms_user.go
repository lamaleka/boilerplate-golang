package http

import (
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/labstack/echo/v4"
)

type MsUserUseCase interface {
	GetAll(params *MsUserParams) (*[]MsUser, error)
	Detail(uuid string) (*MsUser, error)
	ResetPassword(uuid string) (*MsUser, error)
	Create(payload *MsUserPayload) (*MsUser, error)
	Update(uuid string, payload *MsUserPayload) (*MsUser, error)
	Delete(uuid string) error
}

type MsUserHandler struct {
	UseCase MsUserUseCase
}

func NewMsUserHandler(usecase MsUserUseCase) *MsUserHandler {
	return &MsUserHandler{
		UseCase: usecase,
	}
}

func (h *MsUserHandler) GetAll(c echo.Context) error {
	pageSize := c.QueryParam("page_size")
	pageNumber := c.QueryParam("page_number")
	defaultParams := utils.ParseLimitAndOffset(pageSize, pageNumber)
	defaultParams.Keyword = c.QueryParam("keyword")
	params := &MsUserParams{
		DefaultParams: defaultParams,
	}
	if err := c.Bind(params); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	res, err := h.UseCase.GetAll(params)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Paginate(c, res, params.DefaultParams)
}

func (h *MsUserHandler) Detail(c echo.Context) error {
	uuid := c.Param("ID")
	res, err := h.UseCase.Detail(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Okz(c, res)
}

func (h *MsUserHandler) Create(c echo.Context) error {
	payload := new(MsUserPayload)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	data, err := h.UseCase.Create(payload)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}

	return rest.Okz(c, data)
}

func (h *MsUserHandler) ResetPassword(c echo.Context) error {
	uuid := c.Param("ID")
	data, err := h.UseCase.ResetPassword(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Okz(c, data)
}
func (h *MsUserHandler) Update(c echo.Context) error {
	uuid := c.Param("ID")
	payload := new(MsUserPayload)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	data, err := h.UseCase.Update(uuid, payload)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}

	return rest.Okz(c, data)
}

func (h *MsUserHandler) Delete(c echo.Context) error {
	uuid := c.Param("ID")
	err := h.UseCase.Delete(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Deleted(c)
}
