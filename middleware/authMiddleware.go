package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

var secretKey = os.Getenv("SECRET_KEY")

// GenerateToken crate a JWT token
//func GenerateToken(userID string) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"user_id": userID,
//		"ext":     time.Now().Add(time.Hour * 24 * 30).Unix(),
//	})
//	return token.SignedString([]byte(secretKey))
//}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

//func AuthMiddlewares() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("Authorization")
//		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, jwt.ErrSignatureInvalid
//			}
//			return []byte(secretKey), nil
//		})
//		if err != nil || !token.Valid {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//		claims, ok := token.Claims.(jwt.MapClaims)
//		if !ok {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//		c.Set("user_id", claims["user_id"])
//		c.Next()
//	}
//}
