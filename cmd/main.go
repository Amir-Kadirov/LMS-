package main

import (
	"backend_course/lms/api"
	"backend_course/lms/config"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/service"
	"backend_course/lms/storage/postgres"
	"backend_course/lms/storage/redis"
	"context"
	"fmt"
)

func main() {
	cfg := config.Load()
	newRedis := redis.New(cfg)

	store, err := postgres.New(context.Background(), cfg, newRedis)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}

	defer store.CloseDB()

	
	log := logger.New(cfg.ServiceName)
	
	service := service.New(store,log)
	
	c := api.New(service, log)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8080")
}
