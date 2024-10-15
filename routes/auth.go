package routes

import (
	"go-auth/contrrollers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", contrrollers.Login)
	r.POST("/signup", contrrollers.Signup)
	r.GET("/home", contrrollers.Home)
	r.GET("/premium", contrrollers.Premium)
	r.GET("/logout", contrrollers.Logout)
}
