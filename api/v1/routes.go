package v1

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.Register)

	return r
}
