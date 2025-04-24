package usecase

import (
	"github.com/lamaleka/boilerplate-golang/common/enum"
	"github.com/lamaleka/boilerplate-golang/common/errs"
	"github.com/lamaleka/boilerplate-golang/common/utils"

	"golang.org/x/crypto/bcrypt"
)

type MsUserUseCase struct {
	Repository           MsUserRepository
	MsEmployeeRepository MsEmployeeRepository
	MsRoleRepository     MsRoleRepository
}

func NewMsUserUseCase(
	repository MsUserRepository,
	msEmployeeRepository MsEmployeeRepository,
	msRoleRepository MsRoleRepository,
) *MsUserUseCase {
	return &MsUserUseCase{
		Repository:           repository,
		MsEmployeeRepository: msEmployeeRepository,
		MsRoleRepository:     msRoleRepository,
	}
}

func (u *MsUserUseCase) GetAll(params *MsUserParams) (*[]MsUser, error) {
	data, err := u.Repository.GetAll(params)
	params.PageSize = utils.CalculatePageSize(len(*data), params.PageSize)
	return data, err
}
func (u *MsUserUseCase) Detail(uuid string) (*MsUser, error) {
	data, err := u.Repository.FindByUUID(uuid)
	return data, err
}

func (u *MsUserUseCase) Create(payload *MsUserRequest) (*MsUser, error) {

	newPassword, errGeneratePassword := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if errGeneratePassword != nil {
		return nil, errs.ErrEncryptPassword
	}
	payload.Password = string(newPassword)

	employee, err := u.MsEmployeeRepository.FindByUUID(payload.EmployeeID)
	if err != nil {
		return nil, err
	}

	roles, err := u.MsRoleRepository.FindByIDs([]int{enum.Officer.Value()})
	if err != nil {
		return nil, err
	}
	data := payload.ToMsUserEntity(employee, roles)
	user, err := u.Repository.Create(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *MsUserUseCase) Update(uuid string, payload *MsUserRequest) (*MsUser, error) {
	oldData, err := u.Repository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	var employee *MsEmployee
	foundEmployee, err := u.MsEmployeeRepository.FindByUUID(payload.EmployeeID)
	if err != nil {
		return nil, err
	}
	employee = foundEmployee
	roles, err := u.MsRoleRepository.FindByIDs([]int{enum.Officer.Value()})
	if err != nil {
		return nil, err
	}
	newData := payload.ToMsUserEntity(employee, roles)
	return u.Repository.Update(oldData, newData)
}
func (u *MsUserUseCase) Delete(uuid string) error {
	return u.Repository.Delete(uuid)
}

func (u *MsUserUseCase) ResetPassword(uuid string) (*MsUser, error) {
	oldData, err := u.Repository.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	newPassword, errGeneratePassword := bcrypt.GenerateFromPassword([]byte(oldData.Username), 12)
	if errGeneratePassword != nil {
		return nil, errs.ErrEncryptPassword
	}
	return u.Repository.ResetPassword(oldData, string(newPassword))
}
