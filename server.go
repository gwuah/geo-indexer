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
	DB:       2,
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

			return
		}

		lat, _ := strconv.ParseFloat(data.Lng, 64)
		lng, _ := strconv.ParseFloat(data.Lat, 64)

		h3Index := Utils.IndexLatLng(h3.GeoCoord{Latitude: lat, Longitude: lng})
		stringifiedIndex := Utils.H3IndexToString(h3Index)

		lastDriverLocationIndex, err := driverClient.Get(data.DriverId).Result()

		if err != redis.Nil && err != nil {

			c.JSON(500, gin.H{
				"message": "Last driver location lookup error",
				"error":   err,
			})
			return
		}

		if stringifiedIndex == lastDriverLocationIndex {
			c.JSON(200, gin.H{
				"message": "Driver hasn't changed position",
			})
			return
		}

		_, err = driverClient.Set(data.DriverId, uint64(h3Index), 0).Result()

		if err != nil {

			c.JSON(500, gin.H{
				"message": "Updating driver location failed ",
				"error":   err,
			})
			return
		}

		_, err = carClient.LRem(lastDriverLocationIndex, 0, data.DriverId).Result()

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Updating old index failed",
				"error":   err,
			})
			return
		}

		_, err = carClient.LPush(stringifiedIndex, data.DriverId).Result()

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Updating new index failed",
				"error":   err,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Index was successful",
			"data": gin.H{
				"driver_id":           data.DriverId,
				"last_driver_index":   lastDriverLocationIndex,
				"latest_driver_index": stringifiedIndex,
			},
		})

	})

	r.Run()
}

func main() {
	// Server()
	print("" || nil)
}
