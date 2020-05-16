package repository

import (
	"context"
	"database/sql"
)

type IFolder interface {
	GetUserFolder(context.Context, int64) ([]interface{}, error)
	GetAllFolders(context.Context) ([]interface{}, error)
	CreateFolder(context.Context, string, string) (interface{}, error)
	NewUserFolder(context.Context, interface{}) (interface{}, error)
	GetParentFolders(context.Context, int64, int64) ([]interface{}, error)
	DeleteFolderInFolderById(context.Context, int64) (sql.Result, error)
	DeleteFolderById(context.Context, int64) (sql.Result, error)
	UpdateFolderInFolder(context.Context, interface{}) (interface{}, error)
	GetFolderUser(context.Context, int64) ([]interface{}, error)
	CheckIsFolderUser(context.Context, int64, int64) (interface{}, error)
}

type Folder struct {
}

func (folder *Folder) GetUserFolder(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (folder *Folder) CreateFolder(cntx context.Context, obj interface{}, FolderName string, PathName string) (fobj interface{}, err error) {
	return
}

func (folder *Folder) NewUserFolder(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (folder *Folder) UpdateFolderInFolder(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (folder *Folder) GetAllFolders(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (folder *Folder) GetParentFolders(cntx context.Context, UserID int64, FolderID int64) (obj []interface{}, err error) {
	return
}

func (folder *Folder) DeleteFolderInFolderById(cntx context.Context, obj interface{}, id int64) (fobj sql.Result, err error) {
	return
}

func (folder *Folder) DeleteFolderById(cntx context.Context, obj interface{}, id int64) (fobj sql.Result, err error) {
	return
}

func (folder *Folder) GetFolderUser(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (folder *Folder) CheckIsFolderUser(cntx context.Context, userId int64, folderId int64) (cobj interface{}, err error) {
	return
}
