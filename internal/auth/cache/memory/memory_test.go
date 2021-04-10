package memory_test

import (
	"testing"
	"time"

	"github.com/freonservice/freon/internal/auth/cache/memory"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Memory_GetEmpty(t *testing.T) {
	storage := memory.NewStorage()
	item := storage.Get("MY_KEY")
	assert.Nil(t, item)
}

func TestStorage_Memory_GetValue(t *testing.T) {
	storage := memory.NewStorage()
	var (
		userID int64 = 1
		status int64 = 2
	)
	storage.Set("MY_KEY", memory.Item{
		UserID:     userID,
		Status:     status,
		Expiration: time.Now().UTC().Unix() + 10000,
	})
	item := storage.Get("MY_KEY")
	assert.NotNil(t, item)
	assert.EqualValues(t, userID, item.UserID)
	assert.EqualValues(t, status, item.Status)
}

func TestStorage_Memory_GetExpiredValue(t *testing.T) {
	storage := memory.NewStorage()
	storage.Set("MY_KEY", memory.Item{Expiration: 1})
	item := storage.Get("MY_KEY")
	assert.Nil(t, item)
}
