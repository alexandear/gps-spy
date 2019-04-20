package util

import (
	"math"

	"github.com/devchallenge/spy-api/internal/model"
)

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(coordinate1 model.Coordinate, coordinate2 model.Coordinate) float64 {
	lat1 := float64(coordinate1.Latitude)
	lon1 := float64(coordinate1.Longitude)
	lat2 := float64(coordinate2.Latitude)
	lon2 := float64(coordinate2.Longitude)
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
