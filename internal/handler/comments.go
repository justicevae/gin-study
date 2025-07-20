package handler

import (
	"gin-study/api/request"
	"gin-study/internal/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var (
		param  request.CreateComment
		create mysql.Comments
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := map[string]interface{}{
		"user_id": c.GetInt("user_id"),
		"post_id": param.PostID,
		"content": param.Content,
	}
	if err := create.CreateComment(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetCommentList(c *gin.Context) {
	var (
		param request.GetCommentList
		list  mysql.CommentList
	)
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := list.GetCommentList(param.PostID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": "list"})
}
