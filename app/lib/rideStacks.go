package lib

import (
	"fmt"
	"math"
	"time"
)

type RideStacks struct {
	CoordsA   string
	CoordsB   string
	StartTime time.Time
	EndTime   time.Time

	lat1 float64
	lon1 float64
	lat2 float64
	lon2 float64

	latDiff float64
	lonDiff float64

	Duration time.Duration
	Distance float64

	Price int
}

func (c RideStacks) euclideanDistance() float64 {
	c.latDiff = c.lat2 - c.lat1
	c.lonDiff = c.lon2 - c.lon1

	distance := math.Sqrt(c.latDiff*c.latDiff+c.lonDiff*c.lonDiff) * 111.32

	return distance
}

func (c RideStacks) calculatePrice() int {
	c.Distance = c.euclideanDistance()
	price := 3500 + c.Distance*1000 + c.Duration.Minutes()*200
	c.Price = int(price)
	return c.Price
}

func (c RideStacks) converterCoordinates(coords string) (float64, float64, error) {
	var lat, lon float64

	n, err := fmt.Sscanf(coords, "%f, %f", &lat, &lon)

	if n != 2 || err != nil {
		return 0, 0, fmt.Errorf("no se pudieron escanear las dos coordenadas")
	}

	return lat, lon, nil
}

func (c RideStacks) calculateDuration() time.Duration {
	c.Duration = c.EndTime.Sub(c.StartTime)
	return c.Duration
}

func CalculateRide(
	CoordsA string,
	CoordsB string,
	StartTime time.Time,
	EndTime time.Time,
) (*RideStacks, error) {

	c := &RideStacks{
		CoordsA:   CoordsA,
		CoordsB:   CoordsB,
		StartTime: StartTime,
		EndTime:   EndTime,
	}

	if c.StartTime.IsZero() || c.EndTime.IsZero() {
		return c, fmt.Errorf("no se pudieron obtener las fechas")
	}

	if c.CoordsA == "" || c.CoordsB == "" {
		return c, fmt.Errorf("no se pudieron obtener las coordenadas")
	}

	lat1, lon1, err := c.converterCoordinates(c.CoordsA)
	if err != nil {
		return c, err
	}

	c.lat1 = lat1
	c.lon1 = lon1

	lat2, lon2, err := c.converterCoordinates(c.CoordsB)

	if err != nil {
		return c, err
	}

	c.lat2 = lat2
	c.lon2 = lon2

	c.calculateDuration()

	c.calculatePrice()

	return c, nil
}
