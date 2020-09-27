package cache

import "testing"

// TestSet check simple setting elements to cache.
func TestSet(t *testing.T) {
	c, err := NewCache(3)
	if err != nil {
		t.Errorf("unexpected error in NewCache: %v", err)
	}

	updated := c.Set("key1", 2)
	if updated {
		t.Errorf("expected updated: %v, got: %v", false, updated)
	}

	updated = c.Set("key2", "hello")
	if updated {
		t.Errorf("expected updated: %v, got: %v", false, updated)
	}

	v, ok := c.Get("key1")
	if !ok {
		t.Errorf("expected ok: %v, got: %v", true, ok)
	}

	vInt, ok := v.(int)
	if !ok {
		t.Errorf("can not assert value type as int for value: %v", v)
	}

	if vInt != 2 {
		t.Errorf("expected vInt: %v, got: %v", 2, vInt)
	}

	v, ok = c.Get("key2")
	if !ok {
		t.Errorf("expected ok: %v, got: %v", true, ok)
	}

	vStr, ok := v.(string)
	if !ok {
		t.Errorf("can not assert value type as string for value: %v", v)
	}

	if vStr != "hello" {
		t.Errorf("expected vStr: %v, got: %v", "hello", vStr)
	}
}

// TestSetPop checks that last element will be
// deleted from the cache if cache size is exceeded.
func TestSetPop(t *testing.T) {
	c, err := NewCache(2)
	if err != nil {
		t.Errorf("unexpected error in NewCache: %v", err)
	}

	c.Set("k1", 4)
	c.Set("k2", 5)
	c.Set("k3", 9)

	v, ok := c.Get("k1")
	if ok {
		t.Errorf("expected ok: %v, got: %v", false, ok)
	}

	if v != nil {
		t.Errorf("expected v: %v, got: %v", nil, v)
	}

	v, ok = c.Get("k2")
	if !ok {
		t.Errorf("expected ok: %v, got: %v", true, ok)
	}

	if v != 5 {
		t.Errorf("expected v: %v, got: %v", 5, v)
	}

	v, ok = c.Get("k3")
	if !ok {
		t.Errorf("expected ok: %v, got: %v", true, ok)
	}

	if v != 9 {
		t.Errorf("expected v: %v, got: %v", 9, v)
	}
}

// TestSetPopLeastUsed checks that the least used
// item will be deleted from the queue to set
// new item if cache size is exceeded.
func TestSetPopLeastUsed(t *testing.T) {
	c, err := NewCache(3)
	if err != nil {
		t.Errorf("unexpected error in NewCache: %v", err)
	}

	c.Set("k1", 4)
	c.Set("k2", 5)
	c.Set("k3", 9)

	_, ok := c.Get("k3")
	if !ok {
		t.Errorf("expected ok: %v, got: %v", true, ok)
	}

	updated := c.Set("k2", 13)
	if !updated {
		t.Errorf("expected updated: %v, got: %v", true, updated)
	}

	updated = c.Set("k3", 2)
	if !updated {
		t.Errorf("expected updated: %v, got: %v", true, updated)
	}

	updated = c.Set("k4", 90)
	if updated {
		t.Errorf("expected updated: %v, got: %v", false, updated)
	}

	v, ok := c.Get("k1")
	if ok {
		t.Errorf("expected ok: %v, got: %v", false, ok)
	}

	if v != nil {
		t.Errorf("expected v: %v, got: %v", nil, v)
	}
}
