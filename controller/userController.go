package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mall/common"
	"mall/model"
	"mall/utils"
	"net/http"
)

func Register(ctx *gin.Context) {

	DB := common.GetDB()
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
		name = utils.RandomString(10)
	}
	log.Println(name, phone, password)
	// 判断手机号是否存在

	if isPhoneExists(DB, phone) {
		ctx.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)
	// 返回结果

	ctx.JSON(200, gin.H{"msg": "注册成功"})
}

func isPhoneExists(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
