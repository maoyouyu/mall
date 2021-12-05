package main

import (
	"github.com/gin-gonic/gin"
	"mall/api/v1"
	"mall/common"
	"mall/controller"
)

func main() {
	DB := common.GetDB()
	//defer DB.Close()

	r := gin.Default()
	r = v1.CollectRoute(r)

	panic(r.Run())
}
