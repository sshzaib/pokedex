package pokecache

import (
	"testing"
	"time"
)

func TestAddGetCache(t *testing.T) {
	interval := 3 * time.Second
	cache := NewCache(interval)

	cases := []struct {
		input    string
		data     []byte
		expected []byte
	}{
		{
			input:    "key1",
			data:     []byte("value1"),
			expected: []byte("value1"),
		},
	}
	for _, cs := range cases {
		cache.AddCache(cs.input, cs.data)
	}
	for _, tcase := range cases {
		val, ok := cache.GetCache(tcase.input)
		if !ok {
			t.Errorf("%v not found", tcase.input)
		}
		if string(val) != string(tcase.expected) {
			t.Errorf("Error: cache data not the same\n expected: %v \n actual: %v\n", tcase.expected, val)
		}
	}
}

func TestReapLoop(t *testing.T) {
	interval := 3 * time.Second
	cache := NewCache(interval)
	go cache.reapLoop(3 * time.Second)
	cache.AddCache("key1", []byte("val1"))

	if val, ok := cache.GetCache("key1"); !ok {
		t.Errorf("test1: failed\n returned no cache\n expected %v \ngot: %v", "val1", val)
	}
	time.Sleep(4 * time.Second)

	if val, ok := cache.GetCache("key1"); ok {
		t.Errorf("test2: failed\n returned the key\n expected: %v got: %v", "val1", val)
	}

}
