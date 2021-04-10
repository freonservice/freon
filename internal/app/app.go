package app

import (
	"context"

	"github.com/freonservice/freon/internal/dao"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	Ctx         = context.Context
	AccessToken = string

	Appl interface {
		AuthorizeUser(ctx Ctx, email, password string) (AccessToken, *User, error)
		RegisterUser(ctx Ctx, email, password, firstName, secondName string, role int64) (*User, error)
		GetUserByUUID(ctx Ctx, uuid string) (*User, error)
		GetUserByID(ctx Ctx, userID int64) (*User, error)
		GetUserByEmail(ctx Ctx, email string) (*User, error)
		LogoutUser(ctx Ctx, token string) error
		UpdateStatus(ctx Ctx, userID, status int64) error
		UpdatePassword(ctx Ctx, userID int64, changePassword ChangePassword) error
		UpdateProfile(ctx Ctx, userID int64, email, firstName, secondName string, role, status int64) error
		GetUsers(ctx Ctx) ([]*User, error)

		CreateLocalization(ctx Ctx, creatorID int64, name, code, icon string) error
		GetLocalizations(ctx Ctx) ([]*Localization, error)
		DeleteLocalization(ctx Ctx, id int64) error

		CreateIdentifier(
			ctx Ctx, creatorID, categoryID, parentID int64, name, description, exampleText string,
			platforms, namedList []string,
		) error
		GetIdentifiers(ctx Ctx, categoryID int64) ([]*Identifier, error)
		DeleteIdentifier(ctx Ctx, id int64) error
		UpdateIdentifier(
			ctx Ctx, id, categoryID, parentID int64, name, description, exampleText string,
			platforms, namedList []string,
		) error

		CreateCategory(ctx Ctx, name string) error
		GetCategories(ctx Ctx) ([]*Category, error)
		DeleteCategory(ctx Ctx, id int64) error
		UpdateCategory(ctx Ctx, id int64, name string) error

		CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, text string) error
		GetTranslations(ctx Ctx, localizationID int64) ([]*Translation, error)
		DeleteTranslation(ctx Ctx, id int64) error
		UpdateTranslation(ctx Ctx, id int64, text string) error
		HideTranslation(ctx Ctx, id int64, hide bool) error

		GetStatistic(ctx Ctx) (*Statistic, error)

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

		CreateLocalization(ctx Ctx, creatorID int64, locale, languageName, icon string) (*dao.Localization, error)
		GetLocalizations(ctx Ctx) ([]*dao.Localization, error)
		DeleteLocalization(ctx Ctx, id int64) error

		CreateIdentifier(
			ctx Ctx, createID, categoryID, parentID int64,
			name, description, exampleText, platforms, namedList string,
		) error
		GetIdentifiers(ctx Ctx, filter dao.IdentifierFilter) ([]*dao.Identifier, error)
		DeleteIdentifier(ctx Ctx, id int64) error
		UpdateIdentifier(
			ctx Ctx, id, categoryID, parentID int64, name, description, exampleText, platforms, namedList string) error
		UpdateHideStatusTranslation(ctx Ctx, id int64, hide bool) error

		CreateCategory(ctx Ctx, name string) error
		GetCategories(ctx Ctx) ([]*dao.Category, error)
		DeleteCategory(ctx Ctx, id int64) error
		UpdateCategory(ctx Ctx, id int64, name string) error

		CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, text string) error
		GetTranslations(ctx Ctx, filter dao.TranslationFilter) ([]*dao.Translation, error)
		DeleteTranslation(ctx Ctx, id int64) error
		UpdateTranslation(ctx Ctx, id int64, text string) error

		GetStatistic(ctx Ctx) (*dao.Statistic, error)
	}

	Password interface {
		Hashing(password string) ([]byte, error)
		Compare(hashedPassword []byte, password []byte) bool
		Generate(length int) string
	}

	appl struct {
		repo Repo
		auth Auth
		pass Password
	}

	UserSession struct {
		UserID   int64
		UserUUID string
	}
)

func New(repo Repo, auth Auth, pass Password) Appl {
	return &appl{
		repo: repo,
		auth: auth,
		pass: pass,
	}
}

func (a *appl) HealthCheck(_ Ctx) (interface{}, error) {
	return "OK", nil
}
