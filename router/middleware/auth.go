package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type AuthUserInfo struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func AuthUser(c *gin.Context) {
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
		c.Abort()
		return
	}

	claims := &AuthUserInfo{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("gin_study"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("uid", claims.Id)
	c.Next()
}
