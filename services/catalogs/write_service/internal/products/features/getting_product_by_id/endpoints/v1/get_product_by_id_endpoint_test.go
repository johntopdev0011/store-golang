package v1

import (
	"context"
	"github.com/gavv/httpexpect/v2"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/test"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/delivery"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/shared/test_fixtures/e2e"
	"net/http"
	"testing"
)

// we could also run the server on docker and then send rest call to the api
func Test_Product_By_Id_E2E(t *testing.T) {
	t.Log("Running Test_Product_By_Id_E2E")
	test.SkipCI(t)
	fixture := e2e.NewE2ETestFixture()

	e := NewGetProductByIdEndpoint(delivery.NewProductEndpointBase(fixture.InfrastructureConfiguration, fixture.V1.ProductsGroup))
	e.MapRoute()

	fixture.Run()
	defer fixture.Cleanup()

	// create httpexpect instance
	expect := httpexpect.New(t, fixture.HttpServer.URL)

	expect.GET("/api/v1/products/{id}").
		WithPath("id", "1b088075-53f0-4376-a491-ca6fe3a7f8fa").
		WithContext(context.Background()).
		Expect().
		Status(http.StatusOK)
}
