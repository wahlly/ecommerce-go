package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wahlly/ecommerce-go/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("users/addProduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("users/productView", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}
