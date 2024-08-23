package api

type MockStore struct {
	users []Snacker
}

func CreateMockStore() *MockStore {
	return &MockStore{
		[]Snacker{
			{
				ID:        0,
				FirstName: "Michael",
				LastName:  "Goppert",
				Snacks:    []string{"じゃがりこ"},
				Hearts:    6,
			},
			{
				ID:        1,
				FirstName: "Dylan",
				LastName:  "Horton",
				Snacks:    []string{"Goldfish", "Braum's Cheeseburger"},
				Hearts:    5,
			},
			{
				ID:        2,
				FirstName: "Jason",
				LastName:  "Chavez",
				Snacks:    []string{"Dust"},
				Hearts:    0,
			},
		},
	}
}

func (m *MockStore) CreateSnacker(s *SnackerRegistration) (uint64, error) {
	id := uint64(len(m.users))
	snacker := &Snacker{
		ID:        id,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Snacks:    s.Snacks,
		Hearts:    0,
	}
	m.users = append(m.users, *snacker)
	return id, nil
}

func (m *MockStore) DeleteSnacker(id uint64) error {
	m.users[id] = Snacker{}
	return nil
}

func (m *MockStore) UpdateSnacker(s *Snacker) error {
	m.users[s.ID] = *s
	return nil
}

func (m *MockStore) FindSnacker(id uint64) (*Snacker, error) {
	if id >= uint64(len(m.users)) {
		return &Snacker{}, nil
	}

	return &m.users[id], nil
}

func (m *MockStore) FetchSnackers() ([]*Snacker, error) {
	var snackers []*Snacker
	for _, snacker := range m.users {
		snackers = append(snackers, &snacker)
	}
	return snackers, nil
}
