package dal

import (
	"encoding/json"
	"errors"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/pkg/freonApi"

	"github.com/dgraph-io/badger/v3"
	"google.golang.org/protobuf/proto"
)

const (
	settingStorageKey         = "settingStorage"
	settingTranslationKey     = "settingTranslation"
	settingFirstTimeLaunchKey = "settingFirstTimeLaunch"
)

type FirstLaunch struct {
	State bool `json:"state"`
}

func (s *SettingRepo) get(key string, txn *badger.Txn) ([]byte, error) {
	item, err := txn.Get([]byte(key))
	switch err {
	default:
		return nil, err
	case badger.ErrKeyNotFound:
		return []byte{}, nil
	case nil:
	}
	return item.ValueCopy(nil)
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
			Auto:         data.Auto,
			Use:          int32(data.Use),
			MainLanguage: data.MainLanguage,
		}
		return nil
	})
}

func (s *SettingRepo) SetTranslationConfiguration(ctx Ctx, data domain.TranslationConfiguration) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		val := freonApi.TranslationConfiguration{
			Auto:         data.Auto,
			Use:          freonApi.TranslationSource(data.Use),
			MainLanguage: data.MainLanguage,
		}
		value, err := proto.Marshal(&val)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(settingTranslationKey), value)
		if err != nil {
			return err
		}
		s.state.Translation = data
		return nil
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
		if data.Use == freonApi.StorageType_STORAGE_TYPE_S3 {
			if data.GetS3Configuration() == nil {
				return errors.New("badger serialization S3Configuration is nil")
			}
			s.state.Storage.S3Configuration = new(domain.S3Configuration)
			s.state.Storage.S3Configuration.SecretAccessKey = data.GetS3Configuration().GetSecretAccessKey()
			s.state.Storage.S3Configuration.AccessKeyID = data.GetS3Configuration().GetAccessKeyID()
			s.state.Storage.S3Configuration.Region = data.GetS3Configuration().GetRegion()
			s.state.Storage.S3Configuration.Endpoint = data.GetS3Configuration().GetEndpoint()
			s.state.Storage.S3Configuration.DisableSSL = data.GetS3Configuration().GetDisableSSL()
			s.state.Storage.S3Configuration.ForcePathStyle = data.GetS3Configuration().GetForcePathStyle()
		}
		return nil
	})
}

func (s *SettingRepo) SetStorageConfiguration(ctx Ctx, data domain.StorageConfiguration) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		val := freonApi.StorageConfiguration{
			Use:             freonApi.StorageType(data.Use),
			S3Configuration: nil,
		}
		if data.S3Configuration != nil {
			val.S3Configuration = &freonApi.S3Configuration{
				SecretAccessKey: data.S3Configuration.SecretAccessKey,
				AccessKeyID:     data.S3Configuration.AccessKeyID,
				Region:          data.S3Configuration.Region,
				Endpoint:        data.S3Configuration.Endpoint,
				DisableSSL:      data.S3Configuration.DisableSSL,
				ForcePathStyle:  data.S3Configuration.ForcePathStyle,
			}
		}
		value, err := proto.Marshal(&val)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(settingStorageKey), value)
		if err != nil {
			return err
		}
		s.state.Storage = data
		return nil
	})
}

func (s *SettingRepo) getSettingFirstLaunchState() error {
	return s.DB.View(func(txn *badger.Txn) error {
		value, err := s.get(settingFirstTimeLaunchKey, txn)
		if err != nil {
			return err
		}
		if len(value) == 0 {
			s.state.FirstLaunch = true
			return nil
		}
		var data FirstLaunch
		err = json.Unmarshal(value, &data)
		if err != nil {
			return err
		}
		s.state.FirstLaunch = data.State
		return nil
	})
}

func (s *SettingRepo) DisableSettingFirstLaunch(ctx Ctx) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		val := FirstLaunch{State: false}
		value, err := json.Marshal(&val)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(settingFirstTimeLaunchKey), value)
		if err != nil {
			return err
		}
		s.state.FirstLaunch = false
		return nil
	})
}

func (s SettingRepo) GetCurrentSettingState() domain.SettingConfiguration {
	return s.state
}
