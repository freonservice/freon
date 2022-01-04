package setting

import (
	"github.com/dgraph-io/badger/v3"
)

type Storage struct {
	DB *badger.DB
}

func NewSetting(path string) (*Storage, error) {
	db, err := badger.Open(badger.DefaultOptions(path).WithInMemory(false))
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}
