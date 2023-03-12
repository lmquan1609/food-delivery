package main

import (
	"food-delivery/component"
	"food-delivery/modules/restauranttransport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

	return r.Run()
}
