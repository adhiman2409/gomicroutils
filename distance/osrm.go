package distance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Coordinate represents a geographic coordinate with latitude and longitude
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// OSRMRoute represents the response from OSRM routing API
type OSRMRoute struct {
	Code   string `json:"code"`
	Routes []struct {
		Geometry string  `json:"geometry"`
		Distance float64 `json:"distance"` // in meters
		Duration float64 `json:"duration"` // in seconds
	} `json:"routes"`
}

// RouteResult contains the decoded polyline points and route metadata
type RouteResult struct {
	Points   []Coordinate
	Distance float64 // in meters
	Duration float64 // in seconds
}

// RouteCache is a simple in-memory cache for routes
type RouteCache struct {
	mu    sync.RWMutex
	cache map[string]*RouteResult
}

var (
	routeCache = &RouteCache{
		cache: make(map[string]*RouteResult),
	}
)

// GetRoute fetches a route between coordinates using OSRM
func GetRoute(coordinates []Coordinate) (*RouteResult, error) {
	if len(coordinates) < 2 {
		return nil, fmt.Errorf("at least 2 coordinates are required")
	}

	// Create cache key
	cacheKey := createCacheKey(coordinates)

	// Check cache
	if result := routeCache.get(cacheKey); result != nil {
		return result, nil
	}

	// Build coordinate string for OSRM API
	coordParts := make([]string, len(coordinates))
	for i, c := range coordinates {
		coordParts[i] = fmt.Sprintf("%.6f,%.6f", c.Lng, c.Lat)
	}
	coordString := strings.Join(coordParts, ";")

	// Build URL
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s?overview=full&geometries=polyline", coordString)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("OSRM request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var osrmResp OSRMRoute
	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Validate response
	if osrmResp.Code != "Ok" || len(osrmResp.Routes) == 0 {
		return nil, fmt.Errorf("no route found")
	}

	// Decode polyline
	encodedPolyline := osrmResp.Routes[0].Geometry
	decodedPoints := decodePolyline(encodedPolyline)

	result := &RouteResult{
		Points:   decodedPoints,
		Distance: osrmResp.Routes[0].Distance,
		Duration: osrmResp.Routes[0].Duration,
	}

	// Cache the result
	routeCache.set(cacheKey, result)

	return result, nil
}

// GetDistanceBetweenPoints calculates the distance between two coordinates
func GetDistanceBetweenPoints(from, to Coordinate) (float64, error) {
	result, err := GetRoute([]Coordinate{from, to})
	if err != nil {
		return 0, err
	}
	return result.Distance, nil
}

// createCacheKey creates a unique key for caching
func createCacheKey(coordinates []Coordinate) string {
	parts := make([]string, len(coordinates))
	for i, c := range coordinates {
		parts[i] = fmt.Sprintf("%.6f,%.6f", c.Lat, c.Lng)
	}
	return strings.Join(parts, "|")
}

// Cache methods
func (rc *RouteCache) get(key string) *RouteResult {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.cache[key]
}

func (rc *RouteCache) set(key string, result *RouteResult) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.cache[key] = result

	// Clear old cache entries if size exceeds 1000
	if len(rc.cache) > 1000 {
		// Delete oldest 500 entries (simple approach - delete first 500 keys)
		count := 0
		for k := range rc.cache {
			if count >= 500 {
				break
			}
			delete(rc.cache, k)
			count++
		}
	}
}

// decodePolyline decodes a polyline string into an array of coordinates
// Based on the Polyline encoding algorithm
func decodePolyline(encoded string) []Coordinate {
	var coordinates []Coordinate
	index := 0
	lat := 0
	lng := 0

	for index < len(encoded) {
		var result int
		var shift uint
		var b int

		// Decode latitude
		for {
			b = int(encoded[index]) - 63
			index++
			result |= (b & 0x1f) << shift
			shift += 5
			if b < 0x20 {
				break
			}
		}

		var deltaLat int
		if result&1 != 0 {
			deltaLat = ^(result >> 1)
		} else {
			deltaLat = result >> 1
		}
		lat += deltaLat

		// Decode longitude
		result = 0
		shift = 0
		for {
			b = int(encoded[index]) - 63
			index++
			result |= (b & 0x1f) << shift
			shift += 5
			if b < 0x20 {
				break
			}
		}

		var deltaLng int
		if result&1 != 0 {
			deltaLng = ^(result >> 1)
		} else {
			deltaLng = result >> 1
		}
		lng += deltaLng

		coordinates = append(coordinates, Coordinate{
			Lat: float64(lat) / 1e5,
			Lng: float64(lng) / 1e5,
		})
	}

	return coordinates
}
