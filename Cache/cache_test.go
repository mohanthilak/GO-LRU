package cache

import (
	"testing"
	"time"
)

func TestIntantiateCache(t *testing.T) {
	cache := InstantiateCache(10, 10*time.Second)
	t.Cleanup(func() {
		cache = nil
	})
	if cache.Capacity != 10 {
		t.Errorf("Expected capacity to be 10, got %d", cache.Capacity)
	}
	if cache.EvictionDuration != 10*time.Second {
		t.Errorf("Expected eviction duration to be 10s, got %v", cache.EvictionDuration)
	}
}

func TestSetGet(t *testing.T) {
	t.Run("Testing set and Get", func(t *testing.T) {
		cache := InstantiateCache(2, 0) // Disable eviction for this test
		t.Cleanup(func() {
			cache = nil
		})
		if cache.Capacity != 2 {
			t.Errorf("Expected capacity to be 2, got %d", cache.Capacity)
		}
		cache.Set("key1", "value1")
		if value := cache.Get("key1"); value != "value1" {
			t.Errorf("Expected value1 for key1, got %v", value)
		}

		cache.Set("key2", "value2")
		if value := cache.Get("key2"); value != "value2" {
			t.Errorf("Expected value2 for key2 after adding key2, got %v", value)
		}
	})

	t.Run("Testing stability after change in order", func(t *testing.T) {
		cache := InstantiateCache(2, 0) // Disable eviction for this test

		t.Cleanup(func() {
			cache = nil
		})

		cache.Set("key1", "value1")
		if value := cache.Get("key1"); value != "value1" {
			t.Fatalf("Expected value1 for key1 after adding key2, got %v", value)
		}
		cache.Set("key2", "value2")
		if value := cache.Get("key1"); value != "value1" {
			t.Fatalf("Expected value1 for key1 after adding key2, got %v", value)
		}
		if value := cache.Get("key2"); value != "value2" {
			t.Errorf("Expected value2 for key2, got %v", value)
		}
	})

	t.Run("Testing stability after updating value", func(t *testing.T) {
		cache := InstantiateCache(2, 0) // Disable eviction for this test

		t.Cleanup(func() {
			cache = nil
		})

		cache.Set("key1", "updatedValue1")
		if value := cache.Get("key1"); value != "updatedValue1" {
			t.Errorf("Expected updatedValue1 for key1 after update, got %v", value)
		}
	})

	t.Run("Testing Time based Eviction", func(t *testing.T) {
		cache := InstantiateCache(2, 2*time.Second)
		if cache.EvictionDuration != 2*time.Second {
			t.Errorf("Expected eviction duration to be 2s, got %v", cache.EvictionDuration)
		}
		cache.Set("key1", "value1")
		cache.Set("key2", "value2")

		time.Sleep(3 * time.Second)
		cache.Set("key3", "value3")
		_ = cache.Get("key1")

		if value := cache.Get("key1"); value != -1 {
			t.Errorf("Expected key1 to be evicted after timeout, got %v", value)
		}
		if value := cache.Get("key2"); value != -1 {
			t.Errorf("Expected key2 to be evicted after timeout, got %v", value)
		}
		if value := cache.Get("key3"); value != "value3" {
			t.Errorf("Expected value3 for key3, got %v", value)
		}

	})
}
