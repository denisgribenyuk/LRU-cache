package lru

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lru(t *testing.T) {
	req := require.New(t)
	cache := NewLRUCahce(4)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")
	cache.Add("key4", "value4")

	t.Run("Correct get", func(t *testing.T) {
		value, ok := cache.Get("key2")
		req.True(ok)
		req.Equal("value2", value)
	})

	t.Run("Correct remove", func(t *testing.T) {
		ok := cache.Remove("key2")
		req.True(ok)
		_, ok = cache.Get("key2")
		req.False(ok)
		cache.Add("key2", "value2")
	})

	t.Run("Correct update", func(t *testing.T) {
		ok := cache.Add("key3", "newvalue3")
		req.True(ok)
		v, ok := cache.Get("key3")
		req.True(ok)
		req.Equal("newvalue3", v)
		cache.Add("key2", "value2")
	})

	t.Run("Last element was deleted", func(t *testing.T) {
		_, ok := cache.Get("key1")
		req.True(ok)
		_, ok = cache.Get("key3")
		req.True(ok)
		_, ok = cache.Get("key4")
		req.True(ok)
		cache.Add("key5", "value5")
		_, ok = cache.Get("key2")
		req.False(ok)
	})

}
