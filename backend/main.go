package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/Access-Control-list/backend/config"
	"github.com/Access-Control-list/backend/driver"
	"github.com/Access-Control-list/backend/handler"
	httpHandler "github.com/Access-Control-list/backend/handler/http"
)

var (
	handlers = []handler.IHTTPHandler{}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbConn, err := driver.NewMysqlConnection(config.Config().Database)
	if nil != err {
		log.Printf("Error while creating db connectiion:%s", err.Error())
		os.Exit(1)
	}

	handlers = []handler.IHTTPHandler{
		httpHandler.NewUserHandler(dbConn),
		httpHandler.NewGroupHandler(dbConn),
		httpHandler.NewFileHandler(dbConn),
		httpHandler.NewFolderHandler(dbConn),
	}
}

func createRouterGroup(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		for _, hdlr := range handlers { // register all handlers
			for _, hlr := range hdlr.GetHTTPHandler() {
				path := fmt.Sprintf("/webapi/v1/%s", hlr.Path)
				switch hlr.Method {
				case http.MethodGet:
					r.Get(path, hlr.Func)
				case http.MethodPost:
					r.Post(path, hlr.Func)
				case http.MethodPut:
					r.Put(path, hlr.Func)
				case http.MethodDelete:
					r.Delete(path, hlr.Func)
				default:
					log.Println("Invalid method")
				}
			}
		}
	})
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	createRouterGroup(router)
	log.Println("Server started")
	http.ListenAndServe(fmt.Sprintf("%s:%d",
		config.Config().Host, config.Config().Port), router)
}
