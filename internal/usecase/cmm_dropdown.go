package usecase

type BIUseCase struct {
	Employee MsEmployeeRepository
}

func NewDropdownUseCase(

	employee MsEmployeeRepository,
) *BIUseCase {
	return &BIUseCase{
		Employee: employee,
	}
}

func (s *BIUseCase) GetAllEmployeeUnregistered(params *MsEmployeeParams) (*[]MsEmployee, error) {
	return s.Employee.GetAllUnregistered(params)
}

func (s *BIUseCase) GetAllEmployee(params *MsEmployeeParams) (*[]MsEmployee, error) {
	return s.Employee.GetAll(params)
}
