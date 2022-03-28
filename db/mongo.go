package db

import (
	"DriverLocation/exception"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connection() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	option := options.Client().
		ApplyURI(getURL()).
		SetMinPoolSize(10).
		SetMaxPoolSize(50).
		SetMaxConnIdleTime(time.Duration(30) * time.Second)

	client, err := mongo.NewClient(option)
	exception.IsPanic(err)

	err = client.Connect(ctx)
	exception.IsPanic(err)
	database := client.Database("test")

	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func getURL() string {
	if viper.GetBool("isLocal") {
		return viper.GetString("database.local")
	} else {
		return fmt.Sprintf("mongodb://%s:%s@%s:%d",
			viper.GetString("database.user"),
			viper.GetString("database.password"),
			viper.GetString("database.host"),
			viper.GetInt("database.port"))
	}
}
