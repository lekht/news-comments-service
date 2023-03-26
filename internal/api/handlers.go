package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/lekht/news-comments-service/pkg/storage"
)

func (a *API) commentsByIDHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	idValue := q.Get("id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := a.db.CommentsByNewsID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("api - commentsByIDHandler - db data getting error: %s\n", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comments)

}

func (a *API) addCommentHandler(w http.ResponseWriter, r *http.Request) {

	var comment *storage.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		body, _ := io.ReadAll(r.Body)
		fmt.Println(string(body))
		http.Error(w, fmt.Sprintf("wrong JSON: %s\n", err), http.StatusBadRequest)
		return
	}
	comment.PubTime = time.Now().Unix()
	err = a.db.AddComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
