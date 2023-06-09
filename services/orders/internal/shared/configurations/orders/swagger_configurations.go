package orders

import (
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (c *ordersServiceConfigurator) configSwagger() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Orders Service Api"
	docs.SwaggerInfo.Description = "Orders Service Api."

	c.echoServer.GetEchoInstance().GET("/swagger/*", echoSwagger.WrapHandler)
}
