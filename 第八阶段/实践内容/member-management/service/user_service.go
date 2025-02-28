package service

import (
	"errors"
	"member-management/model"
	"member-management/repository"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

// UserService 结构体，封装用户服务相关操作
type UserService struct {
	UserRepo *repository.UserRepository
}

// Register 方法，处理用户注册逻辑
func (s *UserService) Register(user *model.User) error {
	// 简单的输入验证，确保关键信息不为空
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" ||
		strings.TrimSpace(user.Nickname) == "" || strings.TrimSpace(user.Phone) == "" {
		return errors.New("email, password, nickname and phone cannot be empty")
	}
	// 对用户输入的密码进行加密处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	// 调用 repository 层的方法创建用户
	return s.UserRepo.CreateUser(user)
}

// Login 方法，处理用户登录逻辑
func (s *UserService) Login(email, password string) (*model.User, error) {
	// 简单的输入验证，确保邮箱和密码不为空
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, errors.New("email and password cannot be empty")
	}
	// 根据邮箱查找用户
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// 验证用户输入的密码与数据库中加密后的密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password or not approved")
	}
	// 检查用户是否已通过审核
	if !user.IsApproved {
		return nil, errors.New("invalid email or password or not approved")
	}
	return user, nil
}

// UpdateInfo 方法，处理用户信息更新逻辑
func (s *UserService) UpdateInfo(userID uint, input struct {
	// 允许用户更新昵称、电话、邮箱和密码
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}) error {
	// 简单的输入验证，确保用户 ID 不为空
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	// 普通用户只能更新自己的信息
	if user.IsAdmin {
		return errors.New("this method is for normal users only")
	}
	// 更新用户信息
	if input.Nickname != "" {
		user.Nickname = input.Nickname
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}
	if input.Email != "" {
		// 可添加邮箱唯一性检查逻辑
		user.Email = input.Email
	}
	// 更新密码时对密码进行加密处理
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return s.UserRepo.UpdateUser(user)
}

// DeleteUser 方法，管理员删除用户
func (s *UserService) DeleteUser(adminID, userID uint) error {
	admin, err := s.UserRepo.GetUserByID(adminID)
	if err != nil || !admin.IsAdmin {
		return errors.New("unauthorized access")
	}
	// 管理员不能删除自己
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user.IsAdmin {
		return errors.New("user not found or is an admin")
	}
	return s.UserRepo.DeleteUser(user)
}

// GetAllUsers 方法，管理员获取所有非管理员用户信息
func (s *UserService) GetAllUsers(adminID uint) ([]model.User, error) {
	admin, err := s.UserRepo.GetUserByID(adminID)
	if err != nil || !admin.IsAdmin {
		return nil, errors.New("unauthorized access")
	}
	return s.UserRepo.GetAllUsers()
}

// ApproveUser 方法，管理员批准用户注册
func (s *UserService) ApproveUser(adminID, userID uint) error {
	admin, err := s.UserRepo.GetUserByID(adminID)
	if err != nil || !admin.IsAdmin {
		return errors.New("unauthorized access")
	}
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user.IsApproved {
		return errors.New("user not found or already approved")
	}
	// 批准用户注册
	user.IsApproved = true
	return s.UserRepo.UpdateUser(user)
}