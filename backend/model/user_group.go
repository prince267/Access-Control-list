package model

type UserGroup struct {
	UserID    int64  `json:"user_id"  column:"user_id"`
	GroupID   int64  `json:"group_id" column:"groups.group_id"`
	GroupName string `json:"group_name" column:"group_name"`
}

func (userGroup *UserGroup) UserGroupTable() string {
	return "user_group"
}

func (userGroup *UserGroup) GroupTable() string {
	return "groups"
}

func (userGroup *UserGroup) String() string {
	return Stringify(userGroup)
}
