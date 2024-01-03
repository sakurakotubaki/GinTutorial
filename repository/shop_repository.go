// repository/shop_repository.go
package repository

import (
    "github.com/jinzhu/gorm"
    "shop/model"
)

type ShopRepository interface {
    GetByID(id uint) (*model.Shop, error)
    GetAll() ([]*model.Shop, error)
    Create(shop *model.Shop) error
    Update(shop *model.Shop) error
    Delete(shop *model.Shop) error
}

type shopRepository struct {
    db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
    return &shopRepository{db}
}

func (r *shopRepository) GetByID(id uint) (*model.Shop, error) {
    var shop model.Shop
    if err := r.db.First(&shop, id).Error; err != nil {
        return nil, err
    }
    return &shop, nil
}

func (r *shopRepository) GetAll() ([]*model.Shop, error) {
    var shops []*model.Shop
    if err := r.db.Find(&shops).Error; err != nil {
        return nil, err
    }
    return shops, nil
}

func (r *shopRepository) Create(shop *model.Shop) error {
    return r.db.Create(shop).Error
}

func (r *shopRepository) Update(shop *model.Shop) error {
    return r.db.Save(shop).Error
}

func (r *shopRepository) Delete(shop *model.Shop) error {
    return r.db.Delete(shop).Error
}