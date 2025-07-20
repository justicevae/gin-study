package handler

import (
	"gin-study/api/request"
	"gin-study/internal/models/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostList(c *gin.Context) {
	var (
		param request.PostList
		list  mysql.PostList
	)
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := list.GetPostList(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "list": list})
}

func GetPostDetail(c *gin.Context) {
	var (
		param request.PostDetail
		info  mysql.PostsInfo
	)
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := info.GetInfoById(param.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": info})
}

func CreatePost(c *gin.Context) {
	var (
		param  request.CreatePost
		create mysql.PostsInfo
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := map[string]interface{}{
		"title":   param.Title,
		"content": param.Content,
	}
	if err := create.CreatePost(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func UpdatePost(c *gin.Context) {
	var (
		param  request.UpdatePost
		info   mysql.PostsInfo
		update mysql.Posts
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := info.GetInfoById(param.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if info.UserId != c.MustGet("uid").(int) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	data := map[string]interface{}{
		"title":   param.Title,
		"content": param.Content,
	}
	if err := update.UpdatePost(param.Id, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeletePost(c *gin.Context) {
	var (
		param  request.DeletePost
		info   mysql.PostsInfo
		update mysql.Posts
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := info.GetInfoById(param.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if info.UserId != c.MustGet("uid").(int) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	data := map[string]interface{}{
		"active": 0,
	}
	if err := update.UpdatePost(param.Id, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
