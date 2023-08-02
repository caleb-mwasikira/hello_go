package nasa

import (
	"fmt"
	"math"
	"os"
	"strings"
	"text/tabwriter"
)

/*
Write a program that determines the distance between each pair of
landing sites in table 22.1

Print out each of the locations in decimal degrees.
Use the distance method to write a program that determines the distance
 between each pair of landing sites in table 22.1.
Which two landing sites are the closest?
Which two are farthest apart?
*/

type LandingSite struct {
	Rover      string
	Site       string
	Coordinate Coordinate // a struct with latitude and longitude
}

func CreateLandingSites(data [][]string) ([]LandingSite, error) {
	var (
		landingSites []LandingSite
	)

	for _, record := range data {
		var landingSite LandingSite

		latitude, err := parseCoordinate(record[2])
		if err != nil {
			return nil, err
		}

		longitude, err := parseCoordinate(record[3])
		if err != nil {
			return nil, err
		}

		landingSite.Rover = record[0]
		landingSite.Site = record[1]
		landingSite.Coordinate.Latitude = latitude
		landingSite.Coordinate.Longitude = longitude

		landingSites = append(landingSites, landingSite)
	}

	return landingSites, nil
}

func CalcLandingSiteDistances(world World, landingSites []LandingSite) map[string]float64 {
	distances := make(map[string]float64)

	for _, siteOne := range landingSites {
		for _, siteTwo := range landingSites {
			if siteOne.Site == siteTwo.Site {
				// Refrain from calculating distances between a Location and itself
				continue
			}

			Locations := fmt.Sprintf("%v-%v", siteOne.Site, siteTwo.Site)
			distance := world.distanceBetweenTwoCoords(siteOne.Coordinate, siteTwo.Coordinate)
			distances[Locations] = distance
		}
	}
	return distances
}

func CalcClosestLandingSites(distances map[string]float64) {
	closestSites := ""
	smallestDistance := math.MaxFloat64

	for Coordinate, distance := range distances {
		if distance < smallestDistance {
			smallestDistance = distance
			closestSites = Coordinate
		}
	}

	sites := strings.Split(closestSites, "-")
	fmt.Printf("Closest landing sites are '%v' and '%v' with a distance of %v\n", sites[0], sites[1], smallestDistance)
}

func CalcFarthestLandingSites(distances map[string]float64) {
	farthestSites := ""
	largestDistance := -math.MaxFloat64

	for Coordinate, distance := range distances {
		if distance > largestDistance {
			largestDistance = distance
			farthestSites = Coordinate
		}
	}

	sites := strings.Split(farthestSites, "-")
	fmt.Printf("Farthest landing sites are '%v' and '%v' with a distance of %v\n", sites[0], sites[1], largestDistance)
}

func printLandingSites(landingSites []LandingSite) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Rover\tSite\tLatitude\tLongitude")
	for _, landingSite := range landingSites {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", landingSite.Rover, landingSite.Site, landingSite.Coordinate.Latitude, landingSite.Coordinate.Longitude)
	}
	w.Flush()
}
