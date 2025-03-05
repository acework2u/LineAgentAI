package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"linechat/conf"
	"linechat/handler"
	"linechat/repository"
	"linechat/router"
	"linechat/services"
	"log"
	"time"
)

const (
	MemberCollectionName   = "members"
	EventsCollectionName   = "events"
	StaffCollectionName    = "staffs"
	SettingsCollectionName = "settings"
)

var (
	server    *gin.Engine
	client    *mongo.Client
	ctx       context.Context
	appConfig *conf.AppConfig

	// Member Collection
	memberCollection   *mongo.Collection
	settingsCollection *mongo.Collection
	// LineApp
	LineHandler       *handler.LineWebhookHandler
	LineRouter        *router.LineRouter
	lineBotService    services.LineBotService
	memberRepo        repository.MemberRepository
	memberService     services.MemberService
	memberHandler     *handler.MemberHandler
	memberRouter      *router.MemberRouter
	eventsRepo        repository.EventsRepository
	eventsService     services.EventsService
	eventHandler      *handler.EventHandler
	eventRouter       *router.EventRouter
	reportService     services.ReportService
	reportHandler     *handler.ReportHandler
	reportRouter      *router.ReportRouter
	staffRepo         repository.StaffRepository
	staffService      services.StaffService
	staffHandler      *handler.StaffHandler
	staffRouter       *router.StaffRouter
	settingRepo       repository.AppSettingsRepository
	appSettingService services.AppSettingsService
	appSettingHandler *handler.AppSettingHandler
	appSettingRouter  *router.AppSettingRouter
	fileServeHandler  *handler.FileServeHandler
	fileServeRouter   *router.FileServeRouter
)

func init() {
	var err error
	appConfig, err = conf.NewAppConfig()
	ctx = context.TODO()
	if err != nil {
		fmt.Println(err)
	}
	// DB Connection
	client = conf.ConnectionDB()
	memberCollection = conf.GetCollection(client, MemberCollectionName)
	eventsCollection := conf.GetCollection(client, EventsCollectionName)
	staffCollection := conf.GetCollection(client, StaffCollectionName)
	settingsCollection = conf.GetCollection(client, SettingsCollectionName)

	// Service

	memberRepo = repository.NewMemberRepository(ctx, memberCollection)
	eventsRepo = repository.NewEventRepository(ctx, eventsCollection)
	lineBotService = services.NewLineBotService(memberRepo, eventsRepo)
	LineHandler = handler.NewLineWebhookHandler(lineBotService)
	LineRouter = router.NewLineRouter(LineHandler)

	eventsService = services.NewEventsService(eventsRepo)
	eventHandler = handler.NewEventHandler(eventsService)
	eventRouter = router.NewEventRouter(eventHandler)

	// Report
	reportService = services.NewReportService(eventsRepo, memberRepo)
	reportHandler = handler.NewReportHandler(reportService)
	reportRouter = router.NewReportRouter(reportHandler)

	// Staff
	staffRepo = repository.NewStaffRepository(ctx, staffCollection)
	staffService = services.NewStaffService(staffRepo)
	staffHandler = handler.NewStaffHandler(staffService)
	staffRouter = router.NewStaffRouter(staffHandler)
	//Members
	memberService = services.NewMemberService(memberRepo)
	memberHandler = handler.NewMemberHandler(memberService)
	memberRouter = router.NewMemberRouter(memberHandler)
	//Settings
	settingRepo = repository.NewSettingsRepository(ctx, settingsCollection)
	appSettingService = services.NewAppSettingsService(settingRepo)
	appSettingHandler = handler.NewAppSettingHandler(appSettingService)
	appSettingRouter = router.NewAppSettingRouter(appSettingHandler)
	// File serve
	fileServeHandler = handler.NewFileServHandler()
	fileServeRouter = router.NewFileServeRouter(fileServeHandler)

	// Set server
	server = gin.Default()
}

func StartServer() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*", "http://localhost:5173", "https://f325fcd7ea2b.ngrok.app"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.MaxAge = 12 * 30 * 24 * time.Hour
	// new config
	server.Use(cors.New(corsConfig))
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	//server.Use(middleware.AuthMiddleware())
	// default page not found
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	//static page
	server.LoadHTMLGlob("views/*.html")
	server.Static("/static", "./static")
	server.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	server.GET("/attend", func(c *gin.Context) {
		c.HTML(200, "attend.html", nil)
	})
	server.GET("/events", func(c *gin.Context) {
		c.HTML(200, "event.html", nil)
	})
	server.GET("/events-calendar", func(c *gin.Context) {
		c.HTML(200, "events-calendar.html", nil)
	})
	server.GET("/event-cal", func(c *gin.Context) {
		c.HTML(200, "event-cal.html", nil)
	})
	server.GET("/news-event", func(c *gin.Context) {
		c.HTML(200, "news-event-calendar.html", nil)
	})
	server.GET("/event-check-in", func(c *gin.Context) {
		c.HTML(200, "event-checkin.html", nil)
	})
	server.GET("/event-attend", func(c *gin.Context) {
		c.HTML(200, "event-attend.html", nil)
	})

	// router
	routers := server.Group("/api/v1")
	LineRouter.LineHookRouter(routers)
	eventRouter.EventRouter(routers)
	reportRouter.ReportRouter(routers)
	staffRouter.StaffRouter(routers)
	memberRouter.MemberRouter(routers)
	appSettingRouter.AppSettingRouter(routers)
	fileServeRouter.FileServeRouter(routers)

	//server.Run(appConfig.App.Port)
	log.Fatal(server.Run(":" + appConfig.App.Port + ""))

}

func main() {
	StartServer()

}
