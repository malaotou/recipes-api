package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection

func init() {
	var recipe Recipe
	recipe.ID = xid.New().String()
	recipe.Name = "test"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:password@localhost:27017/test?authSource=admin"))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection = client.Database("test").Collection("test")
	result, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/l", ListData)

	router.Run()
}

func ListData(c *gin.Context) {
	//collection := client.Database("test").Collection("test")
	if collection != nil {
		log.Println(collection.Name())
	}
	cur, err := collection.Find(ctx, bson.M{})
	log.Println(cur)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	defer cur.Close(ctx)
	recipes = make([]Recipe, 0)
	for cur.Next(ctx) {
		var recipe Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

func IndexHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PulishedAt   time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PulishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}


func demo() {
	router := gin.Default()
	router.Use(AuthMiddleware())
	router.Group("/admin", gin.BasicAuth(gin.Accounts{}))

}

func AuthMiddleware() gin.HandlerFunc {
	token := jwt.
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != os.Getenv("X_API_KEY") {
			c.AbortWithStatus(401)
			c.MustGet(gin.AuthUserKey)
			gin.Recovery()
		}
	}
	

}
func test(){
		
}