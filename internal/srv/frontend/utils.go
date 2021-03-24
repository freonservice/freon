package frontend

import "github.com/MarcSky/freon/pkg/api"

func getUserStatusByInteger(status api.UserStatus) string {
	switch status {
	case api.UserStatus_USER_NOT_ACTIVE:
		return "not active"
	case api.UserStatus_USER_IS_BANNED:
		return "banned"
	default:
		return "active"
	}
}

func getUserStatusByString(status string) api.UserStatus {
	switch status {
	case "not active":
		return api.UserStatus_USER_NOT_ACTIVE
	case "banned":
		return api.UserStatus_USER_IS_BANNED
	default:
		return api.UserStatus_USER_ACTIVE
	}
}

func getUserRoleByInteger(role api.UserRole) string {
	switch role {
	case api.UserRole_USER_ROLE_ADMIN:
		return "admin"
	case api.UserRole_USER_ROLE_MODERATOR:
		return "moderator"
	default:
		return "translator"
	}
}

func getUserRoleByString(role string) api.UserRole {
	switch role {
	case "admin":
		return api.UserRole_USER_ROLE_ADMIN
	case "moderator":
		return api.UserRole_USER_ROLE_MODERATOR
	default:
		return api.UserRole_USER_ROLE_TRANSLATOR
	}
}
