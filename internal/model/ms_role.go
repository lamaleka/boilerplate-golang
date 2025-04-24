package model

type MsRoleRequest struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type MsRoleParams struct {
	*DefaultParams
}

func (payload MsRoleRequest) ToMsRoleEntity() *MsRole {
	return &MsRole{
		Slug: payload.Slug,
		Name: payload.Name,
	}
}
