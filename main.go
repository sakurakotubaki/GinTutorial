package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "shop/model"
    "shop/repository"
    "shop/usecase"
    "shop/controller"
)

func main() {
    db, err := gorm.Open("sqlite3", "shop.db")
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&model.Shop{})

    shopRepo := repository.NewShopRepository(db)
    shopUsecase := usecase.NewShopUsecase(shopRepo)
    shopController := controller.NewShopController(shopUsecase)

    r := gin.Default()

    r.POST("/shops", shopController.Create)
	r.GET("/shops/:id", shopController.GetByID)
	r.GET("/shops", shopController.GetAll)
	r.PUT("/shops/:id", shopController.Update)
	r.DELETE("/shops/:id", shopController.Delete)
    r.Run()
}