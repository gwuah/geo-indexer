package main

import (
	"fmt"

	Controllers "github.com/electra-systems/athena/controllers"
	Utils "github.com/electra-systems/athena/utils"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
)

type UserLocationData struct {
	DriverId string `json:"id"`
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

func print(data ...interface{}) {
	fmt.Println(data)
}

func Server() {
	r := gin.Default()

	r.Use(Utils.CORSMiddleware())

	r.GET("/", Controllers.ServeBasicView)

	r.POST("/index-location", Controllers.IndexLocation)

	r.Run()
}

func main() {
	Server()
}
