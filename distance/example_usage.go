package distance

import "fmt"

// Example usage of the distance package

func ExampleUsage() {
	// Example 1: Get distance between two points
	from := Coordinate{Lat: 28.6139, Lng: 77.2090}
	to := Coordinate{Lat: 28.7041, Lng: 77.1025}

	distance, err := GetDistanceBetweenPoints(from, to)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Distance: %.2f meters (%.2f km)\n", distance, distance/1000)

	// Example 2: Get full route with multiple waypoints
	waypoints := []Coordinate{
		{Lat: 28.6139, Lng: 77.2090},
		{Lat: 28.6500, Lng: 77.2167},
		{Lat: 28.7041, Lng: 77.1025},
	}

	result, err := GetRoute(waypoints)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Distance: %.2f meters\n", result.Distance)
	fmt.Printf("Total Duration: %.2f seconds (%.2f minutes)\n", result.Duration, result.Duration/60)
	fmt.Printf("Route has %d points\n", len(result.Points))

	// You can iterate through the decoded polyline points
	fmt.Println("\nFirst 5 points of the route:")
	for i := 0; i < 5 && i < len(result.Points); i++ {
		fmt.Printf("  Point %d: Lat=%.6f, Lng=%.6f\n", i+1, result.Points[i].Lat, result.Points[i].Lng)
	}
}
