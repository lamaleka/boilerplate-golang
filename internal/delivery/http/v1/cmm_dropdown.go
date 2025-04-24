package http

import (
	"github.com/lamaleka/boilerplate-golang/common/rest"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"github.com/labstack/echo/v4"
)

type DropdownUseCase interface {
	GetAllEmployeeUnregistered(params *MsEmployeeParams) (*[]MsEmployee, error)
	GetAllEmployee(params *MsEmployeeParams) (*[]MsEmployee, error)
}

type DropdownHandler struct {
	UseCase DropdownUseCase
}

func NewDropdownHandler(usecase DropdownUseCase) *DropdownHandler {
	return &DropdownHandler{
		UseCase: usecase,
	}
}

func (h *DropdownHandler) GetAllEmployeeUnregistered(c echo.Context) error {
	defaultParams := utils.ParseLimitAndOffset("10", "1")
	defaultParams.Keyword = c.QueryParam("keyword")
	params := &MsEmployeeParams{
		DefaultParams: defaultParams,
	}
	res, err := h.UseCase.GetAllEmployeeUnregistered(params)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Okz(c, res)
}
func (h *DropdownHandler) GetAllEmployee(c echo.Context) error {
	defaultParams := utils.ParseLimitAndOffset("10", "1")
	defaultParams.Keyword = c.QueryParam("keyword")
	params := &MsEmployeeParams{
		DefaultParams: defaultParams,
	}
	res, err := h.UseCase.GetAllEmployee(params)
	if err != nil {
		return rest.BadRequest(c, err.Error())
	}
	return rest.Okz(c, res)
}
