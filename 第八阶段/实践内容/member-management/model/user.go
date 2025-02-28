package model

import (
	"gorm.io/gorm"
)

// User 定义用户结构体
type User struct {
	// gorm.Model 结构体包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	gorm.Model
	Password   string `gorm:"not null"`
	Nickname   string `gorm:"not null"`
	Phone      string `gorm:"not null"`
	// Email 字段为唯一索引，不允许为空
	Email      string `gorm:"not null;unique"`
	IsAdmin    bool   `gorm:"default:false"`
	IsApproved bool   `gorm:"default:false"`
}