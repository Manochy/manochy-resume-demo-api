package apps

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitMongo() {

	// URI : mongodb+srv://portfolio_access_user:wF2huzzJokGcC1hZ@personal-cluster0.zs0hkbl.mongodb.net/?appName=Personal-Cluster0
	if DB != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("manochy_db")

	log.Println("Mongo connected")
}
