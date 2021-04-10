package cache

import (
	"github.com/freonservice/freon/internal/auth/cache/memory"
)

type Storage interface {
	Get(key string) *memory.Item
	Set(key string, item memory.Item)
}
