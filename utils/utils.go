package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

type DriverData struct {
	id  string
	lng float64
	lat float64
}

func RunTests() {
	fmt.Println("Testing GEO lib and redis")
}

func TestRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
}

func CheckIfDriverExists() {
	driver_data := DriverData{lat: 5.683750933739772, lng: -0.24672031402587888}
	fmt.Println(driver_data)
}
