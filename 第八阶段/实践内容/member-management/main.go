package main

import (
	"member-management/controller"
	"member-management/model"
	"member-management/repository"
	"member-management/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 配置数据库连接信息
	dsn := "tizi6666:Tizi@1234@tcp(192.168.228.128:3306)/member_management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移数据库表结构
	db.AutoMigrate(&model.User{})

	// 直接初始化 repository 层实例
	userRepo := &repository.UserRepository{
		Db: db,
	}
	// 直接初始化 service 层实例
	userService := &service.UserService{
		UserRepo: userRepo,
	}

	// 直接初始化 controller 层实例
	userController := &controller.UserController{
		UserService: userService,
	}

	// 创建 Gin 引擎
	r := gin.Default()
	r.POST("/register", userController.Register)
	r.PUT("/user/update_info", userController.UpdateInfo)
	r.DELETE("/admin/delete_user/:user_id", userController.DeleteUser)
	r.GET("/admin/get_all_users", userController.GetAllUsers)
	r.PUT("/admin/approve_user/:user_id", userController.ApproveUser)

	// 启动服务器
	r.Run(":8080")
}
