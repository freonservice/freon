package cache

import (
	"github.com/MarcSky/freon/internal/auth/cache/memory"
)

type Storage interface {
	Get(key string) *memory.Item
	Set(key string, item memory.Item)
}
