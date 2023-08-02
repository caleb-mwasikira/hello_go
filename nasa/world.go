package nasa

import (
	"math"
)

type World struct {
	Radius float64
}

type Coordinate struct {
	Latitude, Longitude float64
}

func degreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

/* Distance calculation using the Spherical Law of Cosines. */
func (w World) distanceBetweenTwoCoords(coord1, coord2 Coordinate) float64 {
	s1, c1 := math.Sincos(degreesToRadians(coord1.Latitude))
	s2, c2 := math.Sincos(degreesToRadians(coord2.Latitude))

	clong := math.Cos(degreesToRadians(coord1.Longitude - coord2.Longitude))
	return w.Radius * math.Acos(s1*s2+c1*c2*clong)
}

/* Calculates the distance between two locations */
func (w World) DistanceBetweenTwoLocations(fromLocation, toLocation string) (float64, error) {
	fromCoord, err := parseCoordinatePair(fromLocation)
	if err != nil {
		return 0, err
	}

	toCoord, err := parseCoordinatePair(toLocation)
	if err != nil {
		return 0, err
	}

	distance := w.distanceBetweenTwoCoords(*fromCoord, *toCoord)
	return distance, nil
}
