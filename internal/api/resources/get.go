package resources

import (
	"github.com/go-chi/chi/v5"
	"github.com/ppal31/mygo/internal/api/render"
	"github.com/ppal31/mygo/internal/logger"
	"github.com/ppal31/mygo/internal/store/database"
	"net/http"
	"strconv"
)

// HandleGet returns an http.HandlerFunc that writes a json-encoded object to the response body.
func HandleGet(rs database.ResourceStoreInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, err := strconv.ParseInt(chi.URLParam(r, "resourceId"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.Debug(ctx, "cannot parse resource id", "err", err)
			return
		}
		resource, err := rs.Get(ctx, id)
		if err != nil {
			render.InternalError(w, err)
			logger.Error(ctx, "cannot retrieve list", "err", err)
		} else {
			render.JSON(w, resource, 200)
		}
	}
}
