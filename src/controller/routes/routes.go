package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karineaf/GoWithReactProject/src/controller"
)

func InitRoutes(routerGroup *gin.RouterGroup){

	routerGroup.POST("/carts/products/", controller.AddProduct) // add a product in cart
	routerGroup.DELETE("/carts/products/:productId", controller.DeleteProductById) // delete a product in cart
	routerGroup.PUT("/carts/products/:productId", controller.UpdateProductById) // update a product in cart
	routerGroup.GET("/carts/quantity/:productId", controller.GetProductQuantityById) // get quantity of products is in cart
	routerGroup.GET("/carts/totalPrice", controller.GetProductsTotalPrice) // get total price

}