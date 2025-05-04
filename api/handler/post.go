package handler

import (
	"context"
	"minimedium/config"
	"minimedium/helpers"
	"minimedium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost godoc
// @ID create_post
// @Router /post [POST]
// @Summary Create Post
// @Description Create Post
// @Tags Post
// @Accept json
// @Produce json
// @Param object body models.CreatePost true "CreatePostRequestBody"
// @Success 200 {object} Response{data=models.Post} "PostBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
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

// GetByIdPost godoc
// @ID get_post
// @Router /post/{id} [GET]
// @Summary Get By Id Post
// @Description Get By Id Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Post} "GetListPostResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
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
// GetListPost godoc
// @ID get_list_post
// @Router /posts [GET]
// @Summary Get List Post
// @Description Get List Post
// @Tags Post
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param search query string false "search"
// @Success 200 {object} Response{data=models.GetListPostResponse} "GetListPostResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
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

// UpdatePost godoc
// @ID update_post
// @Router /post [PUT]
// @Summary PostUpdate
// @Description PostUpdate
// @Tags Post
// @Accept json
// @Produce json
// @Param object body models.UpdatePost true "UpdatePostRequestBody"
// @Success 200 {object} Response{data=models.UpdatePost} "GetListPostResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func(h *Handler) UpdatePost(c *gin.Context){
	var updadepost models.UpdatePost

	result,err := h.strg.Posts().UpdatePost(context.Background(),&models.UpdatePost{
		Id: updadepost.Id,
		Title: updadepost.Title,
		Body: updadepost.Body,
	})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}
	HandleResponse(c,http.StatusOK,result)

}

// DeletePost godoc
// @ID delete_post
// @Router /post/{id} [DELETE]
// @Summary PostDelete
// @Description PostDelete
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.PostPrimaryKey} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
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
