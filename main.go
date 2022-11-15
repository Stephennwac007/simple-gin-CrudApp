package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stephennwachukwu/go-RestAPI/controllers"
	"github.com/stephennwachukwu/go-RestAPI/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// server         *gin.Engine
	userservice    helpers.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	// err            error
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	ctx = context.TODO()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Please set MONGODB_URI environment variable")
	}

	mongo_connection := options.Client().ApplyURI(uri)
	mongoclient, err = mongo.Connect(ctx, mongo_connection)
	if err != nil {
		log.Fatal(err)
	}
	//  we need to ping the connection
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MONGODB CONNECTION ESTABLISHED")

	usercollection = (*mongo.Collection)(mongoclient.Database("userdb").Collection("users"))
	userservice = helpers.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)

}

func main() {
	server := gin.Default()
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
	port := os.Getenv("PORT")
	log.Fatal(server.Run(":" + port))
}
