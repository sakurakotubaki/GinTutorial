// controller/shop_controller.go
package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "shop/model"
    "shop/usecase"
    "strconv"
)

type ShopController struct {
    shopUsecase usecase.ShopUsecase
}

func NewShopController(shopUsecase usecase.ShopUsecase) *ShopController {
    return &ShopController{shopUsecase}
}

func (c *ShopController) Create(ctx *gin.Context) {
    var shop model.Shop
    if err := ctx.ShouldBindJSON(&shop); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.shopUsecase.Create(&shop); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, shop)
}

func (c *ShopController) GetByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    shop, err := c.shopUsecase.GetByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, shop)
}

func (c *ShopController) GetAll(ctx *gin.Context) {
    shops, err := c.shopUsecase.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, shops)
}

func (c *ShopController) Update(ctx *gin.Context) {
    var shop model.Shop
    if err := ctx.ShouldBindJSON(&shop); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.shopUsecase.Update(&shop); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, shop)
}

func (c *ShopController) Delete(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    shop, err := c.shopUsecase.GetByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := c.shopUsecase.Delete(shop); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}