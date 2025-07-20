package handler

import (
	"gin-study/api/public"
	"gin-study/api/request"
	"gin-study/internal/models/mysql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var (
		register request.Register
		user     mysql.UserInfo
		add      mysql.Users
	)
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.GetUserInfoByName(register.Name); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is exist"})
		return
	}
	addMap := map[string]interface{}{
		"username":   register.Name,
		"password":   public.Md5(register.Password),
		"created_at": time.Now().Format(time.DateTime),
	}
	if err := add.CreateUser(addMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func Login(c *gin.Context) {
	var (
		login request.Login
		user  mysql.UserInfo
	)
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.GetUserInfoByName(login.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Password != public.Md5(login.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password error"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("gin_study"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
