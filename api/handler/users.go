package handler

import (
	"context"
	"net/http"
	"project/config"
	"project/helpers"
	"project/models"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) GetListUser(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query(" limit"), 1)
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
		Page:   limit,
		Limit:  offset,
		Search: search,
	})
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusOK, resp)

}

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
		Id:        updateuser.Id,
		FirstName: updateuser.FirstName,
		LastName:  updateuser.LastName,
		Password:  updateuser.Password,
	})

	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusOK, rows)
}

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
