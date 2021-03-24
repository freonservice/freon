package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MarcSky/freon/internal/app"
	"github.com/MarcSky/freon/internal/auth/cache"
	"github.com/MarcSky/freon/internal/auth/cache/memory"
	"github.com/MarcSky/freon/internal/config"
	"github.com/MarcSky/freon/pkg/api"

	"github.com/dgrijalva/jwt-go"
	openapierrors "github.com/go-openapi/errors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

const (
	audienceToken = "freon"
	issuerToken   = "auth-freon"
	sessionExpire = 10 * time.Second
)

var (
	ErrTokenInvalid = openapierrors.New(http.StatusUnauthorized, "Token is invalid")
)

type (
	auth struct {
		storage   cache.Storage
		repo      app.Repo
		secretKey []byte
		logger    *structlog.Logger
	}

	session struct {
		UserID int64
		Status int64
	}
)

func NewAuth(secret string, repo app.Repo, logger *structlog.Logger) app.Auth {
	return &auth{
		secretKey: []byte(secret),
		repo:      repo,
		storage:   memory.NewStorage(),
		logger:    logger,
	}
}

func (a *auth) IsAuthorized(accessToken string) (*app.UserSession, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt: unexpected signing method %s %s", "alt", token.Header["alg"])
		}
		return a.secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, ErrTokenInvalid
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, ErrTokenInvalid
	}

	if claims.Audience != audienceToken || claims.Issuer != issuerToken {
		a.logger.PrintErr("jwt: invalid claims", "aud", claims.Audience, "iss", claims.Issuer)
		return nil, ErrTokenInvalid
	}

	uID, err := uuid.Parse(claims.Id)
	if err != nil {
		a.logger.PrintErr(errors.WithStack(err))
		return nil, ErrTokenInvalid
	}

	sess, err := a.session(uID.String(), accessToken)
	if err != nil {
		a.logger.PrintErr(errors.WithStack(err))
		return nil, ErrTokenInvalid
	}

	switch sess.Status {
	case int64(api.UserStatus_USER_IS_BANNED):
		return nil, app.ErrUserIsBanned
	case int64(api.UserStatus_USER_NOT_ACTIVE):
		return nil, app.ErrUserNotActive
	default:
	}

	return &app.UserSession{
		UserID:   sess.UserID,
		UserUUID: uID.String(),
	}, nil
}

func (a *auth) GenerateAuthToken(uID uuid.UUID) (string, error) {
	claims := jwt.StandardClaims{
		Id:        uID.String(),
		Audience:  audienceToken,
		ExpiresAt: time.Now().Add(config.JwtTokenLifetime).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuerToken,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString(a.secretKey)
}

func (a *auth) session(uuid, token string) (*session, error) {
	item := a.storage.Get(uuid)
	if item != nil {
		return &session{
			UserID: item.UserID,
			Status: item.Status,
		}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sess, err := a.repo.SessionByAccessToken(ctx, token)
	if err != nil {
		a.logger.Println(errors.WithStack(err))
		return nil, ErrTokenInvalid
	}

	a.storage.Set(uuid, memory.Item{
		UserID:     sess.UserID,
		Status:     sess.User.Status,
		Expiration: time.Now().Add(sessionExpire).Unix(),
	})

	return &session{
		UserID: sess.UserID,
		Status: sess.User.Status,
	}, nil
}
