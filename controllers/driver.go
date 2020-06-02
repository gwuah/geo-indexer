package controllers

import (
	"strconv"

	Utils "github.com/electra-systems/athena/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/uber/h3-go"
)

type UserLocationData struct {
	Id       string `json:"id"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}

var userClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       1,
})

var carClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       2,
})

func IndexLocation(c *gin.Context) {
	var data UserLocationData

	if c.BindJSON(&data) != nil {

		c.JSON(500, gin.H{
			"message": "Error",
		})

		return
	}

	lat, _ := strconv.ParseFloat(data.Lat, 64)
	lng, _ := strconv.ParseFloat(data.Lng, 64)

	h3Index := Utils.IndexLatLng(h3.GeoCoord{Latitude: lat, Longitude: lng})
	stringifiedIndex := Utils.H3IndexToString(h3Index)

	lastLocationIndex, err := userClient.Get(data.Id).Result()

	if err != redis.Nil && err != nil {

		c.JSON(500, gin.H{
			"message": "Last user location lookup error",
			"error":   err,
		})
		return
	}

	if stringifiedIndex == lastLocationIndex {
		c.JSON(200, gin.H{
			"message": "User hasn't changed position",
		})
		return
	}

	_, err = userClient.Set(data.Id, uint64(h3Index), 0).Result()

	if err != nil {

		c.JSON(500, gin.H{
			"message": "Updating user location failed ",
			"error":   err,
		})
		return
	}

	_, err = carClient.LRem(lastLocationIndex, 0, data.Id).Result()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Updating old index failed",
			"error":   err,
		})
		return
	}

	_, err = carClient.LPush(stringifiedIndex, data.Id).Result()

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
			"id":           data.Id,
			"last_index":   lastLocationIndex,
			"latest_index": stringifiedIndex,
			"lat":                 lat,
			"lng":                 lng,
		},
	})

}

func ServeBasicView(c *gin.Context) {
	h3Index := Utils.IndexLatLng(h3.GeoCoord{Latitude: 5.678981813723179, Longitude: -0.24087421107286566})

	rings := h3.KRing(h3Index, 1)

	c.JSON(200, gin.H{
		"message": "Homepage",
		"rings":   Utils.GeneratePolygons(rings),
	})
}
