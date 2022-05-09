package main

import (
	"fmt"
	"redigo-example/config"
	"redigo-example/driver"
	"redigo-example/handler"
	"redigo-example/repository"
	"redigo-example/router"
	"redigo-example/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	appConfig, dbConfig, redisConfig := config.InitConfig()
	db, err := driver.InitDB(dbConfig)
	if err != nil {
		log.Fatal("can't connect db")
	}

	redisDB := driver.InitRedis(redisConfig)
	err = driver.Ping(redisDB)
	if err != nil {
		log.Fatal("can't connect redis")
	}

	e := echo.New()

	userRepo := repository.NewUserRepository(db)
	redisRepo := repository.NewRedisRepository(redisDB)
	userService := service.NewUserService(userRepo, redisRepo)
	userHandler := handler.NewUserHandler(userService)

	router.InitUserRouter(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)))

}
