package lru

import "testing"

func Test_LRUCache_Empty(t *testing.T) {
	cache := NewLRUCache(0)
	cache.Put("foo", "bar")
	expect(t, cache.Get("foo"), "")
}

func Test_LRUCache_Expire(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("1", "1")

	expect(t, cache.Get("1"), "1")

	cache.Put("2", "2")
	expect(t, cache.Get("2"), "2")

	// expire 1
	cache.Put("3", "3")
	expect(t, cache.Get("3"), "3")

	expect(t, cache.Get("1"), "")
	// access 2, 3 will expire next time
	expect(t, cache.Get("2"), "2")

	// 3 expired
	cache.Put("1", "1")
	expect(t, cache.Get("3"), "")

	// change 2，1 will expire next time
	cache.Put("2", "foo")
	cache.Put("3", "3")
	expect(t, cache.Get("1"), "")

}

func expect(t *testing.T, got, exp string) {
	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
