package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-static-server/utils"
	"time"
)

func Init() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@testcluster.5xx4g.mongodb.net/test", utils.GoDotEnvVariable("MONGO_USER_NAME"), utils.GoDotEnvVariable("MONGO_PASSWORD"))))

	if err != nil {
		panic(err)
	}

	defer Disconnect(ctx, client)
	return client

}

func Disconnect(ctx context.Context,  client *mongo.Client) {
	client.Disconnect(ctx);
}



