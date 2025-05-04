package handler

import (
	"context"
	"minimedium/config"
	"minimedium/helpers"
	"minimedium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param object body models.CreateUser true "CreateUserRequestBody"
// @Success 200 {object} Response{data=models.User} "UserBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		HandleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.Users().CreateUser(ctx, &createUser)
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusCreated, resp)
}
// GetByIdUser godoc
// @ID get_user
// @Router /user/{id} [GET]
// @Summary Get By Id User
// @Description Get By Id User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.User} "GetListUserResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdUser(c *gin.Context) {

	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		HandleResponse(c, http.StatusBadRequest, "it is not a uuid")
		return
	}
	_, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Users().GetByIdUser(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}
	HandleResponse(c, http.StatusOK, resp)

}

// GetListUser godoc
// @ID get_list_user
// @Router /users [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListUserResponse} "GetListUserResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListUser(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query(" limit"), 10)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query(" offset"), 0)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Users().GetListUser(ctx, &models.GetListUserRequest{
		Page:   offset,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusOK, resp)

}
// UpdateUser godoc
// @ID update_user
// @Router /user [PUT]
// @Summary UserUpdate
// @Description UserUpdate
// @Tags User
// @Accept json
// @Produce json
// @Param object body models.UpdateUser true "UpdateUserRequestBody"
// @Success 200 {object} Response{data=models.UpdateUser} "GetListUserResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var updateuser models.UpdateUser

	err := c.ShouldBindJSON(&updateuser)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rows, err := h.strg.Users().UpdateUser(ctx, &models.UpdateUser{
		Id:       updateuser.Id,
		Name:     updateuser.Name,
		UserName: updateuser.UserName,
		Bio:      updateuser.Bio,
		MediaUrl: updateuser.MediaUrl,
		Password: updateuser.Password,
	})

	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusOK, rows)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary UserDelete
// @Description UserDelete
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.UserPrimaryKey} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		HandleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Users().DeleteUser(ctx, &models.UserPrimaryKey{Id: id})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}
	HandleResponse(c, http.StatusNoContent, nil)
}
