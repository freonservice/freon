package dal

import (
	"strings"
)

// func isNotFoundError(err error) bool {
//	return err != nil && err == sql.ErrNoRows
// }

func isDuplicateKeyValue(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
