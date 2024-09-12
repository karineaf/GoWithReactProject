package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/karineaf/GoWithReactProject/src/configuration/validation"
	"github.com/karineaf/GoWithReactProject/src/controller/model/request"
)

func AddProduct(c *gin.Context) {

	var productRequest request.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		handlerError := validation.ValidadeProductError(err)

		c.JSON(handlerError.Code, handlerError)
		return
	}

	fmt.Println(productRequest)
}
