package api

type Snacker struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Snack     string `json:"snack"`
	Hearts    int    `json:"hearts"`
}

type SnackerRegistration struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Snack     string `json:"snack"`
}
