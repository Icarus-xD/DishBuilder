package main

import (
	"log"

	"github.com/Icarus-xD/DishBuilder/internal/config"
	"github.com/Icarus-xD/DishBuilder/internal/database"
	"github.com/Icarus-xD/DishBuilder/internal/repository"
	"github.com/Icarus-xD/DishBuilder/internal/service"
	"github.com/Icarus-xD/DishBuilder/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db := database.Init(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	router := gin.Default()

	dishRepo := repository.New(db) 
	dishService := service.NewDishService(dishRepo)

	handler := rest.NewHandler(dishService)
	handler.InitRouter(router)

	if err = router.Run(config.Port); err != nil {
		log.Fatalln(err)
	}
}

// 15