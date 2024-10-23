package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wahlly/ecommerce-go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

	type Application struct{
		prodCollection *mongo.Collection
		UserCollection *mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
		return &Application{
			prodCollection: prodCollection,
			userCollection: userCollection
		}
	}


func (app *Application) AddToCart() gin.HandlerFunc{
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is ready"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err != primitive.ObjectIDFromHex(productQueryId)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.statusInternal sellor)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection,, app.userCollection, productID, userqueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfuly added to the cart")
	}
}

func (app *Application) Removeitem() gin.HandlerFunc{
	return func (c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is ready"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err != primitive.ObjectIDFromHex(productQueryId)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.statusInternal sellor)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.UserCollection, productQueryId, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "Successfully remove item from cart")
	}
}

func GetItemFromCart() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledcart models.UserID
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: usert_id}}).Decode()
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "not found")
			return
		}

		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total",  Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}

		var listing []bson.M
		if err = pointCursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(http.StatusOK, json["total"])
			c.IndentedJSON(http.StatusOK, filledcart.UserCart)
		}
		ctx.Done()
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc{
	return func(c *gin.Context) {
		userQueryId = c.Query("id")
		if userQueryId == "" {
			log.Panic("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.UserCollection, app.prodCollection, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON("successfully placed the order")

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc{
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is ready"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err != primitive.ObjectIDFromHex(productQueryId)

		if err != nil{
			log.Println(err)
			c.AbortWithStatus(http.statusInternal sellor)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.UserCollection, productQueryId, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "successfully placed the order")
	}
}