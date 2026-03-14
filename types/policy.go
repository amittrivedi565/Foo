package types

type Policy struct {
	Who string `json:"name"`
	Contracts []string `json:"contracts"`
}