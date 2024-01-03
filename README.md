# GinTutorial

1. Goの環境構築をする
```bash
go mod init shop
```

2. Ginを導入する
```bash
go get -u github.com/gin-gonic/gin
```

3. `main.go`を作成して`Hello World`してみる

curlコマンドを使って`http://localhost:8080/`でサーバーからデータをGETできるが、Thunder ClientやPostmanを使うと簡単にローカルサーバーから、HTTP GETできる。
```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello World")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```

4. sqlite3を導入する
```bash
go get github.com/mattn/go-sqlite3
```

5. gormを導入する
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

`go mod tidy`を実行して依存関係を追加する:
```bash
go mod tidy
```

CRUDと特定のデータを検索することができるアプリ:
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Shop struct {
	gorm.Model
	Title       string
	Description string
}

func main() {
	db, err := gorm.Open("sqlite3", "shop.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Shop{})

	r := gin.Default()

	// Create
	r.POST("/shops", func(c *gin.Context) {
		var shop Shop
		if err := c.ShouldBindJSON(&shop); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&shop)
		c.JSON(http.StatusOK, shop)
	})

	// Read
	r.GET("/shops/:id", func(c *gin.Context) {
		var shop Shop
		if err := db.First(&shop, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		c.JSON(http.StatusOK, shop)
	})

	// Read all
	r.GET("/shops", func(c *gin.Context) {
		var shops []Shop
		if err := db.Find(&shops).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		} else {
			db.Find(&shops)
			c.JSON(http.StatusOK, shops)
		}
	})

	// Update
	r.PUT("/shops/:id", func(c *gin.Context) {
		var shop Shop
		if err := db.First(&shop, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&shop); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&shop)
		c.JSON(http.StatusOK, shop)
	})

	// Delete
	r.DELETE("/shops/:id", func(c *gin.Context) {
		var shop Shop
		if err := db.First(&shop, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&shop)
		c.JSON(http.StatusOK, gin.H{"success": "Record has been deleted!"})
	})

	r.Run()
}
```