package nasa

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
Coordinates can be represented as degrees° minutes' seconds"(DMS) format
or in decimal form
*/

type DMSCoordinate struct {
	Degrees   float64
	Minutes   float64
	Seconds   float64
	Direction string
}

/* Converts a DMS coordinate to its decimal equivalent */
func (coord *DMSCoordinate) toDecimal() float64 {
	// Calculate the decimal degrees based on the DMS values
	decimalDegrees := float64(coord.Degrees) + (float64(coord.Minutes) / 60) + (float64(coord.Seconds) / 3600)

	// Reverse the sign for S and W directions
	if coord.Direction == "S" || coord.Direction == "W" {
		decimalDegrees = -decimalDegrees
	}

	return decimalDegrees
}

/*
Takes in a coordinate string and checks if it is a valid
degrees,minutes,seconds(DMS) coordinate

The regex used:

		^: 						Start of the string
	    [+-]?:					An optional plus or minus sign for positive or negative coordinates.
	    (\d{1,3}): 				Capturing group for 1 to 3 digits representing degrees.
	    °?: 					An optional degree symbol.
	    \s?: 					An optional space character to separate degrees from minutes.
	    (\d{1,2}): 				Capturing group for 1 to 2 digits representing minutes.
	    '?: 					An optional single quote character (apostrophe) to separate minutes from seconds.
	    (?: ... )?: 			A non-capturing group to make the seconds and the second's symbol optional.
	    \s?:					An optional space character to separate minutes from seconds (if seconds are present).
	    (\d{1,2}(?:\.\d+)?)?: 	Capturing group for 1 to 2 digits with an optional decimal point and additional digits, representing seconds.
								The (?: ... ) is a non-capturing group to include the decimal point and additional digits, if present. The ? at the end makes the seconds and its decimal part optional.
	    "?: 					An optional double quote character to represent seconds (if seconds are present).
	    \s?: 					An optional space character before the direction (if seconds are present).
	    [NSWE]?: 				An optional character representing the direction (N, S, E, or W).
	    $: 						End of the string.
*/
func isDMSCoord(coord string) bool {
	pattern := `^[+-]?(\d{1,3})°?\s?(\d{1,2})'?(?:\s?(\d{1,2}(?:\.\d+)?)"?)?\s?[NSWE]?$`
	regex := regexp.MustCompile(pattern)
	return regex.Match([]byte(coord))
}

/*
Checks if a direction string is a valid compass direction (N, S, E, W)
*/
func isCompassDirection(direction string) bool {
	compassDirections := []string{"N", "S", "E", "W"}
	compassDirectionsStr := strings.Join(compassDirections, "")
	return strings.ContainsAny(strings.ToUpper(direction), compassDirectionsStr)
}

/*
Takes in a coord string and converts it to its corresponding data type;
If the coord string is in the degrees° minutes' seconds"(DMS) format
then it is converted into a CoordinatesDMS struct

If the coord string is in the decimal format, then it is converted
into its decimal form

For example:

	parseCoordinate("135°54'0\" E")
	Returns: CoordinatesDMS{Degrees=135 Minutes=54 Seconds=0 Direction=E}

While:

	parseCoordinate("135.90000")
	Returns: 135.90000
*/
func parseCoordinate(coord string) (float64, error) {
	if !isDMSCoord(coord) {
		decimalCoord, err := strconv.ParseFloat(coord, 64)
		if err != nil {
			return 0, fmt.Errorf("Invalid coordinate format: %v", err)
		}

		return decimalCoord, nil
	}

	// Removes all DMS characters(°'") and replaces them with white space
	replacer := strings.NewReplacer("°", " ", "'", " ", "\"", " ")
	coord = replacer.Replace(strings.Trim(coord, " "))

	// Splits the DMS coordinate string into its distinctive parts
	parts := strings.Fields(coord)

	// Create a new coordinate
	DMSCoord := &DMSCoordinate{}

	for i, item := range parts {
		lastIndex := i == len(parts)-1

		if !lastIndex {
			switch i {
			case 0: // Degrees
				val, err := strconv.ParseFloat(item, 64)
				if err != nil {
					return 0, err
				}
				DMSCoord.Degrees = val
			case 1: // Minutes
				val, err := strconv.ParseFloat(item, 64)
				if err != nil {
					return 0, err
				}
				DMSCoord.Minutes = val
			default: // Seconds
				val, err := strconv.ParseFloat(item, 64)
				if err != nil {
					return 0, err
				}
				DMSCoord.Seconds = val
			}
		} else {
			if isCompassDirection(item) {
				DMSCoord.Direction = item
			}
		}
	}

	return DMSCoord.toDecimal(), nil
}

func parseCoordinatePair(coordPair string) (*Coordinate, error) {
	coordPair = strings.Trim(coordPair, "()")
	coordArr := strings.Split(coordPair, ",")

	latitude, err := parseCoordinate(strings.Trim(coordArr[0], " "))
	if err != nil {
		return nil, err
	}
	longitude, err := parseCoordinate(strings.Trim(coordArr[1], " "))
	if err != nil {
		return nil, err
	}

	return &Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}, err
}
