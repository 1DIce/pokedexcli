package pokecache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/1DIce/pokedexcli/pokecache"
)

type mockEntry struct {
	v string
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val mockEntry
	}{
		{
			key: "https://example.com",
			val: mockEntry{v: "testdata"},
		},
		{
			key: "https://example.com/path",
			val: mockEntry{v: "moretestdata"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache[mockEntry](interval)
			cache.Set(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val.v) != string(c.val.v) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache[mockEntry](baseTime)
	cache.Set("https://example.com", mockEntry{v: "testdata"})

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
