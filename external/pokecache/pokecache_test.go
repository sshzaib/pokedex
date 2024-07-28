package pokecache

import "testing"

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

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
