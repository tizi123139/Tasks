package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
	"fmt"
)

type User struct {
    ID uint
    Name  string
    Email string
}

func connectDB() (*gorm.DB,error) {
    //配置数据库连接
	dsn := "tizi6666:Tizi@1234@tcp(192.168.228.128:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
    //连接数据库
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        return nil, err
    }
    return db, nil
}

//创建用户
func CreateUser(db *gorm.DB, user User) (User,error) {
	result := db.Create(&user)
    if result.Error != nil {
        return User{}, result.Error
    }
    return user, nil
}

//更新用户信息
func UpdateUser(db *gorm.DB, id uint, updatedUser User) (User, error) {
    var user User
    //获取用户信息
    result := db.First(&user, id)
    if result.Error != nil {
        return User{}, result.Error
    }

    // 更新用户信息
    // 如果更新的字段为空，则不更新
    if updatedUser.Name != "" {
        user.Name = updatedUser.Name
    }
    if updatedUser.Email != "" {
        user.Email = updatedUser.Email
    }

    // 保存更新后的用户信息
    result = db.Save(&user)
    if result.Error != nil {
        return User{}, result.Error
    }
    return user, nil
}

//通过id获取用户信息
func GetUserByID(db *gorm.DB, id uint) (User, error) {
    var user User
    //获取用户信息
    result := db.First(&user, id)
    if result.Error != nil {
        return User{}, result.Error
    }
    return user, nil
}

//删除用户
// func DeleteUser(db *gorm.DB, id uint) error {
//     var user User
//     result := db.Delete(&user, id)
//     if result.Error != nil {
//         return result.Error
//     }
//     return nil
// }

func main() {
	//连接数据库
    db, err := connectDB()
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    log.Println("Successfully connected to the database")

    // 自动迁移模型
    if err := db.AutoMigrate(&User{}); err != nil {
        log.Fatalf("failed to migrate table: %v", err)
    }
    log.Println("Table migration completed successfully")

    // 创建一个新用户
    newUser := User{Name: "Wu Zhiyi", Email: "wuzhiyi@example.com"}
	createdUser, err := CreateUser(db, newUser)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }
    fmt.Printf("Created user: %+v\n", createdUser)

    log.Println("User created successfully!")
	fetchedUser, err := GetUserByID(db, createdUser.ID)
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }
    fmt.Printf("Fetched user: %+v\n", fetchedUser)

    // 更新用户信息
    updatedUserInfo := User{
        Name:  "Tizi",
        Email: "tizi@example.com",
    }
    updatedUser, err := UpdateUser(db, createdUser.ID, updatedUserInfo)
    if err != nil {
        log.Fatalf("Failed to update user: %v", err)
    }
    fmt.Printf("Updated user: %+v\n", updatedUser)

    //删除用户
	// err = DeleteUser(db, createdUser.ID)
    // if err != nil {
    //     log.Fatalf("Failed to delete user: %v", err)
    // }
    // fmt.Println("User deleted successfully")
}