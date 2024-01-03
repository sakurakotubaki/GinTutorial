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