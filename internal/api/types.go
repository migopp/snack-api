package api

type Snacker struct {
	ID        uint64   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Snacks    []string `json:"snacks"`
	Hearts    uint32   `json:"hearts"`
}

type SnackerRegistration struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Snacks    []string `json:"snacks"`
}
