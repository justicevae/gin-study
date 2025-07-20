package main

import (
	"gin-study/internal/handler"
	"gin-study/internal/models/mysql"
	"gin-study/router/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	mysql.GetDB()
	r := gin.Default()

	// 登录相关
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	// 文章
	post := r.Group("/posts")
	{
		post.GET("/list", handler.GetPostList)
		post.GET("/detail", handler.GetPostDetail)

		// 认证
		authPosts := post.Group("/").Use(middleware.AuthUser)
		{
			authPosts.POST("/create", handler.CreatePost)
			authPosts.POST("/update", handler.UpdatePost)
			authPosts.POST("/delete", handler.DeletePost)

			// 评论文章
			authPosts.POST("/create/comment", handler.CreateComment)
			authPosts.POST("/comment/list", handler.GetCommentList)
		}
	}

	// 启动服务器
	r.Run(":8080")
}
