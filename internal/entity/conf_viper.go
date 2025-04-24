package entity

type ConfViper struct {
	App *ConfApp
	Web *ConfWeb
	Log *ConfLog
	Jwt *ConfJwtConfig
	Db  *ConfDbConfig
	Api *ConfApiConfig
}
type ConfApp struct {
	AppName string
}
type ConfWeb struct {
	Prefork bool
	Port    int
}
type ConfLog struct {
	Level int
}
type ConfJwtConfig struct {
	Access  *ConfJwt
	Refresh *ConfJwt
}
type ConfJwt struct {
	Secret  string
	Expired int
	MaxAge  int
}
type ConfDbConfig struct {
	App *ConfDb
	Bi  *ConfDb
}
type ConfDb struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Pool     *ConfDbPool
}
type ConfDbPool struct {
	Idle     int
	Max      int
	Lifetime int
}
type ConfApiConfig struct {
	Sso    *ConfApiSso
	Webdav *ConfApiWebdav
}
type ConfApiSso struct {
	Url    string
	Secret string
}

type ConfApiWebdav struct {
	Url    string
	User   string
	Secret string
	Path   string
}
