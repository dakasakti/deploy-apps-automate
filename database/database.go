package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dakasakti/deploy-apps-hexagonal/config"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection(config *config.AppConfig) (*gorm.DB, *mongo.Collection) {
	var db *gorm.DB

	switch config.DB_Driver {
	case "postgree":
		db = initPostgree(config)
	case "mysql":
		db = initMySQL(config)
	}

	db.AutoMigrate(&models.User{})

	client := initMongoDB(config)
	mc := initCollection(client, config, "users")

	return db, mc
}

func initMySQL(config *config.AppConfig) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Address,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func initPostgree(config *config.AppConfig) *gorm.DB {
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		config.DB_Address,
		config.DB_Username,
		config.DB_Password,
		config.DB_Name,
		config.DB_Port,
	)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func initMongoDB(config *config.AppConfig) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(config.Mongo_URI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func initCollection(client *mongo.Client, config *config.AppConfig, name string) *mongo.Collection {
	collection := client.Database(config.Mongo_Database).Collection(name)
	return collection
}
