package repository

import (
	"member-management/model"

	"gorm.io/gorm"
)

// UserRepository 结构体，用于封装数据库操作
type UserRepository struct {
	Db *gorm.DB
}

// CreateUser 方法，用于在数据库中创建新用户
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.Db.Create(user).Error
}

// GetUserByEmail 方法，根据邮箱查找用户
func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	// 根据邮箱查找用户
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	// 定义一个 user 变量
	var user model.User
	// 根据 ID 查找用户
	err := r.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 方法，更新用户信息
func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.Db.Save(user).Error
}

// DeleteUser 方法，删除用户
func (r *UserRepository) DeleteUser(user *model.User) error {
	return r.Db.Delete(user).Error
}

// GetAllUsers 方法，获取所有非管理员用户信息
func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	// 查询所有非管理员用户
	err := r.Db.Where("is_admin = false").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
