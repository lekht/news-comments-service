package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lekht/news-comments-service/config"
)

type API struct {
	r  *mux.Router
	db store
}

// Регистрация методов в маршрутизаторе
func (a *API) endpoints() {
	a.r.Use(a.accessMiddleware, a.requestIdMiddlware, a.logRequestMiddlware)
	a.r.Name("comments").Path("/comment").Methods(http.MethodGet).HandlerFunc(a.commentsByIDHandler)
	a.r.Name("add_comment").Path("/comment").Methods(http.MethodPost).HandlerFunc(a.addCommentHandler)
}

func (a *API) Router() *mux.Router {
	return a.r
}

// Конструктор API
func New(cfg *config.Server, db store) *API {
	a := API{
		db: db,
		r:  mux.NewRouter(),
	}
	a.endpoints()
	return &a
}
