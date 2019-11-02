package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/henne-/ctdogeusnbuhasys/backend/routes"
	"github.com/sirupsen/logrus"
)

var (
	// config values
	httpListenAddress string

	// internal
	httpServer *http.Server
)

func initServer() {
	router := mux.NewRouter()

	router.NewRoute().Path("/product").Methods("GET").HandlerFunc(routes.ListProducts)
	router.NewRoute().Path("/product").Methods("POST").HandlerFunc(routes.CreateProduct)
	router.NewRoute().Path("/product/{id}").Methods("GET").HandlerFunc(routes.GetProduct)
	router.NewRoute().Path("/product/{id}").Methods("PUT").HandlerFunc(routes.UpdateProduct)

	router.NewRoute().Path("/product/{id}/image").Methods("GET").HandlerFunc(routes.GetProductImage)
	router.NewRoute().Path("/product/{id}/image").Methods("PUT").HandlerFunc(routes.UpdateProductImage)

	router.NewRoute().Path("/category").Methods("GET").HandlerFunc(routes.ListCategories)
	router.NewRoute().Path("/category").Methods("POST").HandlerFunc(routes.CreateCategory)
	router.NewRoute().Path("/category/{id}").Methods("GET").HandlerFunc(routes.GetCategory)
	router.NewRoute().Path("/category/{id}").Methods("PUT").HandlerFunc(routes.UpdateCategory)

	router.NewRoute().Path("/barcode").Methods("GET").HandlerFunc(routes.ListBarcodes)
	router.NewRoute().Path("/barcode").Methods("POST").HandlerFunc(routes.CreateBarcode)
	router.NewRoute().Path("/barcode/{id}").Methods("GET").HandlerFunc(routes.GetBarcode)
	router.NewRoute().Path("/barcode/{id}").Methods("PUT").HandlerFunc(routes.UpdateBarcode)
	router.NewRoute().Path("/barcode/ean/{ean}").Methods("GET").HandlerFunc(routes.GetEanBarcode)

	router.NewRoute().Path("/user").Methods("GET").HandlerFunc(routes.ListUsers)
	router.NewRoute().Path("/user").Methods("POST").HandlerFunc(routes.CreateUser)
	router.NewRoute().Path("/user").Methods("OPTIONS").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})
	router.NewRoute().Path("/user/{id}").Methods("GET").HandlerFunc(routes.GetUser)
	router.NewRoute().Path("/user/{id}").Methods("PUT").HandlerFunc(routes.UpdateUser)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT"})
	headersOk := handlers.AllowedHeaders([]string{"content-type"})
	router.Use(func(handler http.Handler) http.Handler {
		return handlers.CORS(headersOk, originsOk, methodsOk)(handler)
	})
	http.Handle("/", router)
}

func startServer(wg *sync.WaitGroup) {
	if httpServer != nil {
		logrus.Infof("Closing server on `%s`", httpServer.Addr)
		err := httpServer.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}

	httpServer = &http.Server{
		Addr: httpListenAddress,
	}

	wg.Add(1)
	go func(httpServer *http.Server) {
		defer wg.Done()
		logrus.Infof("Starting server on `%s`", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}(httpServer)
}
