package repository

import (
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (*model.User, error)
	GetByName(name string) *model.User //ユーザー作成とサインイン判定用
	GetAll() ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (u *UserService) GetById(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetByName(name string) *model.User {
	var user model.User
	u.db.Where("username = ?", name).First(&user)
	return &user
}

func (u *UserService) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) Create(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserService) Update(user *model.User) error {
	return u.db.Save(user).Error
}

func (u *UserService) Delete(user *model.User) error {
	return u.db.Delete(user).Error
}
