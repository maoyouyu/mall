package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null"`
	Password string `gorm:"size;not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("pssword")

		// 数据验证
		if len(phone) != 11 {
			ctx.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "手机号必须为11味"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "密码不能为空"})
			return
		}
		// 如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, phone, password)
		// 判断手机号是否存在

		if isPhoneExists(db, phone) {
			ctx.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		// 创建用户
		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newUser)
		// 返回结果

		ctx.JSON(200, gin.H{"msg": "注册成功"})
	})
	panic(r.Run())
}

func isPhoneExists(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklklwerytyyui")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := 3306
	datebase := "test"
	username := "root"
	password := "12345678"
	charset := "utf8mb4"
	//args:=fmt.Sprintf('%s:%s@tcp(%s:%s)/%s?charset=%s')
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, datebase, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("发现错误" + err.Error())
	}
	return db
}
