package post

import (
	"github.com/factly/dega-vito/model"
	"github.com/go-chi/chi"
)

type postData struct {
	model.Post
	Authors []model.Author `json:"authors"`
	Claims  []model.Claim  `json:"claims"`
}

// Router posts router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Get("/{post_id}", details)

	return r
}
