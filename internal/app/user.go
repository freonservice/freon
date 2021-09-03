package app

import (
	"github.com/freonservice/freon/internal/entities"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

func (a *appl) AuthorizeUser(ctx Ctx, email, password string) (AccessToken, *entities.User, error) {
	user, err := a.repo.GetUserByEmail(email)
	switch err {
	default:
		return "", nil, err
	case reform.ErrNoRows:
		return "", nil, ErrNotFound
	case nil:
	}

	switch user.Status {
	case int64(api.UserStatus_USER_NOT_ACTIVE):
		return "", nil, ErrUserNotActive
	case int64(api.UserStatus_USER_IS_BANNED):
		return "", nil, ErrUserIsBanned
	case int64(api.UserStatus_USER_ACTIVE):
	}

	if !a.pass.Compare([]byte(user.Password), []byte(password)) {
		return "", nil, ErrWrongPassword
	}

	userID, err := uuid.Parse(user.UUIDID)
	if err != nil {
		return "", nil, errors.Wrap(err, "app.AuthorizeUser uuid parsing error")
	}

	token, err := a.auth.GenerateAuthToken(userID)
	if err != nil {
		return "", nil, errors.Wrap(err, "app.AuthorizeUser generation auth token error")
	}

	err = a.repo.SaveSession(ctx, user.ID, token)
	if err != nil {
		return "", nil, errors.Wrap(err, "app.AuthorizeUser save session error")
	}

	return token, mappingUser(user), nil
}

func (a *appl) RegisterUser(ctx Ctx, email, password, firstName, secondName string, role int64) (*entities.User, error) {
	user, err := a.repo.GetUserByEmail(email)
	if err == nil && user != nil {
		return nil, ErrEmailIsUsed
	}

	passwordHash, err := a.pass.Hashing(password)
	if err != nil {
		return nil, errors.Wrap(err, "app.RegisterUser hashing password error")
	}

	newUser, err := a.repo.CreateUser(ctx, email, string(passwordHash), firstName, secondName, role)
	if err != nil {
		return nil, errors.Wrap(err, "app.RegisterUser create user error")
	}
	return mappingUser(newUser), nil
}

func (a *appl) GetUserByUUID(ctx Ctx, userUUID string) (*entities.User, error) {
	user, err := a.repo.GetUserByUserUUID(userUUID)
	if err != nil {
		return nil, err
	}
	return mappingUser(user), nil
}

func (a *appl) GetUserByID(ctx Ctx, userID int64) (*entities.User, error) {
	user, err := a.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return mappingUser(user), nil
}

func (a *appl) GetUserByEmail(ctx Ctx, email string) (*entities.User, error) {
	user, err := a.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return mappingUser(user), nil
}

func (a *appl) LogoutUser(ctx Ctx, token string) error {
	return a.repo.DeleteSession(ctx, token)
}

func (a *appl) UpdatePassword(ctx Ctx, userID int64, changePassword entities.ChangePassword) error {
	user, err := a.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if !a.pass.Compare([]byte(user.Password), []byte(changePassword.PreviousPassword)) {
		return ErrPasswordNotCorrect
	}

	passwordHash, err := a.pass.Hashing(changePassword.NewPassword)
	if err != nil {
		return err
	}

	return a.repo.UpdatePassword(ctx, user.ID, string(passwordHash))
}

func (a *appl) UpdateProfile(ctx Ctx, userID int64, email, firstName, secondName string, role, status int64) error {
	return a.repo.UpdateProfile(ctx, userID, email, firstName, secondName, role, status)
}

func (a *appl) GetUsers(ctx Ctx) ([]*entities.User, error) {
	u, err := a.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayUser(u), nil
}

func (a *appl) UpdateStatus(ctx Ctx, userID, status int64) error {
	return a.repo.UpdateStatus(ctx, userID, status)
}
