package handler

import (
	"context"
	"minimedium/config"
	"minimedium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @ID sign_up
// @Router /signup [POST]
// @Summary SignUp
// @Description SignUp
// @Tags SignUp
// @Accept json
// @Produce json
// @Param object body models.SignUp true "SignUpRequestBody"
// @Success 200 {object} Response{data=string} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) SignUp(c *gin.Context) {

	var signup models.SignUp

	err := c.ShouldBindJSON(&signup)

	if err != nil {
		HandleResponse(c, 400, "ShouldBindJSON signup")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, _ := h.strg.Users().CreateUser(ctx, &models.CreateUser{
		Name:     signup.Name,
		UserName: signup.UserName,
		Bio:      signup.Bio,
		MediaUrl: signup.MediaUrl,
		Password: signup.Password,
	})

	HandleResponse(c, 200, resp)

}

// SignIn godoc
// @ID sign_in
// @Router /signin [GET]
// @Summary Sign In
// @Description Sign In
// @Tags SignUp
// @Accept json
// @Produce json
// @Param user_name query string false "user_name"
// @Param password query string false "password"
// @Success 200 {object} Response{data=string} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) SignIn(c *gin.Context) {

	var user_name = c.Query("user_name")
	var password = c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	_, err := h.strg.Login().SignIn(ctx, &models.SignIn{UserName: user_name, Password: password})

	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}

	HandleResponse(c, http.StatusOK, "login successful")
}

// Follow godoc
// @ID follow
// @Router /follow [POST]
// @Summary Follow
// @Description Follow
// @Tags Follow
// @Accept json
// @Produce json
// @Param FollowerId query string false "follower_id"
// @Param FollowedId query string false "followed_id"
// @Success 200 {object} Response{data=string} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Follow(c *gin.Context) {
	var follower_id = c.Query("follower_id")
	var followed_id = c.Query("followed_id")

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	err := h.strg.Follow().Follow(ctx, &models.Follow{
		FollowerId: follower_id,
		FollowedId: followed_id,
	})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}
	HandleResponse(c, http.StatusOK, "follow successful")
}
