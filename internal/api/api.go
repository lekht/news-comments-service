package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lekht/news-comments-service/config"
	"github.com/lekht/news-comments-service/pkg/storage"
)

type store interface {
	CommentsByNewsID(id int) (*[]storage.Comment, error)
	AddComment(c *storage.Comment) error
}

type API struct {
	r  *mux.Router
	db store
}

// Регистрация методов в маршрутизаторе
func (a *API) endpoints() {
	a.r.Name("comments").Path("/comment").Methods(http.MethodPost).HandlerFunc(a.commentsByIDHandler)
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
