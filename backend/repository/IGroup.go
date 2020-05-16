package repository

import (
	"context"
)

type IGroup interface {
	GetGroupUsers(context.Context, int64) ([]interface{}, error)
}
type Group struct {
}

func (group *Group) GetGroupUsers(cntx context.Context, id int64) (obj []interface{}, err error) {
	return
}
