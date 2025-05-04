package api

import (
	"minimedium/api/handler"
	"minimedium/config"
	"minimedium/storage"

	_ "minimedium/api/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/users", handler.GetListUser)
	r.PUT("/user", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/post", handler.CreatePost)
	r.GET("/post/:id", handler.GetByIdPost)
	r.GET("/posts", handler.GetListPost)
	r.PUT("/post", handler.UpdatePost)
	r.DELETE("/post/:id", handler.DeletePost)

	r.POST("/signup", handler.SignUp)
	r.GET("/signin", handler.SignIn)

	r.POST("/follow", handler.Follow)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
