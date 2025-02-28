package controller

import (
	"member-management/model"
	"member-management/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// UserController 结构体，封装用户控制器相关操作
type UserController struct {
	UserService *service.UserService
}

// Register 方法，处理用户注册请求
func (c *UserController) Register(ctx *gin.Context) {
	var user model.User
	// 从请求中解析 JSON 数据并绑定到 user 结构体
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data: " + err.Error()})
		return
	}
	// 调用 UserService 的 Register 方法处理用户注册逻辑
	if err := c.UserService.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Registration failed: " + err.Error()})
		return
	}
	// 返回注册成功的消息
	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful, waiting for approval."})
}

// Login 方法，处理用户登录请求
func (c *UserController) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// 从请求中解析 JSON 数据并绑定到 input 结构体
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data: " + err.Error()})
		return
	}
	// 调用 UserService 的 Login 方法处理用户登录逻辑
	user, err := c.UserService.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed: " + err.Error()})
		return
	}
	// 返回登录成功的消息
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": user.ID, "is_admin": user.IsAdmin})
}

// UpdateInfo 方法，处理用户信息更新请求
func (c *UserController) UpdateInfo(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	// 将用户 ID 从字符串转换为 uint 类型
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	var input struct {
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data: " + err.Error()})
		return
	}
	// 调用 UserService 的 UpdateInfo 方法处理用户信息更新逻辑
	if err := c.UserService.UpdateInfo(uint(userID), input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Update failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User information updated successfully"})
}

// DeleteUser 方法，处理管理员删除用户请求
func (c *UserController) DeleteUser(ctx *gin.Context) {
	adminIDStr := ctx.Query("admin_id")
	// 将管理员 ID 从字符串转换为 uint 类型
	adminID, err := strconv.ParseUint(adminIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin ID"})
		return
	}
	// 将用户 ID 从 URL 参数中解析并转换为 uint 类型
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	// 调用 UserService 的 DeleteUser 方法处理管理员删除用户逻辑
	if err := c.UserService.DeleteUser(uint(adminID), uint(userID)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetAllUsers 方法，处理管理员获取所有非管理员用户信息请求
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	adminIDStr := ctx.Query("admin_id")
	adminID, err := strconv.ParseUint(adminIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin ID"})
		return
	}
	// 调用 UserService 的 GetAllUsers 方法获取所有非管理员用户信息
	users, err := c.UserService.GetAllUsers(uint(adminID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// ApproveUser 方法，处理管理员批准用户注册请求
func (c *UserController) ApproveUser(ctx *gin.Context) {
	adminIDStr := ctx.Query("admin_id")
	adminID, err := strconv.ParseUint(adminIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin ID"})
		return
	}
	// 将用户 ID 从 URL 参数中解析并转换为 uint 类型
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	// 调用 UserService 的 ApproveUser 方法处理管理员批准用户注册逻辑
	if err := c.UserService.ApproveUser(uint(adminID), uint(userID)); err != nil {
		// 返回错误信息
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User approved successfully"})
}
