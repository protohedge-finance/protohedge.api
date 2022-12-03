package config

type Env uint64

const (
	Local Env = iota
	Production
)

func (s Env) String() string {
	switch s {
	case Local:
		return "local"
	case Production:
		return "production"
	}
	return "unknown"
}
