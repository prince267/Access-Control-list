package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/Access-Control-list/backend/handler"
	"github.com/Access-Control-list/backend/model"
	"github.com/Access-Control-list/backend/repository"
	"github.com/Access-Control-list/backend/repository/file"
)

type File struct {
	handler.HTTPHandler
	repo repository.IFile
}

func NewFileHandler(conn *sql.DB) *File {
	return &File{
		repo: file.NewFileRepository(conn),
	}
}

func (file *File) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{

		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "allFiles", Func: file.GetAllFiles},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userFile", Func: file.NewUserFile},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "file", Func: file.CreateFile},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "rfile", Func: file.ReadFile},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "wfile", Func: file.WriteFile},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "files/{id}", Func: file.GetUserFiles},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "parentFiles/", Func: file.GetParentFiles},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "file", Func: file.UpdateFileInFolder},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "FileInFolder/{id}", Func: file.DeleteFileInFolderById},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "file/{id}", Func: file.DeleteFileById},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "fileUser/{id}", Func: file.GetFileUser},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "fileUser/{id}/{fileId}", Func: file.CheckIsFileUser},
	}
}
func (file *File) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("get all files called")
	usrs, err := file.repo.GetAllFiles(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
func (file *File) NewUserFile(w http.ResponseWriter, r *http.Request) {
	var usr model.NewFileInFolder
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = file.repo.NewUserFile(r.Context(), usr)
		break
	}
	if err != nil {
		handler.WriteJSONResponse(w, r, usr, http.StatusBadRequest, err)
	} else {
		handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	}
}

func (file *File) CreateFile(w http.ResponseWriter, r *http.Request) {
	var payload string
	var usr model.Files
	var err error
	var id interface{}
	err = json.NewDecoder(r.Body).Decode(&usr)
	FileName := usr.FileName
	PathName := usr.PathName
	for {
		id, err = file.repo.CreateFile(r.Context(), FileName, PathName)
		if nil != err {
			break
		}
		payload = "File created successfully"
		break
	}
	if payload == "" {
		payload = "not successfull"
		handler.WriteJSONResponse(w, r, payload, http.StatusBadRequest, err)
	} else {
		handler.WriteJSONResponse(w, r, id, http.StatusOK, err)
	}
}

func (file *File) ReadFile(w http.ResponseWriter, r *http.Request) {
	var file1 model.ReadAndWriteFile
	var err error
	err = json.NewDecoder(r.Body).Decode(&file1)
	PathName := file1.PathName
	// PathName := r.URL.Query()["path"]
	data, err := file.repo.ReadFile(r.Context(), PathName)
	if nil == err {
		handler.WriteJSONResponse(w, r, data, http.StatusOK, err)
	} else {
		handler.WriteJSONResponse(w, r, data, http.StatusBadRequest, err)
	}

}

func (file *File) WriteFile(w http.ResponseWriter, r *http.Request) {
	var file1 model.ReadAndWriteFile
	var err error
	err = json.NewDecoder(r.Body).Decode(&file1)
	// PathName := r.URL.Query()["path"]
	// Content := r.URL.Query()["content"]
	PathName := file1.PathName
	Content := file1.Content
	var payload string
	err = file.repo.WriteFile(r.Context(), PathName, Content)
	if nil == err {
		payload = "File written"
		handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
	} else {
		payload = "Unable File written"
		handler.WriteJSONResponse(w, r, payload, http.StatusBadRequest, err)
	}

}

func (file *File) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = file.repo.GetUserFiles(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (file *File) GetParentFiles(w http.ResponseWriter, r *http.Request) {
	UserID, _ := strconv.ParseInt(r.URL.Query()["UserID"][0], 10, 64)
	FolderID, _ := strconv.ParseInt(r.URL.Query()["FolderID"][0], 10, 64)
	var usr interface{}
	var err error
	for {
		usr, err = file.repo.GetParentFiles(r.Context(), UserID, FolderID)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (file *File) UpdateFileInFolder(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	usr := model.NewFileInFolder{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		iUsr, err = file.repo.UpdateFileInFolder(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.NewFileInFolder)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (file *File) DeleteFileInFolderById(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		_, err = file.repo.DeleteFileInFolderById(r.Context(), id)

		if nil != err {
			break
		}
		payload = "Folder deleted successfully from File In Folder"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (file *File) DeleteFileById(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		_, err = file.repo.DeleteFileById(r.Context(), id)

		if nil != err {
			break
		}
		payload = "File deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (file *File) GetFileUser(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = file.repo.GetFileUser(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (file *File) CheckIsFileUser(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	UserId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	fileId, err := strconv.ParseInt(chi.URLParam(r, "fileId"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = file.repo.CheckIsFileUser(r.Context(), UserId, fileId)
		if usr == nil {
			usr = ""
		}
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
