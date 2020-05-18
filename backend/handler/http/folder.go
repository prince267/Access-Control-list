package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Access-Control-list/backend/handler"
	"github.com/Access-Control-list/backend/model"
	"github.com/Access-Control-list/backend/repository"
	"github.com/Access-Control-list/backend/repository/folder"
	"github.com/go-chi/chi"
)

type Folder struct {
	handler.HTTPHandler
	repo repository.IFolder
}

func NewFolderHandler(conn *sql.DB) *Folder {
	return &Folder{
		repo: folder.NewFolderRepository(conn),
	}
}

func (folder *Folder) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{

		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "folders/{id}", Func: folder.GetUserFolder},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "allFolders", Func: folder.GetAllFolders},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "folder", Func: folder.CreateFolder},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userFolder", Func: folder.NewUserFolder},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "parentFolders/", Func: folder.GetParentFolders},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "folder", Func: folder.UpdateFolderInFolder},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "FolderInFolder/{id}", Func: folder.DeleteFolderInFolderById},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "folder/{id}", Func: folder.DeleteFolderById},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "folderUser/{id}", Func: folder.GetFolderUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "folderUser/{id}/{folderId}", Func: folder.CheckIsFolderUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "FolderInFolder/userID/{id}", Func: folder.DeleteFolderInFolderByUserId},
	}
}

func (folder *Folder) GetUserFolder(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = folder.repo.GetUserFolder(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (folder *Folder) GetParentFolders(w http.ResponseWriter, r *http.Request) {
	UserID, _ := strconv.ParseInt(r.URL.Query()["UserID"][0], 10, 64)
	FolderID, _ := strconv.ParseInt(r.URL.Query()["FolderID"][0], 10, 64)
	var usr interface{}
	var err error
	for {

		usr, err = folder.repo.GetParentFolders(r.Context(), UserID, FolderID)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (folder *Folder) NewUserFolder(w http.ResponseWriter, r *http.Request) {
	var usr model.NewFolderInFolder
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = folder.repo.NewUserFolder(r.Context(), usr)
		break
	}
	if err != nil {
		handler.WriteJSONResponse(w, r, usr, http.StatusBadRequest, err)
	} else {
		handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	}

}

func (folder *Folder) UpdateFolderInFolder(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	usr := model.NewFolderInFolder{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		// usr.UserID = id
		// if nil != err {
		// 	break
		// }

		// set logged in user id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = folder.repo.UpdateFolderInFolder(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.NewFolderInFolder)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
func (folder *Folder) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var payload string
	var usr model.Folders
	var err error
	var id interface{}
	err = json.NewDecoder(r.Body).Decode(&usr)
	FolderName := usr.FolderName
	PathName := usr.PathName
	for {
		id, err = folder.repo.CreateFolder(r.Context(), FolderName, PathName)
		if nil != err {
			break
		}
		payload = "Folder Created successfully"
		break
	}
	if payload == "" {
		payload = "not successfull"
		handler.WriteJSONResponse(w, r, payload, http.StatusBadRequest, err)
	} else {
		handler.WriteJSONResponse(w, r, id, http.StatusOK, err)
	}
}

func (folder *Folder) DeleteFolderInFolderById(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		_, err = folder.repo.DeleteFolderInFolderById(r.Context(), id)

		if nil != err {
			break
		}
		payload = "Folder deleted successfully from Folder In Folder"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (folder *Folder) DeleteFolderInFolderByUserId(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		_, err = folder.repo.DeleteFolderInFolderByUserId(r.Context(), id)

		if nil != err {
			break
		}
		payload = "Folder deleted successfully from Folder In Folder"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (folder *Folder) DeleteFolderById(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		_, err = folder.repo.DeleteFolderById(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Folder deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (folder *Folder) GetAllFolders(w http.ResponseWriter, r *http.Request) {
	usrs, err := folder.repo.GetAllFolders(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (folder *Folder) GetFolderUser(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = folder.repo.GetFolderUser(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (folder *Folder) CheckIsFolderUser(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	UserId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	folderId, err := strconv.ParseInt(chi.URLParam(r, "folderId"), 10, 64)
	log.Println(folderId, UserId)
	for {
		if nil != err {
			break
		}

		usr, err = folder.repo.CheckIsFolderUser(r.Context(), UserId, folderId)
		if usr == nil {
			usr = ""
		}
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
