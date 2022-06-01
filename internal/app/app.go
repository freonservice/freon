package app

import (
	"context"

	filter2 "github.com/freonservice/freon/internal/dal/filter"
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/storage"
	iface "github.com/freonservice/freon/internal/translation"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
	"golang.org/x/text/language"
)

type (
	Ctx         = context.Context
	AccessToken = string

	Appl interface {
		AuthorizeUser(ctx Ctx, email, password string) (AccessToken, *domain.User, error)
		RegisterUser(ctx Ctx, email, password, firstName, secondName string, role int64) (*domain.User, error)
		GetUserByUUID(ctx Ctx, uuid string) (*domain.User, error)
		GetUserByID(ctx Ctx, userID int64) (*domain.User, error)
		GetUserByEmail(ctx Ctx, email string) (*domain.User, error)
		LogoutUser(ctx Ctx, token string) error
		UpdateStatus(ctx Ctx, userID, status int64) error
		UpdatePassword(ctx Ctx, userID int64, changePassword domain.ChangePassword) error
		UpdateProfile(ctx Ctx, userID int64, email, firstName, secondName string, role, status int64) error
		GetUsers(ctx Ctx) ([]*domain.User, error)

		CreateLocalization(ctx Ctx, creatorID int64, name, code string) error
		GetLocalizations(ctx Ctx) ([]*domain.Localization, error)
		DeleteLocalization(ctx Ctx, id int64) error

		CreateIdentifier(
			ctx Ctx, creatorID, categoryID, parentID int64, name, description, textSingular, textPlural string, platforms []string,
		) error
		GetIdentifiers(ctx Ctx, f filter2.IdentifierFilter) ([]*domain.Identifier, error)
		DeleteIdentifier(ctx Ctx, id int64) error
		UpdateIdentifier(
			ctx Ctx, id, categoryID, parentID int64, name, description, textSingular, textPlural string, platforms []string,
		) error

		CreateCategory(ctx Ctx, name string) error
		GetCategories(ctx Ctx) ([]*domain.Category, error)
		DeleteCategory(ctx Ctx, id int64) error
		UpdateCategory(ctx Ctx, id int64, name string) error

		CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, singular, plural string) error
		GetTranslations(ctx Ctx, f filter2.TranslationFilter) ([]*domain.Translation, error)
		DeleteTranslation(ctx Ctx, id int64) error
		UpdateTranslation(ctx Ctx, id int64, singular, plural string) error
		UpdateStatusTranslation(ctx Ctx, id, status int64) error
		GetTranslation(ctx Ctx, locale, identifierName string) (*domain.Translation, error)
		GetGroupedTranslations(ctx Ctx, f filter2.GroupedTranslationFilter) ([]*domain.GroupedTranslations, error)

		CreateTranslationFile(ctx Ctx, platform, storageType string, creatorID, localizationID int64) error
		GetTranslationFiles(ctx Ctx, f filter2.TranslationFileFilter) ([]*domain.TranslationFile, error)
		DeleteTranslationFile(ctx Ctx, id int64) error

		GetStatistic(ctx Ctx) (*domain.Statistic, error)
		GetVersion(ctx Ctx, localizationID, platform int64) ([]*domain.Version, error)

		GetCurrentSettingState() domain.SettingConfiguration
		SetTranslationConfiguration(ctx Ctx, data domain.TranslationConfiguration) error
		SetStorageConfiguration(ctx Ctx, data domain.StorageConfiguration) error
		DisableSettingFirstLaunch(ctx Ctx) error

		GetSupportedLanguages(ctx Ctx) ([]*domain.Language, error)
		Translate(ctx Ctx, text string, source, target language.Tag) (string, error)
		CreateAutoTranslationByID(ctx Ctx, id int64) error

		HealthCheck(Ctx) (interface{}, error)
	}

	Auth interface {
		IsAuthorized(token string) (*UserSession, error)
		GenerateAuthToken(userUUID uuid.UUID) (string, error)
	}

	Repo interface {
		GetDB() *sqlx.DB

		CreateUser(ctx Ctx, email, password, firstName, secondName string, role int64) (*dao.User, error)
		UpdatePassword(ctx Ctx, userID int64, passwordHash string) error
		UpdateProfile(ctx Ctx, userID int64, email, firstName, secondName string, role, status int64) error
		GetUsers(ctx Ctx) ([]*dao.User, error)
		UpdateStatus(ctx Ctx, userID, status int64) error

		GetUserByUserUUID(userUUID string) (*dao.User, error)
		GetUserByEmail(email string) (*dao.User, error)
		GetUserByID(id int64) (*dao.User, error)

		SaveSession(ctx Ctx, userID int64, token AccessToken) error
		SessionByAccessToken(ctx Ctx, token AccessToken) (*dao.UserSession, error)
		DeleteSession(ctx Ctx, token AccessToken) error

		CreateLocalization(ctx Ctx, creatorID int64, locale, languageName string) (*dao.Localization, error)
		GetLocalization(ctx Ctx, id int64) (*dao.Localization, error)
		GetLocalizations(ctx Ctx) ([]*dao.Localization, error)
		DeleteLocalization(ctx Ctx, id int64) error

		CreateIdentifier(
			ctx Ctx, createID, categoryID, parentID int64, name, description, textSingular, textPlural, platforms string,
		) (int64, error)
		GetIdentifiers(ctx Ctx, f filter2.IdentifierFilter) ([]*dao.Identifier, error)
		DeleteIdentifier(ctx Ctx, id int64) error
		UpdateIdentifier(
			ctx Ctx, id, categoryID, parentID int64, name, description, textSingular, textPlural, platforms string) error
		UpdateStatusTranslation(ctx Ctx, id, status int64) error

		CreateCategory(ctx Ctx, name string) error
		GetCategories(ctx Ctx) ([]*dao.Category, error)
		DeleteCategory(ctx Ctx, id int64) error
		UpdateCategory(ctx Ctx, id int64, name string) error

		CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, singular, plural string) error
		GetTranslations(ctx Ctx, f filter2.TranslationFilter) ([]*dao.Translation, error)
		DeleteTranslation(ctx Ctx, id int64) error
		UpdateTranslation(ctx Ctx, id int64, singular, plural string) error
		UpdateTranslationWithMeta(ctx Ctx, localizationID, identifierID int64, singular, plural string) error
		GetTranslation(ctx Ctx, locale, identifierName string) (*dao.Translation, error)
		GetGroupedTranslations(ctx Ctx, f filter2.GroupedTranslationFilter) (map[string][]*dao.Translation, error)
		GetTranslationByID(ctx Ctx, id int64) (*dao.Translation, error)

		CreateTranslationFile(ctx Ctx, name, path, s3fileID, s3bucket string, platform, storageType, creatorID, localizationID int64) error
		GetTranslationFile(ctx Ctx, id int64) (*dao.TranslationFile, error)
		GetTranslationFiles(ctx Ctx, f filter2.TranslationFileFilter) ([]*dao.TranslationFile, error)
		DeleteTranslationFile(ctx Ctx, id int64) error

		GetStatistic(ctx Ctx) (*dao.Statistic, error)
		GetVersionFromTranslationFiles(ctx Ctx, f filter2.VersionTranslationFilesFilter) ([]*dao.Version, error)
		GetVersionFromTranslations(ctx Ctx, f filter2.VersionTranslationsFilter) ([]*dao.Version, error)

		GetLanguages(ctx Ctx) ([]*dao.Language, error)
	}

	SettingRepo interface {
		GetCurrentSettingState() domain.SettingConfiguration
		SetTranslationConfiguration(ctx Ctx, data domain.TranslationConfiguration) error
		SetStorageConfiguration(ctx Ctx, data domain.StorageConfiguration) error
		DisableSettingFirstLaunch(ctx Ctx) error
	}

	Password interface {
		Hashing(password string) ([]byte, error)
		Compare(hashedPassword []byte, password []byte) bool
		Generate(length int) string
	}

	Svc struct {
		repo        Repo
		auth        Auth
		pass        Password
		setting     SettingRepo
		translation iface.Translation
		storage     storage.Storage
	}

	Config struct {
		TranslationFilesFolder string
	}

	appl struct {
		svc    *Svc
		config Config
		logger *structlog.Logger
	}

	UserSession struct {
		UserID   int64
		UserUUID string
	}
)

func NewSvc(repo Repo, auth Auth, pass Password, setting SettingRepo, translation iface.Translation, s storage.Storage) *Svc {
	return &Svc{
		repo,
		auth,
		pass,
		setting,
		translation,
		s,
	}
}

func New(svc *Svc, config Config, logger *structlog.Logger) Appl {
	return &appl{
		svc:    svc,
		config: config,
		logger: logger,
	}
}

func (a *appl) HealthCheck(_ Ctx) (interface{}, error) {
	return "OK", nil
}
