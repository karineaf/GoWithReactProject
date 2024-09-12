package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/karineaf/GoWithReactProject/src/configuration/handle_error"
	"github.com/karineaf/GoWithReactProject/src/controller/model/request"
)

func AddProduct(c *gin.Context) {

	var productRequest request.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		handlerError := handle_error.BadRequestError(
			fmt.Sprintf("Invalid request body. Please provide a valid product object. Error=%s", err.Error()))
		
			c.JSON(handlerError.Code, handlerError)
		return
	}

	fmt.Println(productRequest)
}