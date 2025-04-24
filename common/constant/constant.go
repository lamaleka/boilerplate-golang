package constant

const (
	DefaultPageNumber    = 1
	DefaultPageSize      = 10
	DefaultMaxUploadSize = 2000000 // 5MB
)

const (
	WhereID        = "id = ?"
	WhereUUID      = "uuid = ?"
	WhereUsername  = "username = ?"
	WhereIsActive  = "is_active = ?"
	WhereBadge     = "badge = ?"
	WhereUUIDIn    = "uuid in ?"
	WhereIDIn      = "id in ?"
	DeletedNotNull = "deleted_at != null"
)
