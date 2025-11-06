package distance

import (
	"fmt"
	"testing"
)

func TestGetDistanceBetweenPoints(t *testing.T) {
	// Example: Distance between two points in Delhi
	from := Coordinate{Lat: 28.6139, Lng: 77.2090} // New Delhi
	to := Coordinate{Lat: 28.7041, Lng: 77.1025}   // Delhi Airport

	distance, err := GetDistanceBetweenPoints(from, to)
	if err != nil {
		t.Fatalf("Failed to get distance: %v", err)
	}

	fmt.Printf("Distance: %.2f meters (%.2f km)\n", distance, distance/1000)

	if distance <= 0 {
		t.Error("Expected positive distance")
	}
}

func TestGetRoute(t *testing.T) {
	// Multiple waypoints
	coordinates := []Coordinate{
		{Lat: 28.6139, Lng: 77.2090},
		{Lat: 28.6500, Lng: 77.2167},
		{Lat: 28.7041, Lng: 77.1025},
	}

	result, err := GetRoute(coordinates)
	if err != nil {
		t.Fatalf("Failed to get route: %v", err)
	}

	fmt.Printf("Distance: %.2f meters\n", result.Distance)
	fmt.Printf("Duration: %.2f seconds\n", result.Duration)
	fmt.Printf("Points in route: %d\n", len(result.Points))

	if result.Distance <= 0 {
		t.Error("Expected positive distance")
	}

	if len(result.Points) == 0 {
		t.Error("Expected decoded polyline points")
	}
}

func TestCaching(t *testing.T) {
	from := Coordinate{Lat: 28.6139, Lng: 77.2090}
	to := Coordinate{Lat: 28.7041, Lng: 77.1025}

	// First call - should hit the API
	distance1, err := GetDistanceBetweenPoints(from, to)
	if err != nil {
		t.Fatalf("Failed to get distance: %v", err)
	}

	// Second call - should use cache
	distance2, err := GetDistanceBetweenPoints(from, to)
	if err != nil {
		t.Fatalf("Failed to get distance: %v", err)
	}

	if distance1 != distance2 {
		t.Error("Cached result should match original result")
	}
}

func TestInvalidInput(t *testing.T) {
	// Test with less than 2 coordinates
	_, err := GetRoute([]Coordinate{{Lat: 28.6139, Lng: 77.2090}})
	if err == nil {
		t.Error("Expected error for single coordinate")
	}
}

func TestLRUCacheEviction(t *testing.T) {
	// Create a small cache for testing
	testCache := NewRouteCache(3)

	// Add 3 entries
	testCache.set("key1", &RouteResult{Distance: 100})
	testCache.set("key2", &RouteResult{Distance: 200})
	testCache.set("key3", &RouteResult{Distance: 300})

	if testCache.Size() != 3 {
		t.Errorf("Expected cache size 3, got %d", testCache.Size())
	}

	// Add 4th entry - should evict key1 (least recently used)
	testCache.set("key4", &RouteResult{Distance: 400})

	if testCache.Size() != 3 {
		t.Errorf("Expected cache size 3 after eviction, got %d", testCache.Size())
	}

	// key1 should be evicted
	if result := testCache.get("key1"); result != nil {
		t.Error("key1 should have been evicted")
	}

	// Other keys should still exist
	if result := testCache.get("key2"); result == nil || result.Distance != 200 {
		t.Error("key2 should still exist in cache")
	}
	if result := testCache.get("key3"); result == nil || result.Distance != 300 {
		t.Error("key3 should still exist in cache")
	}
	if result := testCache.get("key4"); result == nil || result.Distance != 400 {
		t.Error("key4 should exist in cache")
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	testCache := NewRouteCache(3)

	// Add 3 entries
	testCache.set("key1", &RouteResult{Distance: 100})
	testCache.set("key2", &RouteResult{Distance: 200})
	testCache.set("key3", &RouteResult{Distance: 300})

	// Access key1 (moves it to front)
	testCache.get("key1")

	// Add key4 - should evict key2 (now the least recently used)
	testCache.set("key4", &RouteResult{Distance: 400})

	// key2 should be evicted
	if result := testCache.get("key2"); result != nil {
		t.Error("key2 should have been evicted")
	}

	// key1 should still exist (was recently accessed)
	if result := testCache.get("key1"); result == nil || result.Distance != 100 {
		t.Error("key1 should still exist in cache")
	}
}

func TestLRUCacheClear(t *testing.T) {
	testCache := NewRouteCache(10)

	// Add some entries
	testCache.set("key1", &RouteResult{Distance: 100})
	testCache.set("key2", &RouteResult{Distance: 200})

	if testCache.Size() != 2 {
		t.Errorf("Expected cache size 2, got %d", testCache.Size())
	}

	// Clear cache
	testCache.Clear()

	if testCache.Size() != 0 {
		t.Errorf("Expected cache size 0 after clear, got %d", testCache.Size())
	}

	// Entries should be gone
	if result := testCache.get("key1"); result != nil {
		t.Error("Cache should be empty after clear")
	}
}

func TestLRUCacheThreadSafety(t *testing.T) {
	testCache := NewRouteCache(100)

	// Run concurrent operations
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				testCache.set(key, &RouteResult{Distance: float64(j * 100)})
				testCache.get(key)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Cache should not exceed maxSize
	if testCache.Size() > 100 {
		t.Errorf("Cache size %d exceeds maximum 100", testCache.Size())
	}
}
