package handlers

import (
	"net/http"

	mysql_storage "github.com/fathisiddiqi/go-mini-commerce/storage/mysql"
	redis_storage "github.com/fathisiddiqi/go-mini-commerce/storage/redis"
	"github.com/fathisiddiqi/go-mini-commerce/validator"
	"github.com/gorilla/mux"
)

type API struct {
	storage *mysql_storage.MysqlStorage
	tempStorage *redis_storage.RedisStorage
	validate *validator.Validate
}

type APIConfigs struct {
	Storage *mysql_storage.MysqlStorage
	TempStorage *redis_storage.RedisStorage
	Validate *validator.Validate
}

func NewAPI(configs APIConfigs) (*API, error) {
	return &API{storage: configs.Storage, tempStorage: configs.TempStorage, validate: configs.Validate}, nil
}

func (a *API) GetHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/products", a.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products", a.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/products/{id}", a.GetProductById).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", a.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/products/{id}", a.DeleteProduct).Methods("DELETE")

	return router
}