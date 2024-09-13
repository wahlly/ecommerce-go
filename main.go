package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wahlly/ecommerce-go/controllers"
	"github.com/wahlly/ecommerce-go/database"
	"github.com/wahlly/ecommerce-go/middleware"
	"github.com/wahlly/ecommerce-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client))

	router := gin.New()
	router.Use((gin.Logger()))

	routes.UserRoutes(router)
	router.GET(middleware.Authentication)

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.Removeitem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy)

	log.Fatal(router.Run(":" + port))
}