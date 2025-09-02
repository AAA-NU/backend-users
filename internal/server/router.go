package server

import (
	"fmt"

	"github.com/aaanu/backendusers/internal/domain/requests"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router  *gin.RouterGroup
	service UsersService
}

func Register(router *gin.RouterGroup, service UsersService) {
	r := &UserRouter{
		router:  router,
		service: service,
	}
	r.init()
}

func (r *UserRouter) GetUser(ctx *gin.Context) {
	tgID := ctx.Param("tgID")

	user, err := r.service.User(ctx, tgID)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, user)
}

func (r *UserRouter) GetUsers(ctx *gin.Context) {
	role := ctx.Query("role")

	users, err := r.service.Users(ctx, role)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, users)
}

func (r *UserRouter) SaveUser(ctx *gin.Context) {
	var userRequest requests.SaveUserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		HandleError(ctx, err)
		return
	}

	if err := r.service.SaveUser(ctx, &userRequest); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (r *UserRouter) UpdateUser(ctx *gin.Context) {
	tgID := ctx.Param("tgID")
	role := ctx.Query("role")
	fmt.Println(role)
	language := ctx.Query("language")

	if err := r.service.UpdateUser(ctx, tgID, role, language); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (r *UserRouter) DeleteUser(ctx *gin.Context) {
	tgID := ctx.Param("tgID")
	fromUserID := ctx.Query("fromUserID")

	if err := r.service.DeleteUser(ctx, tgID, fromUserID); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (r *UserRouter) init() {
	group := r.router.Group("/users")

	group.GET("/:tgID", r.GetUser)
	group.GET("/", r.GetUsers)
	group.POST("/", r.SaveUser)
	group.PUT("/:tgID", r.UpdateUser)
	group.DELETE("/:tgID", r.DeleteUser)
}
