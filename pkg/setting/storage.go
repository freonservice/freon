package setting

import (
	"github.com/dgraph-io/badger/v3"
)

type Storage struct {
	DB *badger.DB
}

func NewSetting(path string) (*Storage, error) {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}
