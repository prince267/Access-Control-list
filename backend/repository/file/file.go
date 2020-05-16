package file

import (
	"context"
	"database/sql"

	"github.com/Access-Control-list/backend/driver"
	"github.com/Access-Control-list/backend/model"
)

type fileRepository struct {
	conn *sql.DB
}

func NewFileRepository(conn *sql.DB) *fileRepository {
	return &fileRepository{conn: conn}
}

func (file *fileRepository) GetUserFiles(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.FileInFolder)
	return driver.GetUserFiles(file.conn, obj, id)
}

func (file *fileRepository) CreateFile(cntx context.Context, FileName string, PathName string) (interface{}, error) {
	obj := new(model.Files)
	result, err := driver.CreateFile(file.conn, obj, FileName, PathName)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	obj.FileID = id
	return id, err
}

func (file *fileRepository) ReadFile(cntx context.Context, PathName string) (string, error) {
	data, err := driver.ReadFile(PathName)
	if nil != err {
		return "", err
	}

	return data, err
}

func (file *fileRepository) WriteFile(cntx context.Context, PathName string, Content string) error {
	err := driver.WriteFile(PathName, Content)
	return err
}

func (file *fileRepository) NewUserFile(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.NewFileInFolder)
	result, err := driver.NewUserFile(file.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.RowsAffected()
	// usr.UserID = id
	return id, nil
}

func (file *fileRepository) UpdateFileInFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.NewFileInFolder)
	_, err := driver.UpdateFileInFolder(file.conn, &usr)
	return obj, err
}

func (file *fileRepository) GetAllFiles(cntx context.Context) ([]interface{}, error) {
	obj := &model.AllFiles{}
	return driver.GetAllFiles(file.conn, obj)
}

func (file *fileRepository) GetParentFiles(cntx context.Context, UserID int64, FolderID int64) ([]interface{}, error) {
	obj := new(model.NewFileInFolder)
	return driver.GetParentFiles(file.conn, obj, UserID, FolderID)
}

func (file *fileRepository) DeleteFileInFolderById(cntx context.Context, id int64) (sql.Result, error) {
	obj := &model.NewFileInFolder{ChildFileID: id}
	// return driver.SoftDeleteById(file.conn, obj, id)
	return driver.DeleteFileInFolderById(file.conn, obj, id)
}

func (file *fileRepository) DeleteFileById(cntx context.Context, id int64) (sql.Result, error) {
	obj := &model.Files{FileID: id}
	// return driver.SoftDeleteById(file.conn, obj, id)
	return driver.DeleteFileById(file.conn, obj, id)
}

func (file *fileRepository) GetFileUser(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.NewFileInFolder)
	return driver.GetFileUser(file.conn, obj, id)
}

func (file *fileRepository) CheckIsFileUser(cntx context.Context, userId int64, fileId int64) (interface{}, error) {
	obj := new(model.NewFileInFolder)
	return driver.CheckIsFileUser(file.conn, obj, userId, fileId)
}
