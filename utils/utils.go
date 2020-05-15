package utils

import (
	"fmt"

	"github.com/uber/h3-go"
)

type DriverData struct {
	id  string
	lng float64
	lat float64
}

func IndexLatLng(coordinates h3.GeoCoord) h3.H3Index {
	return h3.FromGeo(coordinates, 8)
}

func H3IndexToString(index h3.H3Index) string {
	return fmt.Sprintf("%v", index)
}

func FormatH3Index(index h3.H3Index) string {
	return fmt.Sprintf("%#x\n", index)
}
