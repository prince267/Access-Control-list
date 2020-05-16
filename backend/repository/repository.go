package repository

import (
	"context"
)

type IUser interface {
	GetByID(context.Context, int64) (interface{}, error)
	GetUserGroup(context.Context, int64) ([]interface{}, error)
	Login(context.Context, int64, string) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]interface{}, error)
}

type User struct {
}

func (user *User) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

func (user *User) Login(cntx context.Context, id int64, password string) (obj interface{}, err error) {
	return
}

func (user *User) GetUserGroup(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}

func (user *User) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (user *User) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (user *User) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

func (user *User) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}
