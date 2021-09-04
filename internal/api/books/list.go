package books

import (
	"github.com/ppal31/mygo/internal/api/render"
	"github.com/ppal31/mygo/internal/logger"
	"net/http"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of objects to the response body.
func HandleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var err error
		books := []string{"mygo"}
		if err != nil {
			render.InternalError(w, err)
			logger.Error(ctx, "cannot retrieve list", "err", err)
		} else {
			render.JSON(w, books, 200)
		}
	}
}
