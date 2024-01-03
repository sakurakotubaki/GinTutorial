// usecase/shop_usecase.go
package usecase

import (
    "shop/model"
    "shop/repository"
)

type ShopUsecase interface {
    GetByID(id uint) (*model.Shop, error)
    GetAll() ([]*model.Shop, error)
    Create(shop *model.Shop) error
    Update(shop *model.Shop) error
    Delete(shop *model.Shop) error
}

type shopUsecase struct {
    shopRepo repository.ShopRepository
}

func NewShopUsecase(shopRepo repository.ShopRepository) ShopUsecase {
    return &shopUsecase{shopRepo}
}

func (u *shopUsecase) GetAll() ([]*model.Shop, error) {
    return u.shopRepo.GetAll()
}

func (u *shopUsecase) GetByID(id uint) (*model.Shop, error) {
    return u.shopRepo.GetByID(id)
}

func (u *shopUsecase) Create(shop *model.Shop) error {
    return u.shopRepo.Create(shop)
}

func (u *shopUsecase) Update(shop *model.Shop) error {
    return u.shopRepo.Update(shop)
}

func (u *shopUsecase) Delete(shop *model.Shop) error {
    return u.shopRepo.Delete(shop)
}