package routes

import (
	"github.com/arturbaccarin/go-my-first-crud/src/controller"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {

	r.GET("/users/:id", model.VerifyTokenMiddleware, controller.FindUserById)
	r.GET("/users/emails/:email", model.VerifyTokenMiddleware, controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUserById)
	r.DELETE("/users/:id", controller.DeleteUserById)

	r.POST("/login", controller.LoginUser)
}
