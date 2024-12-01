package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/homeanter/codly/config"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头获取 JWT
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "No Authorization header provided",
			})
			c.Abort()
			return
		}

		// JWT 通常在 Bearer Token 中，需要分割
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(401, gin.H{
				"error": "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		// 解析 JWT
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return config.Config.JwtSecretKey, nil
		})

		if err != nil {
			c.JSON(401, gin.H{
				"error": "Invalid JWT",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"].(uint))
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid JWT",
			})
			c.Abort()
			return
		}
	}
}
