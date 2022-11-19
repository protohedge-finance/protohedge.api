package enums

type Error uint64

const (
	NotFound Error = iota
	ServerError
)

func (s Error) String() string {
	switch s {
	case NotFound:
		return "notFound"
	case ServerError:
		return "serverError"
	}
	return "unknown"
}
