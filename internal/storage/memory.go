package storage

type Store interface {
	Save(code string, url string)
	Get(code string) (string, bool)
}

type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Save(code string, url string) {
	m.data[code] = url
}

func (m *MemoryStore) Get(code string) (string, bool) {
	val, ok := m.data[code]
	return val, ok
}