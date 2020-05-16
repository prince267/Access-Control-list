package model

type Permission struct {
	PermissionID int64  `json:"permission_id,omitempty" key:"primary" autoincr:"1" column:"permission_id"`
	Description  string `json:"descrp" column:"descrp"`
}

func (permission *Permission) PermissionTable() string {
	return "permission"
}

func (permission *Permission) String() string {
	return Stringify(permission)
}
