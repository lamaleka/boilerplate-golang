package entity

import (
	"encoding/json"
)

type MsRole struct {
	ID          uint           `json:"-"`
	UUID        UniqueID       `gorm:"type:uniqueidentifier;default:NEWSEQUENTIALID()" json:"id"`
	Slug        string         `json:"slug"`
	Name        string         `json:"name"`
	Permissions []MsPermission `gorm:"many2many:piv_role_permission" json:"permissions"`
	CreatedAt   *Time          `json:"created_at"`
	UpdatedAt   *Time          `json:"updated_at"`
	DeletedAt   *DeletedAt     `json:"-"`
}

func (MsRole) TableName() string {
	return "ms_role"
}

func (r MsRole) MarshalJSON() ([]byte, error) {
	type Alias MsRole
	return json.Marshal(&struct {
		Alias
		Permissions []string `json:"permissions"`
	}{
		Alias:       Alias(r),
		Permissions: extractPermissionNames(r.Permissions),
	})
}

func extractPermissionNames(permissions []MsPermission) []string {
	names := make([]string, len(permissions))
	for i, permission := range permissions {
		names[i] = permission.Name
	}
	return names
}
