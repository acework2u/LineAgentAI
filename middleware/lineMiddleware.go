package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	conf2 "linechat/conf"
	"net/http"
	"strings"
)

func lineMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Next()
			return
		}
		// remove "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")
		if token == "" {
			c.Next()
		}
		cfg, err := conf2.NewAppConfig()
		if err != nil {
			c.Next()
		}

		// Verify line token (simplified; use line 's public key in production)
		_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.LineApp.ChannelSecret), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
