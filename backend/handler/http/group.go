package http

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Access-Control-list/backend/handler"
	"github.com/Access-Control-list/backend/repository"
	"github.com/Access-Control-list/backend/repository/group"
)

type Group struct {
	handler.HTTPHandler
	repo repository.IGroup
}

func NewGroupHandler(conn *sql.DB) *Group {
	return &Group{
		repo: group.NewGroupRepository(conn),
	}
}

func (group *Group) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groupUsers/{id}", Func: group.GetGroupUsers},
	}
}

func (group *Group) GetGroupUsers(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = group.repo.GetGroupUsers(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
