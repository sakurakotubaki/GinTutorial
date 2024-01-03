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

## レイヤーを分ける
同じファイルにコードを書いているとコードの記述量が増えるのと、保守がしづらい!
クリーンアーキテクチャで設計をしてみた。
```
.
├── README.md
├── controller
│   └── shop_controller.go
├── go.mod
├── go.sum
├── main.go
├── model
│   └── shop.go
├── repository
│   └── shop_repository.go
├── shop.db
└── usecase
    └── shop_usecase.go
```

このプロジェクトは、クリーンアーキテクチャの原則に基づいてレイヤー分けされています。各ディレクトリの役割は以下の通りです：

controller: 

このディレクトリには、HTTPリクエストを受け取り、適切なユースケースを呼び出し、その結果をHTTPレスポンスとして返すコントローラが含まれています。

model:

 このディレクトリには、アプリケーションのビジネスロジックを表すデータモデルが含まれています。これらのモデルは、データベースのテーブルを表現するために使用されます。

repository:

このディレクトリには、データベースとのやり取りを抽象化するリポジトリインターフェースとその実装が含まれています。これにより、ビジネスロジックはデータベースの具体的な実装から分離されます。

usecase: 

このディレクトリには、アプリケーションのビジネスロジックを表すユースケースが含まれています。ユースケースは、リポジトリを通じてデータベースとやり取りを行います。

main.go: 

このファイルはアプリケーションのエントリーポイントで、全ての依存関係をワイヤリングし、サーバーを起動します。

shop.db: 

これはアプリケーションのデータベースファイルです。

go.modとgo.sum: 

これらのファイルはGoの依存関係管理に使用されます。go.modはプロジェクトの依存関係を宣言し、go.sumは各依存関係の期待されるコンテンツを提供します。

README.md: このファイルはプロジェクトのドキュメンテーションを提供します。