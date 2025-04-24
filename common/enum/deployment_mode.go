package enum

type DeploymentMode int

const (
	Prod DeploymentMode = iota + 1
	Dev
	Local
)

func (s DeploymentMode) Label() string {
	switch s {
	case Prod:
		return "Prod"
	case Dev:
		return "Dev"
	case Local:
		return "Local"
	default:
		return "Unknown"
	}
}

func (s DeploymentMode) File() string {
	switch s {
	case Prod:
		return "config.json"
	case Dev:
		return "config.dev.json"
	case Local:
		return "config.local.json"
	default:
		return "Unknown"
	}
}
