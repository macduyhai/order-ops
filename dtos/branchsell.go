package dtos

type BranchSell struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type AddbranchRequest struct {
	BranchSells []BranchSell `json:"branchsells"`
}

type AddbranchResponse struct {
	ID             int64    `json:"id,omitempty"`
	RecordsSuccess []string `json:"recordsSuccess,omitempty"`
	RecordsFailes  []string `json:"recordsFailes,omitempty"`
}
type DeletebranchRequest struct {
	Name string `json:"name"`
}
