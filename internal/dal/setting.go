package dal

import (
	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/pkg/freonApi"

	"github.com/dgraph-io/badger/v3"
	"google.golang.org/protobuf/proto"
)

const (
	settingStorageKey     = "setting.storage"
	settingTranslationKey = "setting.translation"
)

func (s *SettingRepo) get(key string, txn *badger.Txn) ([]byte, error) {
	item, err := txn.Get([]byte(key))
	switch err {
	default:
		return nil, err
	case badger.ErrKeyNotFound:
		return []byte{}, nil
	case nil:
	}

	var value []byte
	_ = item.Value(func(val []byte) error {
		value = val
		return nil
	})

	return value, nil
}

func (s *SettingRepo) getSettingTranslateState() error {
	return s.DB.View(func(txn *badger.Txn) error {
		value, err := s.get(settingTranslationKey, txn)
		if err != nil {
			return err
		}

		var data freonApi.TranslationConfiguration
		err = proto.Unmarshal(value, &data)
		if err != nil {
			return err
		}

		s.state.Translation = domain.TranslationConfiguration{
			Auto: data.Auto,
			Use:  int32(data.Use),
		}
		return nil
	})
}

func (s *SettingRepo) SetTranslationConfiguration(ctx Ctx, data domain.TranslationConfiguration) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		val := freonApi.TranslationConfiguration{
			Auto: data.Auto,
			Use:  freonApi.TranslationSource(data.Use),
		}
		value, err := proto.Marshal(&val)
		if err != nil {
			return err
		}
		return txn.Set(value, []byte(settingTranslationKey))
	})
}

func (s *SettingRepo) getSettingStorageState() error {
	return s.DB.View(func(txn *badger.Txn) error {
		value, err := s.get(settingStorageKey, txn)
		if err != nil {
			return err
		}

		var data freonApi.StorageConfiguration
		err = proto.Unmarshal(value, &data)
		if err != nil {
			return err
		}

		s.state.Storage = domain.StorageConfiguration{
			Use: int32(data.Use),
		}
		return nil
	})
}

func (s *SettingRepo) SetStorageConfiguration(ctx Ctx, data domain.StorageConfiguration) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		val := freonApi.StorageConfiguration{
			Use: freonApi.StorageType(data.Use),
		}
		value, err := proto.Marshal(&val)
		if err != nil {
			return err
		}
		return txn.Set(value, []byte(settingStorageKey))
	})
}

func (s SettingRepo) GetCurrentSettingState() domain.SettingConfiguration {
	return s.state
}
