package router

import (
	"bidfood/internal/auth"
	controller "bidfood/internal/handler"
	"bidfood/internal/logger"
	"bidfood/internal/service"
	"log"

	"github.com/julienschmidt/httprouter"
)

var HttpRouter *Router

var controllerImpl *controller.Controller

type Router struct {
	*httprouter.Router
}

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func New() *Router {

	if controllerImpl == nil {
		productService := service.NewService()
		controllerImpl = controller.NewController(productService)
		log.Println(">>>>>>", controllerImpl)
	}

	if HttpRouter == nil {
		router := httprouter.New()
		routes := AllRoutes()
		for _, route := range routes {
			var handle httprouter.Handle

			handle = route.HandlerFunc
			handle = logger.Logger(handle)
			handle = auth.APIKeyAuthMiddleware(handle)

			router.Handle(route.Method, route.Path, handle)

		}
		return &Router{router}
	}

	return HttpRouter
}

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", controllerImpl.Home},
		Route{"GetAllProduct", "GET", "/products", controllerImpl.GetAllProduct},
		Route{"AddNewProduct", "POST", "/products", controllerImpl.AddNewProduct},
		Route{"GetProductByID", "GET", "/products/:id", controllerImpl.GetProductByID},
		Route{"UpdateProductByID", "PUT", "/products/:id", controllerImpl.UpdateProductByID},
		Route{"DeleteProduct", "DELETE", "/products/:id", controllerImpl.DeleteProduct},

		Route{"InsertDataPagination", "PUT", "/insertdata", controllerImpl.TestAddProduct}, // API insert initial data to test pagination
	}
	return routes
}
