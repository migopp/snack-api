package api

type MockStore struct {
	snackers    map[uint32]*Snacker
	numSnackers uint32
}

func CreateMockStore() *MockStore {
	return &MockStore{
		map[uint32]*Snacker{
			0: {
				ID:        0,
				FirstName: "Michael",
				LastName:  "Goppert",
				Snack:     "じゃがりこ",
				Hearts:    6,
			},
			1: {
				ID:        1,
				FirstName: "Dylan",
				LastName:  "Horton",
				Snack:     "Goldfish",
				Hearts:    5,
			},
			2: {
				ID:        2,
				FirstName: "Jason",
				LastName:  "Chavez",
				Snack:     "Dust",
				Hearts:    0,
			},
		}, 3,
	}
}

func (m *MockStore) CreateSnacker(s *SnackerRegistration) (uint32, error) {
	id := m.numSnackers
	snacker := &Snacker{
		ID:        id,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Snack:     s.Snack,
		Hearts:    0,
	}
	m.snackers[id] = snacker
	m.numSnackers += 1
	return id, nil
}

func (m *MockStore) DeleteSnacker(id uint32) error {
	delete(m.snackers, id)
	return nil
}

func (m *MockStore) UpdateSnacker(s *Snacker) error {
	m.snackers[s.ID] = s
	return nil
}

func (m *MockStore) FindSnacker(id uint32) (*Snacker, error) {
	snacker, ok := m.snackers[id]
	if !ok {
		return &Snacker{}, nil
	}

	return snacker, nil
}

func (m *MockStore) FetchSnackers() ([]*Snacker, error) {
	var snackers []*Snacker
	for _, snacker := range m.snackers {
		snackers = append(snackers, snacker)
	}
	return snackers, nil
}
