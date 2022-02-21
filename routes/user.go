package route

import (
	"net/http"

	controller "github.com/bansaltushar014/GoLang_CRUD/controllers"
	"github.com/gorilla/mux"
)

func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", controller.HandleGetOrder).Methods("GET")
	muxRouter.HandleFunc("/", controller.HandleWriteOrder).Methods("POST")
	muxRouter.HandleFunc("/", controller.HandleUpdateOrder).Methods("PUT")
	return muxRouter
}
