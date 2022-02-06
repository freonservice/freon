package frontend

import "net/http"

type errCode struct {
	status  int
	message string
}

var codeLabels []int

func newErrCode(statusCode int, errorText string) errCode {
	codeLabels = append(codeLabels, statusCode)
	return errCode{status: statusCode, message: errorText}
}

var (
	codeInternal = newErrCode(http.StatusInternalServerError, "internal error")
	codeNotFound = newErrCode(http.StatusNotFound, "Not found")
	// codeTooManyRequests     = newErrCode(http.StatusTooManyRequests, "Too many requests")
	codeEmailUsed = newErrCode(http.StatusConflict, "This email was used")
	// codeUserNotExist        = newErrCode(http.StatusNotFound, "User with this email not exist")
	codeWrongPassword = newErrCode(http.StatusForbidden, "User wrong password")
	// codeUserIsBanned        = newErrCode(http.StatusForbidden, "User is banned")
	codeLocalizationIsExist = newErrCode(http.StatusConflict, "Localization is exist")
	codeIdentifierIsExist   = newErrCode(http.StatusConflict, "Identifier is exist")
	codeCategoryIsExist     = newErrCode(http.StatusConflict, "Category is exist")
	codePasswordsNotEquals  = newErrCode(http.StatusConflict, "Passwords not equals")

	codeAutoTranslationNotSupported = newErrCode(http.StatusBadRequest, "auto translation not supported")
)
