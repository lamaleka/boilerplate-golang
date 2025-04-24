package usecase

import (
	"github.com/lamaleka/boilerplate-golang/common/utils"
)

type MsEmployeeUseCase struct {
	Repo MsEmployeeRepository
}

func NewMsEmployeeUseCase(repo MsEmployeeRepository) *MsEmployeeUseCase {
	return &MsEmployeeUseCase{
		Repo: repo,
	}
}

func (s *MsEmployeeUseCase) GetAll(params *MsEmployeeParams) (*[]MsEmployee, error) {
	data, err := s.Repo.GetAll(params)
	params.PageSize = utils.CalculatePageSize(len(*data), params.PageSize)
	return data, err
}
func (s *MsEmployeeUseCase) Detail(uuid string) (*MsEmployee, error) {
	data, err := s.Repo.FindByUUID(uuid)
	return data, err
}
func (s *MsEmployeeUseCase) Create(payload *MsEmployeeRequest) (*MsEmployee, error) {
	data := payload.ToMsEmployeeEntity()
	employee, err := s.Repo.Create(data)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *MsEmployeeUseCase) Update(uuid string, payload *MsEmployeeRequest) (*MsEmployee, error) {
	oldData, err := s.Repo.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	newData := payload.ToMsEmployeeEntity()
	return s.Repo.Update(oldData, newData)
}
func (s *MsEmployeeUseCase) UpdateStatus(uuid string) error {
	return s.Repo.UpdateStatus(uuid)
}
func (s *MsEmployeeUseCase) Delete(uuid string) error {
	return s.Repo.Delete(uuid)
}
