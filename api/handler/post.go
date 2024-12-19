package handler

import (
	"context"
	"net/http"
	"project/config"
	"project/helpers"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var createPost models.CreatePost
	err := c.ShouldBindJSON(&createPost)
	if err != nil {
		HandleResponse(c, 400, "ShouldBindJSON Error "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.Posts().CreatePost(ctx, &createPost)
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}
	HandleResponse(c, http.StatusCreated, resp)
}

func (h *Handler) GetByIdPost(c *gin.Context) {

	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		HandleResponse(c, http.StatusBadRequest, "it is not a uuid")
		return
	}
	_, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Posts().GetByIdPost(c.Request.Context(), &models.PostPrimaryKey{Id: id})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}
	HandleResponse(c, http.StatusOK, resp)

}
func (h *Handler) GetListPost(c *gin.Context) {
	limit, err := getIntegerOrDefaultValue(c.Query(" limit"), 10)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	page_limit, err := getIntegerOrDefaultValue(c.Query(" offset"), 0)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Posts().GetListPost(ctx, &models.GetListPostRequest{
		Page:   page_limit,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, err)
		return
	}

	HandleResponse(c, http.StatusOK, resp)
}

func (h *Handler) DeletePost(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		HandleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Posts().DeletePost(ctx, &models.PostPrimaryKey{Id: id})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}
	HandleResponse(c, http.StatusNoContent, nil)
}