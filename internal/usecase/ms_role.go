package usecase

type MsRoleUseCase struct {
	Repository MsRoleRepository
}

func NewMsRoleUseCase(repository MsRoleRepository) *MsRoleUseCase {
	return &MsRoleUseCase{
		Repository: repository,
	}
}
