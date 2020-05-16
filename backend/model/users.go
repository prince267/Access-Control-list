package model

type User struct {
	UserID    int64  `json:"user_id,omitempty" key:"primary" autoincr:"1" column:"user_id"`
	FirstName string `json:"first_name" column:"first_name"`
	LastName  string `json:"last_name" column:"last_name"`
	Password  string `json:"password" column:"password"`
}

func (user *User) UserTable() string {
	return "users"
}

func (user *User) String() string {
	// log.Println("^^^^^^^^^", Stringify(user))
	return Stringify(user)
}
