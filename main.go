package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fathisiddiqi/go-mini-commerce/handlers"
	mysql_storage "github.com/fathisiddiqi/go-mini-commerce/storage/mysql"
	redis_storage "github.com/fathisiddiqi/go-mini-commerce/storage/redis"
)

const port = ":8080"

const (
	envKeyDBHost = "MYSQL_HOST"
	envKeyDBPort = "MYSQL_PORT"
	envKeyDBUser = "MYSQL_USER"
	envKeyDBPass = "MYSQL_PASS"
	envKeyDBName = "MYSQL_DBNAME"
	envKeyRedisHost = "REDIS_HOST"
	envKeyRedisPort = "REDIS_PORT"
	envKeyRedisPass = "REDIS_PASS"
)

func main() {
	db, err := mysql_storage.NewMyqlStorage(mysql_storage.MySqlStorageConfigs{
		DBHost: os.Getenv(envKeyDBHost),
		DBPort: os.Getenv(envKeyDBPort),
		DBUser: os.Getenv(envKeyDBUser),
		DBPass: os.Getenv(envKeyDBPass),
		DBName: os.Getenv(envKeyDBName),
	})
	if err != nil {
		log.Fatalf("unable to initialize db due to: %v", err)
	}

	redis, err := redis_storage.NewRedisStorage(redis_storage.RedisStorageConfigs{
		RedisHost: os.Getenv(envKeyRedisHost),
		RedisPort: os.Getenv(envKeyRedisPort),
		RedisPass: os.Getenv(envKeyRedisPass),
	})
	if err != nil {
		log.Fatalf("unable to initialize redis due to: %v", err)
	}

	api, err := handlers.NewAPI(handlers.APIConfigs{
		Storage: db,
		TempStorage: redis,
	})
	if err != nil {
		log.Fatalf("unable to initialize api due to: %v", err)
	}
	
	log.Printf("Server running on port %v", port)
	err = http.ListenAndServe(port, api.GetHandler())
	if err != nil {
		log.Fatalf("unable to execute http server due to: %v", err)
	}
}