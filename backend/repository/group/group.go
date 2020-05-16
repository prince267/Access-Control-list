package group

import (
	"context"
	"database/sql"

	"github.com/Access-Control-list/backend/driver"
	"github.com/Access-Control-list/backend/model"
)

type groupRepository struct {
	conn *sql.DB
}

func NewGroupRepository(conn *sql.DB) *groupRepository {
	return &groupRepository{conn: conn}
}

func (group *groupRepository) GetGroupUsers(cntx context.Context, id int64) ([]interface{}, error) {
	obj := new(model.GroupUsers)
	return driver.GetGroupUsers(group.conn, obj, id)
}
