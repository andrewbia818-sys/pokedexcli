package pokecache

import (
	"testing"
	"time"
)

// Test NewCache, checking for a valid return value.
func TestNewCache(t *testing.T) {
	cache := NewCache(1 * time.Second)
	if cache == nil {
		t.Errorf("NewCache() = nil, want non-nil")
	}
}

// Test Add, checking for a valid return value.
func TestAdd(t *testing.T) {
	cache := NewCache(1 * time.Second)
	cache.Add("key", []byte("value"))
	if _, exists := cache.Get("key"); !exists {
		t.Errorf("Add() = entry not found, want entry")
	}
}

// Test Get, checking for a valid return value.
func TestGet(t *testing.T) {
	cache := NewCache(1 * time.Second)
	cache.Add("key", []byte("value"))
	if _, exists := cache.Get("key"); !exists {
		t.Errorf("Get() = entry not found, want entry")
	}
}

// Test reapLoop, checking for a valid return value.
func TestReapLoop(t *testing.T) {
	cache := NewCache(1 * time.Second)
	cache.Add("key", []byte("value"))
	time.Sleep(2 * time.Second)
	if _, exists := cache.Get("key"); exists {
		t.Errorf("ReapLoop() = entry found, want entry not found")
	}
}
