package model

type GroupUsers struct {
	UserID    int64  `json:"user_id"  column:"users.user_id"`
	FirstName string `json:"first_name" column:"first_name"`
	LastName  string `json:"last_name" column:"last_name"`
}

func (groupUsers *GroupUsers) UserGroupTable() string {
	return "user_group"
}

func (groupUsers *GroupUsers) UserTable() string {
	return "users"
}

func (groupUsers *GroupUsers) String() string {
	return Stringify(groupUsers)
}
