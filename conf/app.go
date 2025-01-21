package conf

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type App struct {
	Port         string `json:"port" yaml:"port"`
	ClientOrigin string `json:"clientOrigin" yaml:"clientOrigin"`
	GinMode      string `json:"ginMode" yaml:"ginMode"`
	SecretKey    string `json:"secretKey" yaml:"secretKey"`
}
type Db struct {
	DbName string `json:"dbname" yaml:"dbname"`
	Port   string `json:"port" yaml:"port"`
	Uri    string `json:"uri" yaml:"uri"`
}
type LineApp struct {
	ChannelSecret string `json:"channelSecret" yaml:"channelSecret"`
	ChannelToken  string `json:"channelToken" yaml:"channelToken"`
}

type AppConfig struct {
	App     App
	Db      Db
	LineApp LineApp
}

func NewAppConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("../../conf/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &AppConfig{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil

}

func ConnectionDB() *mongo.Client {
	cfg, err := NewAppConfig()
	if err != nil {
		log.Println(err)
	}
	log.Println(cfg.Db.Uri)
	// client connected to db
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Uri))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		//log.Println(err)
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	return client

}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	cfg, err := NewAppConfig()
	if err != nil {
		log.Println(err)
	}
	log.Println(cfg.Db.Uri)
	log.Println(cfg.Db.DbName)

	return client.Database(cfg.Db.DbName).Collection(collectionName)
}
