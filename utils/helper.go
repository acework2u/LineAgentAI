package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type ErrorHandler struct {
	ctx *gin.Context
}

type APIResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var SecretKey = []byte("")

func NewCustomErrorHandler(ctx *gin.Context) ErrorHandler {
	return ErrorHandler{ctx: ctx}
}

func (c *ErrorHandler) ValidateError(err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]APIResponse, len(ve))
		for i, e := range ve {
			//out = append(out, APIResponse{
			//	Field:   e.Field(),
			//	Message: e.Tag(),
			//})
			out[i] = APIResponse{e.Field(), e.Tag()}
			//c.ctx.JSON(400, gin.H{"error": e.Field() + " " + e.Tag()})
		}
		c.ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": out,
		})
		return

	}
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashPassword), nil
}
func ComparePassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return errors.New("error comparing password")
	}
	return nil
}
func GenerateToken(ttt time.Duration, lindId string, payload interface{}) (string, error) {

	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     lindId,
		"exp":     time.Now().Add(ttt).Unix(),
		"iat":     time.Now().Unix(),
		"payload": payload,
	})
	// Print information about the creat token
	token, err := claims.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}

func GetPayload(token *jwt.Token) (interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["payload"], nil
	}
	return nil, errors.New("invalid token")
}
func GetSub(token *jwt.Token) (string, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}
	return "", errors.New("invalid token")
}
func GetExp(token *jwt.Token) (int64, error) {
	return token.Claims.(jwt.MapClaims)["exp"].(int64), nil
}
func GetIat(token *jwt.Token) (int64, error) {
	return token.Claims.(jwt.MapClaims)["iat"].(int64), nil
}
func GetPayloadString(token *jwt.Token) (string, error) {
	return token.Claims.(jwt.MapClaims)["payload"].(string), nil
}

func GetSubString(token *jwt.Token) (string, error) {
	return token.Claims.(jwt.MapClaims)["sub"].(string), nil
}
func GetExpString(token *jwt.Token) (string, error) {
	return token.Claims.(jwt.MapClaims)["exp"].(string), nil
}
func GetIatString(token *jwt.Token) (string, error) {
	return token.Claims.(jwt.MapClaims)["iat"].(string), nil
}
func HashPasswordWithSalt(password string, secret string) (string, error) {
	salt := []byte(secret)
	combinedHashPassword := append([]byte(password), salt...)
	hashPassword, err := bcrypt.GenerateFromPassword(combinedHashPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashPassword), nil
}

func ComparePasswordWithSalt(password, hashPassword, secret string) error {
	salt := []byte(secret)
	combinedHashPassword := append([]byte(password), salt...)
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), combinedHashPassword)
	if err != nil {
		return errors.New("error comparing password")
	}
	return nil
}

func DateToString(t time.Time) string {
	return t.Format("2006-01-02")

}
func timeToString(t time.Time) string {
	return t.Format("15:04")
}
func DateToTime(t string) time.Time {
	date, _ := time.Parse("2006-01-02", t)
	return date
}
func TimeToTime(t string) time.Time {
	toTime, _ := time.Parse("15:04", t)
	return toTime
}
