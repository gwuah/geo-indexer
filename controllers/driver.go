// package controller

// import (
// 	"fmt"

// 	"github.com/go-redis/redis/v7"
// )

// type DriverData struct {
// 	id  string
// 	lng float64
// 	lat float64
// }

// func TestRedis() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)
// }

// func CheckIfDriverExists() {
// 	driver_data := DriverData{lat: 5.683750933739772, lng: -0.24672031402587888}
// 	fmt.Println(driver_data)
// }
