package middleware

import (
	"go-rest-project/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		token := parts[1]
		claims, err := auth.ValidateToken(token)
		if err != nil {
			switch err {
			case auth.ErrExpiredToken:
				c.AbortWithStatusJSON(401, gin.H{"error": "Token has expired"})
			case auth.ErrEmptyToken:
				c.AbortWithStatusJSON(401, gin.H{"error": "Token is empty"})
			default:
				c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			}
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
	c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, ok := c.Get("role")
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "User role not found in context"})
			return
		}
		role, _ := roleVal.(string)

		hasRole := false
		for _, allowedRole := range roles {
			if role == allowedRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "Insufficient permissions"})
			return
		}
		c.Next()
	}
}
