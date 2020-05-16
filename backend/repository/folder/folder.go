package folder

import (
	"context"
	"database/sql"

	"github.com/Access-Control-list/backend/driver"
	"github.com/Access-Control-list/backend/model"
)

type folderRepository struct {
	conn *sql.DB
}

func NewFolderRepository(conn *sql.DB) *folderRepository {
	return &folderRepository{conn: conn}
}

func (folder *folderRepository) GetUserFolder(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.FolderInFolder)
	return driver.GetUserFolder(folder.conn, obj, id)
}

func (folder *folderRepository) GetAllFolders(cntx context.Context) ([]interface{}, error) {
	obj := &model.AllFolders{}
	return driver.GetAllFolders(folder.conn, obj)
}

func (folder *folderRepository) CreateFolder(cntx context.Context, FolderName string, PathName string) (interface{}, error) {
	obj := new(model.Folders)
	result, err := driver.CreateFolder(folder.conn, obj, FolderName, PathName)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	obj.FolderID = id
	return id, err
}
func (folder *folderRepository) NewUserFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.NewFolderInFolder)
	result, err := driver.NewUserFolder(folder.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.RowsAffected()
	// usr.UserID = id
	return id, nil
}

func (folder *folderRepository) GetParentFolders(cntx context.Context, UserID int64, FolderID int64) ([]interface{}, error) {
	obj := new(model.NewFolderInFolder)
	return driver.GetParentFolders(folder.conn, obj, UserID, FolderID)
}

func (folder *folderRepository) DeleteFolderInFolderById(cntx context.Context, id int64) (sql.Result, error) {
	obj := &model.NewFolderInFolder{ChildFolderID: id}
	// return driver.SoftDeleteById(folder.conn, obj, id)
	return driver.DeleteFolderInFolderById(folder.conn, obj, id)
}

func (folder *folderRepository) DeleteFolderById(cntx context.Context, id int64) (sql.Result, error) {
	obj := &model.Folders{FolderID: id}
	// return driver.SoftDeleteById(folder.conn, obj, id)
	return driver.DeleteFolderById(folder.conn, obj, id)
}

func (folder *folderRepository) UpdateFolderInFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.NewFolderInFolder)
	_, err := driver.UpdateFolderInFolder(folder.conn, &usr)
	return obj, err
}

func (folder *folderRepository) GetFolderUser(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.NewFolderInFolder)
	return driver.GetFolderUser(folder.conn, obj, id)
}

func (folder *folderRepository) CheckIsFolderUser(cntx context.Context, userId int64, folderId int64) (interface{}, error) {
	obj := new(model.NewFolderInFolder)
	return driver.CheckIsFolderUser(folder.conn, obj, userId, folderId)
}
