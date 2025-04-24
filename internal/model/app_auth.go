package model

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SSO      bool   `json:"sso"`
}

type LoginResponse struct {
	Me           *MsUser `json:"me"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

type SessionData struct {
	IsAuthenticated bool     `gorm:"type:boolean;not null;default:false" json:"is_authenticated"`
	PreferredName   *string  `gorm:"type:varchar(20)" json:"preferred_name"`
	Occupation      string   `gorm:"type:varchar(20)" json:"occupation"`
	Username        string   `gorm:"type:varchar(20)" json:"username"`
	SSO             int      `gorm:"type:boolean;not null;default:0" json:"sso"`
	Roles           []string `gorm:"many2many:piv_user_role" json:"roles"`
	Permissions     []string `gorm:"many2many:piv_user_role" json:"permissions"`
	AccessToken     string   `json:"access_token"`
	RefreshToken    string   `json:"refresh_token"`
	CreatedAt       *Time    `gorm:"not null;default:now()" json:"created_at"`
}

type CheckAuthSSORequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	Host         string `json:"host"`
}
type CheckAuthResponse struct {
	Authorization bool   `json:"authorization"`
	Message       string `json:"message"`
	Mfa           bool   `json:"mfa"`
	NoHp          bool   `json:"no_hp"`
}
