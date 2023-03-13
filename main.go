package main

import (
	"food-delivery/component"
	"food-delivery/middleware"
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
	appCtx := component.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	restaurants := r.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

	return r.Run()
}
