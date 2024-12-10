package api

import (
	"project/api/handler"
	"project/config"
	"project/storage"

	"github.com/gin-gonic/gin"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.POST("/users", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/users", handler.GetListUser)
	r.PUT("/users", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

}
