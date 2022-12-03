package enums

type Network uint64

const (
	Arbitrum Network = iota
)

func (s Network) String() string {
	switch s {
	case Arbitrum:
		return "Arbitrum"
	}
	return "unknown"
}
