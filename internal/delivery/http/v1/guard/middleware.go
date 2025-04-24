package guard

import "github.com/lamaleka/boilerplate-golang/internal/entity"

type MsUser = entity.MsUser
type ConfJwt = entity.ConfJwt
type MsUserRepository interface {
	FindByUsername(username string) (*MsUser, error)
}

type Guard struct {
	JWTAccess        *ConfJwt
	JwtRefresh       *ConfJwt
	MsUserRepository MsUserRepository
}

func NewGuard(
	jwtAccess *ConfJwt,
	jwtRefresh *ConfJwt,
	msUserRepository MsUserRepository,
) *Guard {
	return &Guard{
		JWTAccess:        jwtAccess,
		JwtRefresh:       jwtRefresh,
		MsUserRepository: msUserRepository,
	}
}
