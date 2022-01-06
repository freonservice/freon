package dal

import (
	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/pkg/freonApi"

	"github.com/dgraph-io/badger/v3"
	"google.golang.org/protobuf/proto"
)

const (
	settingTranslationKey = "setting.translation"
)

func (s *SettingRepo) updateSettingTranslateState() error {
	return s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(settingTranslationKey))
		switch err {
		default:
			return err
		case badger.ErrKeyNotFound:
			return nil
		case nil:
		}

		var value []byte
		_ = item.Value(func(val []byte) error {
			value = val
			return nil
		})

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

func (s SettingRepo) GetCurrentSettingState() domain.SettingConfiguration {
	return s.state
}
