package validators

type Quorum struct {
	Validators []FBA `json:"validators"`
}

type FBA struct {
	NodeID string `json:"nodeID"`
	Weight uint64 `json:"weight"`
}
