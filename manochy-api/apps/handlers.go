package apps

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDB() {
	if client != nil {
		return
	}

	c, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	client = c
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}

	c.BindJSON(&req)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": req.Username,
	})

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(200, gin.H{"token": t})
}

func GetMembers(c *gin.Context) {
	col := DB.Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, err)
		return
	}

	var result []bson.M
	cursor.All(ctx, &result)

	c.JSON(200, result)
}

func CreateMember(c *gin.Context) {
	col := DB.Collection("members")

	var body bson.M
	c.BindJSON(&body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := col.InsertOne(ctx, body)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}

func UpdateMember(c *gin.Context) {
	col := DB.Collection("members")

	id := c.Param("id")

	var body bson.M
	c.BindJSON(&body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := col.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": body},
	)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}

func DeleteMember(c *gin.Context) {
	col := DB.Collection("members")

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, res)
}
