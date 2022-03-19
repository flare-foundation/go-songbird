package mocks

type VMState struct {
	GetCurrentHeightF func() (uint64, error)
}

func (v *VMState) GetCurrentHeight() (uint64, error) {
	return v.GetCurrentHeightF()
}
