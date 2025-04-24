package http

import (
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/labstack/echo/v4"
)

type MsEmployeeUseCase interface {
	GetAll(params *MsEmployeeParams) (*[]MsEmployee, error)
	Detail(uuid string) (*MsEmployee, error)
	Create(payload *MsEmployeeRequest) (*MsEmployee, error)
	Update(uuid string, payload *MsEmployeeRequest) (*MsEmployee, error)
	UpdateStatus(uuid string) error
	Delete(uuid string) error
}

type MsEmployeeHandler struct {
	UseCase MsEmployeeUseCase
}

func NewMsEmployeeHandler(usecase MsEmployeeUseCase) *MsEmployeeHandler {
	return &MsEmployeeHandler{
		UseCase: usecase,
	}
}

func (h *MsEmployeeHandler) GetAll(c echo.Context) error {
	pageSize := c.QueryParam("page_size")
	pageNumber := c.QueryParam("page_number")
	defaultParams := utils.ParseLimitAndOffset(pageSize, pageNumber)
	defaultParams.Keyword = c.QueryParam("keyword")
	params := &MsEmployeeParams{
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

func (h *MsEmployeeHandler) Detail(c echo.Context) error {
	uuid := c.Param("ID")
	res, err := h.UseCase.Detail(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Okz(c, res)
}

func (h *MsEmployeeHandler) Create(c echo.Context) error {
	payload := new(MsEmployeeRequest)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	data, err := h.UseCase.Create(payload)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}

	return rest.Okz(c, data)
}

func (h *MsEmployeeHandler) Update(c echo.Context) error {
	uuid := c.Param("ID")
	payload := new(MsEmployeeRequest)
	if err := utils.BindAndValidate(c, payload); err != nil {
		return rest.BadRequest(c, err.Error())
	}
	data, err := h.UseCase.Update(uuid, payload)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}

	return rest.Okz(c, data)
}

func (h *MsEmployeeHandler) UpdateStatus(c echo.Context) error {
	uuid := c.Param("ID")
	err := h.UseCase.UpdateStatus(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Ok(c)
}

func (h *MsEmployeeHandler) Delete(c echo.Context) error {
	uuid := c.Param("ID")
	err := h.UseCase.Delete(uuid)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Deleted(c)
}
