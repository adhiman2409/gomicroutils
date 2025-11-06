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
