package dtos

type Seller struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type AddsellerRequest struct {
	Sellers []Seller `json:"sellers"`
}

type AddsellerResponse struct {
	ID             int64    `json:"id,omitempty"`
	RecordsSuccess []string `json:"recordsSuccess,omitempty"`
	RecordsFailes  []string `json:"recordsFailes,omitempty"`
}
type DeletesellerRequest struct {
	Name string `json:"name"`
}
