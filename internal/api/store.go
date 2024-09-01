package api

type Store interface {
	CreateSnacker(*SnackerRegistration) (uint32, error)
	DeleteSnacker(uint32) error
	UpdateSnacker(*Snacker) error
	FindSnacker(uint32) (*Snacker, error)
	FetchSnackers() ([]*Snacker, error)
}
