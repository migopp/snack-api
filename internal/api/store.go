package api

type Store interface {
	CreateSnacker(*SnackerRegistration) (uint64, error)
	DeleteSnacker(uint64) error
	UpdateSnacker(*Snacker) error
	FindSnacker(uint64) (*Snacker, error)
	FetchSnackers() ([]*Snacker, error)
}
