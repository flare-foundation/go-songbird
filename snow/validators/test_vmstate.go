package validators

type VMStateMock struct {
	GetCurrentHeightF func() (uint64, error)
}

func (v *VMStateMock) GetCurrentHeight() (uint64, error) {
	return v.GetCurrentHeightF()
}
