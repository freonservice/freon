package memory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Memory_GetEmpty(t *testing.T) {
	storage := NewStorage()
	item := storage.Get("MY_KEY")
	if item != nil {
		assert.EqualValues(t, 0, item.UserID)
		assert.EqualValues(t, 0, item.Status)
	}
}

func TestStorage_Memory_GetValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("MY_KEY", 1, 2, 5*time.Second)
	item := storage.Get("MY_KEY")
	if item != nil {
		assert.EqualValues(t, 1, item.UserID)
		assert.EqualValues(t, 2, item.Status)
		return
	}
	t.Error("TestStorage_Memory_GetValue is failed")
}

func TestStorage_Memory_GetExpiredValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("MY_KEY", 1, 2, 10*time.Second)
	time.Sleep(1 * time.Second)
	item := storage.Get("MY_KEY")
	if item != nil {
		assert.EqualValues(t, 1, item.UserID)
		assert.EqualValues(t, 2, item.Status)
		return
	}
	t.Error("TestStorage_Memory_GetExpiredValue is failed")
}
