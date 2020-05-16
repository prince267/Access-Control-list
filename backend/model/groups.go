package model

type Groups struct {
	GroupID   int64  `json:"group_id,omitempty" key:"primary" autoincr:"1" column:"group_id"`
	GroupName string `json:"group_name" column:"group_name"`
}

func (groups *Groups) GroupsTable() string {
	return "groups"
}

func (groups *Groups) String() string {
	return Stringify(groups)
}
