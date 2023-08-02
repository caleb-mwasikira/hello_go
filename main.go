package main

import (
	"fmt"
	"log"

	"example.com/hello_go/lib"
	nasa "example.com/hello_go/nasa"
)

func main() {
	var (
		distance float64
		err      error
	)

	filename := "data/mars_landing_sites.csv"

	data, err := lib.ReadWholeCSVFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	landingSites, err := nasa.CreateLandingSites(data[1:]) // Skipping the csv header line
	if err != nil {
		log.Fatal(err)
	}

	// What are the closest and farthest distances between landing sites
	mars := nasa.World{
		Radius: 3389.5,
	}
	landingSiteDistances := nasa.CalcLandingSiteDistances(mars, landingSites)
	nasa.CalcClosestLandingSites(landingSiteDistances)
	nasa.CalcFarthestLandingSites(landingSiteDistances)

	// Find the distance from London, England (51°30'N, 0°08'W) to Paris, France
	// (48°51'N, 2°21'E).
	earth := nasa.World{
		Radius: 6371.0,
	}

	london := `(51°30'N, 0°08'W)`
	paris := `(48°51'N, 2°21'E)`
	distance, err = earth.DistanceBetweenTwoLocations(london, paris)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Distance from London to Paris is %v\n", distance)

	// Find the distance between Mount Sharp (5°4'48"S, 137°51'E) and Olympus Mons
	// (18°39'N, 226°12'E) on Mars.
	mountSharp := `(5°4'48"S, 137°51'E)`
	olympusMons := `(18°39'N, 226°12'E)`
	distance, err = mars.DistanceBetweenTwoLocations(mountSharp, olympusMons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Distance between Mount Sharp and Olympus Mons on Mars is %v\n", distance)
	return
}
