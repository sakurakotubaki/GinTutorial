// model/shop.go
package model

import "github.com/jinzhu/gorm"

type Shop struct {
    gorm.Model
    Title       string
    Description string
}