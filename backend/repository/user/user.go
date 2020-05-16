package user

import (
	"context"
	"database/sql"

	"github.com/Access-Control-list/backend/driver"
	"github.com/Access-Control-list/backend/model"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{conn: conn}
}

func (user *userRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.User)
	obj1 := new(model.UserGroup)
	return driver.GetById(user.conn, obj, obj1, id)
}

func (user *userRepository) Login(cntx context.Context, id int64, password string) (interface{}, error) {
	obj := new(model.User)
	return driver.Login(user.conn, obj, id, password)
}

func (user *userRepository) GetUserGroup(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.UserGroup)
	return driver.GetUserGroup(user.conn, obj, id)
}

func (user *userRepository) GetGroupUsers(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.GroupUsers)
	return driver.GetGroupUsers(user.conn, obj, id)
}

func (user *userRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.User)
	result, err := driver.Create(user.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.UserID = id
	return id, nil
}

func (user *userRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.User)
	err := driver.UpdateById(user.conn, &usr)
	return obj, err
}

func (user *userRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.User{UserID: id}
	// return driver.SoftDeleteById(user.conn, obj, id)
	return driver.DeleteById(user.conn, obj, id)
}

func (user *userRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.User{}
	return driver.GetAll(user.conn, obj, 0, 0)
}
