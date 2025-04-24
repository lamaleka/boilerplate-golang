package entity

type MsEmployee struct {
	ID           uint         `json:"-"`
	UUID         UniqueID     `gorm:"type:uniqueidentifier;default:NEWSEQUENTIALID()" json:"id"`
	UserID       *string      `json:"-"`
	Nama         string       `json:"nama"`
	Badge        string       `json:"badge"`
	DeptID       *string      `json:"dept_id"`
	DeptTitle    string       `json:"dept_title"`
	Email        *string      `json:"email"`
	PosID        *string      `json:"pos_id"`
	PosTitle     string       `json:"pos_title"`
	EmployeeType EmployeeType `json:"employee_type"`
	IsActive     bool         `json:"is_active"`
	CreatedAt    *Time        `json:"created_at"`
	UpdatedAt    *Time        `json:"updated_at"`
	DeletedAt    *DeletedAt   `json:"-"`
}

func (MsEmployee) TableName() string {
	return "ms_employee"
}
