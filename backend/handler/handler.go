package handler

import (
	"encoding/json"
	"net/http"
)

type IHTTPHandler interface {
	GetHTTPHandler() []*HTTPHandler
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	GetUserGroup(http.ResponseWriter, *http.Request)
	GetGroupUsers(http.ResponseWriter, *http.Request)
	GetUserFiles(http.ResponseWriter, *http.Request)
	NewUserFolder(http.ResponseWriter, *http.Request)
	NewUserFile(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	CreateFolder(http.ResponseWriter, *http.Request)
	ReadFile(http.ResponseWriter, *http.Request)
	WriteFile(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	UpdateFolderInFolder(http.ResponseWriter, *http.Request)
	UpdateFileInFolder(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetParentFolders(http.ResponseWriter, *http.Request)
	GetParentFiles(http.ResponseWriter, *http.Request)
	DeleteFolder(http.ResponseWriter, *http.Request)
	DeleteFile(http.ResponseWriter, *http.Request)
	GetFileUser(http.ResponseWriter, *http.Request)
	GetFolderUser(http.ResponseWriter, *http.Request)
	CheckIsFileUser(http.ResponseWriter, *http.Request)
	CheckIsFolderUser(http.ResponseWriter, *http.Request)
}

type HTTPHandler struct {
	Authenticated bool
	Method        string
	Path          string
	Func          func(http.ResponseWriter, *http.Request)
}

type response struct {
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (hdlr *HTTPHandler) GetHTTPHandler() []HTTPHandler {
	return []HTTPHandler{}
}

func (hdlr *HTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetUserGroup(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetGroupUsers(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetParentFiles(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetParentFolders(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) ReadFile(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) WriteFile(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) NewUserFile(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) NewUserFolder(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) UpdateFolderInFolder(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) UpdateFileInFolder(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetAllFolders(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	return
}
func (hdlr *HTTPHandler) GetFileUser(http.ResponseWriter, *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetFolderUser(http.ResponseWriter, *http.Request) {
	return
}

func (hdlr *HTTPHandler) CheckIsFileUser(http.ResponseWriter, *http.Request) {
	return
}

func (hdlr *HTTPHandler) CheckIsFolderUser(http.ResponseWriter, *http.Request) {
	return
}

func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {
	resp := &response{
		Status: code,
		Data:   payload,
	}

	if nil != err {
		resp.Message = err.Error()
	}

	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}
