package repository

import (
	"context"
	"database/sql"
)

type IFile interface {
	GetUserFiles(context.Context, int64) ([]interface{}, error)
	CreateFile(context.Context, string, string) (interface{}, error)
	ReadFile(context.Context, string) (string, error)
	WriteFile(context.Context, string, string) error
	NewUserFile(context.Context, interface{}) (interface{}, error)
	UpdateFileInFolder(context.Context, interface{}) (interface{}, error)
	GetAllFiles(context.Context) ([]interface{}, error)
	GetParentFiles(context.Context, int64, int64) ([]interface{}, error)
	DeleteFileInFolderById(context.Context, int64) (sql.Result, error)
	DeleteFileById(context.Context, int64) (sql.Result, error)
	GetFileUser(context.Context, int64) ([]interface{}, error)
	CheckIsFileUser(context.Context, int64, int64) (interface{}, error)
	DeleteFileInFolderByUserId(context.Context, int64) (sql.Result, error)
}

type File struct {
}

func (file *File) GetUserFiles(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (file *File) CreateFile(cntx context.Context, obj interface{}, FileName string, PathName string) (fobj interface{}, err error) {
	return
}

func (file *File) ReadFile(cntx context.Context, PathName string) (data string, err error) {
	return
}

func (file *File) WriteFile(cntx context.Context, PathName string, Content string) (err error) {
	return
}

func (file *File) NewUserFile(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (file *File) UpdateFileInFolder(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (file *File) GetAllFiles(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (file *File) GetParentFiles(cntx context.Context, UserID int64, FolderID int64) (obj []interface{}, err error) {
	return
}

func (file *File) DeleteFileInFolderById(cntx context.Context, obj interface{}, id int64) (fobj sql.Result, err error) {
	return
}

func (file *File) DeleteFilesById(cntx context.Context, obj interface{}, id int64) (fobj sql.Result, err error) {
	return
}

func (file *File) GetFileUser(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (folder *Folder) CheckIsFileUser(cntx context.Context, userId int64, fileId int64) (cobj interface{}, err error) {
	return
}

func (file *File) DeleteFileInFolderByUserId(cntx context.Context, obj interface{}, id int64) (fobj sql.Result, err error) {
	return
}
