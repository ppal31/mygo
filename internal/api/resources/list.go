package resources

import (
	"github.com/ppal31/mygo/internal/api/render"
	"github.com/ppal31/mygo/internal/logger"
	"github.com/ppal31/mygo/internal/store/database"
	"github.com/ppal31/mygo/internal/types"
	"net/http"
	"strconv"
)

// HandleList returns an http.HandlerFunc that writes a json-encoded
// list of objects to the response body.
func HandleList(rs database.ResourceStoreInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resources, err := rs.List(ctx, buildPageParams(r))
		if err != nil {
			render.InternalError(w, err)
			logger.Error(ctx, "cannot retrieve list", "err", err)
		} else {
			render.JSON(w, resources, 200)
		}
	}
}

//TODO: Add validation etc. Probably a better thing to do is to build this into a chi-middleware
func buildPageParams(r *http.Request) types.Params {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size == 0 {
		size = 10
	}

	return types.Params{
		Page: page,
		Size: size,
	}
}
