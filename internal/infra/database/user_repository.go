package database

import (
	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) FindAll(page, limit int, sort string) ([]*entity.User, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	var users []*entity.User
	var err error
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(&users).Error
	} else {
		err = u.DB.Order("created_at" + sort).Find(&users).Error
	}
	return users, err
}

func (u *UserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Update(user *entity.User) error {
	err := u.DB.First(&user, "id = ?", user.ID).Error
	if err != nil {
		return err
	}
	return u.DB.Save(user).Error
}

func (u *UserRepository) Delete(id string) error {
	user, err := u.FindByID(id)
	if err != nil {
		return err
	}
	return u.DB.Delete(user).Error
}
