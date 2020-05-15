package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uber/h3-go"
)

type Coord struct {
	Lat, Lng float64
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

func GetCenter(index h3.H3Index) h3.GeoCoord {
	return h3.ToGeo(index)
}

func H3ToPolyline(h3idx h3.H3Index) []Coord {
	hexBoundary := h3.ToGeoBoundary(h3idx)
	hexBoundary = append(hexBoundary, hexBoundary[0])

	arr := []Coord{}

	for _, value := range hexBoundary {
		arr = append(arr, Coord{Lat: value.Latitude, Lng: value.Longitude})
	}

	return arr
}

func GeneratePolygons(rings []h3.H3Index) [][]Coord {
	arr := [][]Coord{}

	for _, value := range rings {
		arr = append(arr, H3ToPolyline(value))
	}

	return arr
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
