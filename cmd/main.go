package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"linechat/conf"
	"log"
	"time"
)

const (
	customerCollection = "customers"
)

var (
	server    *gin.Engine
	client    *mongo.Client
	ctx       context.Context
	appConfig *conf.AppConfig
)

func init() {
	var err error
	appConfig, err = conf.NewAppConfig()
	if err != nil {
		fmt.Println(err)
	}
	//client = conf.ConnectionDB()
	//ctx = context.Background()
	server = gin.Default()
}

func StartServer() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{appConfig.App.ClientOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.MaxAge = 12 * 30 * 24 * time.Hour
	// new config
	server.Use(cors.New(corsConfig))
	server.Use(gin.Recovery())
	server.Use(gin.Logger())

	// default page not found
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})
	//server.Run(appConfig.App.Port)
	log.Fatal(server.Run(":" + appConfig.App.Port + ""))

}

func main() {
	fmt.Println("hello")
	StartServer()

}
