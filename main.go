package main

import (
	"fmt"
	"log"
	"unsplash_analog/config"
	"unsplash_analog/http_api"
	"unsplash_analog/postgres"
	"unsplash_analog/redis"
)

func main() {
	log.Println("Server started")
	// load env config
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// connect to redis
	redis.InitRedis()
	// connect to db
	if err := postgres.InitDB(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Conf.DB_USER, config.Conf.DB_PASSWORD, config.Conf.DB_ADDR, config.Conf.DB_NAME), config.Conf.DB_MIGRATE_VERSION); err != nil {
		log.Fatal(err)
	}
	// start http server
	http_api.StartHttpServer()
}
