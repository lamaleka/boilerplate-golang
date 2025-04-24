package entity

import (
	"encoding/json"
)

type MsUser struct {
	ID         uint        `json:"-"`
	UUID       UniqueID    `gorm:"type:uniqueidentifier;default:NEWSEQUENTIALID()" json:"id"`
	Username   string      `json:"username"`
	EmployeeID uint        `json:"-"`
	Employee   *MsEmployee `gorm:"foreign_key:EmployeeID" json:"employee,omitempty"`
	Roles      []MsRole    `gorm:"many2many:piv_user_role" json:"roles"`
	Password   string      `json:"-"`
	IsActive   bool        `json:"is_active"`
	CreatedAt  *Time       `json:"created_at"`
	UpdatedAt  *Time       `json:"updated_at"`
	DeletedAt  *DeletedAt  `json:"-"`
}

func (MsUser) TableName() string {
	return "ms_user"
}

func (u MsUser) BeforeDelete(tx *DB) (err error) {
	err = tx.Exec("DELETE FROM piv_user_role WHERE ms_user_id = ?", u.ID).Error
	return
}

func (m MsUser) GetRoles() []string {
	return extractRoleNames(m.Roles)
}

func (m MsUser) GetRoleSlugs() []string {
	return extractRoleSlugs(m.Roles)
}

func (m MsUser) GetPermissions() []string {
	return extractUserPermissions(m.Roles)
}

func (u MsUser) MarshalJSON() ([]byte, error) {
	type Alias MsUser
	return json.Marshal(&struct {
		Alias
		Roles       []MsRole `json:"roles"`
		RolesSlug   []string `json:"roles_slugs"`
		Permissions []string `json:"permissions"`
	}{
		Alias:       Alias(u),
		Roles:       u.Roles,
		RolesSlug:   extractRoleSlugs(u.Roles),
		Permissions: extractUserPermissions(u.Roles),
	})
}

// Extract role Name
func extractRoleNames(roles []MsRole) []string {
	slugs := make([]string, len(roles))
	for i, role := range roles {
		slugs[i] = role.Name
	}
	return slugs
}

// Extract role slugs
func extractRoleSlugs(roles []MsRole) []string {
	slugs := make([]string, len(roles))
	for i, role := range roles {
		slugs[i] = role.Slug
	}
	return slugs
}

// Extract unique permission names from all roles
func extractUserPermissions(roles []MsRole) []string {
	permissionSet := make(map[string]struct{})
	for _, role := range roles {
		for _, perm := range role.Permissions {
			permissionSet[perm.Name] = struct{}{}
		}
	}

	permissions := make([]string, 0, len(permissionSet))
	for perm := range permissionSet {
		permissions = append(permissions, perm)
	}
	return permissions
}
