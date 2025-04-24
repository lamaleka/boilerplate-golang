package entity

type MsPermission struct {
	ID        uint       `json:"-"`
	UUID      UniqueID   `gorm:"type:uniqueidentifier;default:NEWSEQUENTIALID()" json:"id"`
	Name      string     `json:"name"`
	CreatedAt *Time      `json:"created_at"`
	UpdatedAt *Time      `json:"updated_at"`
	DeletedAt *DeletedAt `json:"-"`
}

func (MsPermission) TableName() string {
	return "ms_permission"
}
