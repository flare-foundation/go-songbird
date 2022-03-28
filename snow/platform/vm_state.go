package platform

type VMState interface {
	GetCurrentHeight() (uint64, error)
}
