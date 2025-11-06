package distance

import (
	"container/list"
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

// cacheEntry holds a cache value and its key for LRU tracking
type cacheEntry struct {
	key    string
	result *RouteResult
}

// RouteCache is an LRU (Least Recently Used) in-memory cache for routes
type RouteCache struct {
	mu       sync.RWMutex
	cache    map[string]*list.Element
	lruList  *list.List
	maxSize  int
}

// NewRouteCache creates a new LRU cache with the specified maximum size
func NewRouteCache(maxSize int) *RouteCache {
	if maxSize <= 0 {
		maxSize = 1000 // default size
	}
	return &RouteCache{
		cache:   make(map[string]*list.Element),
		lruList: list.New(),
		maxSize: maxSize,
	}
}

var (
	// Default cache with 1000 entries limit
	routeCache = NewRouteCache(1000)
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

// Cache methods with LRU eviction
func (rc *RouteCache) get(key string) *RouteResult {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if elem, exists := rc.cache[key]; exists {
		// Move to front (most recently used)
		rc.lruList.MoveToFront(elem)
		return elem.Value.(*cacheEntry).result
	}
	return nil
}

func (rc *RouteCache) set(key string, result *RouteResult) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	// If key already exists, update and move to front
	if elem, exists := rc.cache[key]; exists {
		rc.lruList.MoveToFront(elem)
		elem.Value.(*cacheEntry).result = result
		return
	}

	// Add new entry
	entry := &cacheEntry{
		key:    key,
		result: result,
	}
	elem := rc.lruList.PushFront(entry)
	rc.cache[key] = elem

	// Evict least recently used if size exceeds max
	if rc.lruList.Len() > rc.maxSize {
		rc.evictOldest()
	}
}

// evictOldest removes the least recently used entry from cache
func (rc *RouteCache) evictOldest() {
	elem := rc.lruList.Back()
	if elem != nil {
		rc.lruList.Remove(elem)
		entry := elem.Value.(*cacheEntry)
		delete(rc.cache, entry.key)
	}
}

// Size returns the current number of entries in the cache
func (rc *RouteCache) Size() int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.lruList.Len()
}

// Clear removes all entries from the cache
func (rc *RouteCache) Clear() {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.cache = make(map[string]*list.Element)
	rc.lruList = list.New()
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
