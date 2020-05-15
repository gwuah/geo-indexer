package main

import (
	"fmt"
	"strconv"

	Utils "github.com/electra-systems/athena/utils"

	"github.com/go-redis/redis"
	"github.com/uber/h3-go"

	"github.com/gin-gonic/gin"
)

type DriverLocationData struct {
	DriverId string `json:"driver_id"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}

var driverClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       1,
})

var carClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func print(data ...interface{}) {
	fmt.Println(data)
}

func Server() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Homepage",
		})
	})

	r.POST("/index-driver-location", func(c *gin.Context) {
		var data DriverLocationData

		if c.BindJSON(&data) != nil {

			c.JSON(500, gin.H{
				"message": "Error",
			})
		}

		lat, _ := strconv.ParseFloat(data.Lng, 64)
		lng, _ := strconv.ParseFloat(data.Lat, 64)

		h3Index := Utils.IndexLatLng(h3.GeoCoord{Latitude: lat, Longitude: lng})

		lastDriverLocation, err := driverClient.Get(data.DriverId).Result()

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error",
				"error":   err,
			})
			return
		}

		if Utils.H3IndexToString(h3Index) == lastDriverLocation {
			c.JSON(200, gin.H{
				"message": "Success",
			})
			return
		}

		_, err = driverClient.Set(data.DriverId, uint64(h3Index), 0).Result()

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error",
				"error":   err,
			})
			return
		}

		// find old index, remove driver
		// find new index, add driver

		c.JSON(200, gin.H{
			"message": "Index was successful",
			"data": gin.H{
				"driver_id":           data.DriverId,
				"last_driver_index":   lastDriverLocation,
				"latest_driver_index": h3Index,
			},
		})

	})

	r.Run()
}

func main() {
	Server()
}
