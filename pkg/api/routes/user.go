package routes

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userHandler *handlers.UserHandler) {
	r.POST("/signup",userHandler.UserSignup)
}
