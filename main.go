package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	olc "github.com/google/open-location-code/go"
)

type OpenLocationCodeLevel int

const (
	// An area of 20° x 20°. The side length of this tile varies with its location on the globe,but can be up to approximately 2200km. Tile addresses will be 2 characters long.
	GLOBAL OpenLocationCodeLevel = 2
	// An area of 1° x 1°. The side length of this tile varies with its location on the globe, but can be up to approximately 110km. Tile addresses will be 4 characters long.
	REGION OpenLocationCodeLevel = 4
	// An area of 0.05° x 0.05°. The side length of this tile varies with its location on the globe, but can be up to approximately 5.5km. Tile addresses will be 6 characters long.
	DISTRICT OpenLocationCodeLevel = 6
	// An area of 0.0025° x 0.0025°. The side length of this tile varies with its location on the globe, but can be up to approximately 275m. Tile addresses will be 8 characters long.
	NEIGHBORHOOD OpenLocationCodeLevel = 8
	// An area of 0.000125° x 0.000125°. The side length of this tile varies with its location on the globe, but can be up to approximately 14m. Tile addresses will be 10 characters long.
	PINPOINT OpenLocationCodeLevel = 10
)

type Area struct {
	lat, lon float64
	// Google open location code a.k.a google plus code
	olc string
}

func NewArea(geolocation string) (area Area, err error) {
	coords := strings.Split(geolocation, ",")
	if len(coords) != 2 {
		err = errors.New("Invalid geolocation")
		return
	}

	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return
	}

	lon, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return
	}

	area = Area{
		lat: lat,
		lon: lon,
		olc: olc.Encode(lat, lon, 0),
	}
	return
}

// Get latitude
func (a Area) GetLatitude() float64 {
	return a.lat
}

// Get longitude
func (a Area) GetLongitude() float64 {
	return a.lon
}

// Get open location code
func (a Area) GetOLC() string {
	return a.olc
}

func (a Area) IsWithinSameArea(target Area, level OpenLocationCodeLevel) bool {
	originCode := a.GetOLC()[:level]
	targetCode := target.GetOLC()[:level]
	fmt.Println(originCode, targetCode)
	return originCode == targetCode
}

func main() {
	// Raw
	origin := "-7.747693404128645,111.97074484841004"
	dest := "-7.7474508,111.9686441"

	originArea, err := NewArea(origin)
	if err != nil {
		panic(err)
	}

	destArea, err := NewArea(dest)
	if err != nil {
		panic(err)
	}

	fmt.Println(originArea.GetOLC())
	fmt.Println(destArea.GetOLC())
	fmt.Println(originArea.IsWithinSameArea(destArea, DISTRICT))
}
