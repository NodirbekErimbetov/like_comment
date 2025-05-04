package handler

import (
	"context"
	"minimedium/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Like godoc
// @ID like
// @Router /like [POST]
// @Summary Like
// @Description Like
// @Tags Like
// @Accept json
// @Produce json
// @Param post_id path string false "post_id"
// @Success 200 {object} Response{data=string} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateLike(c *gin.Context) {
	var like models.Like
	var post_id = c.Param("post_id")
	like.Count = 0
	resp, err := h.strg.Like().Like(context.Background(), &models.Like{
		PostId: post_id,
		Count:  like.Count + 1,
	})
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, err)
		return
	}
	HandleResponse(c, http.StatusOK, resp)

}
