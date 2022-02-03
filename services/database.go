package services

import (
	"context"
	"fmt"
	"go-mongodb-api/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database() *mongo.Database {
	ctx := context.Background()

	clientUri := "mongodb://" + utils.Config().Database.Host + ":" + utils.Config().Database.Port + "/?ssl=" + utils.Config().Database.Ssl + "&w=majority"
	clientOptions := options.Client().ApplyURI(clientUri).SetAuth(options.Credential{Username: utils.Config().Database.User, Password: utils.Config().Database.Password, AuthSource: "admin"})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[DATABASE] Successfully connected")

	return client.Database(utils.Config().Database.Name)
}
