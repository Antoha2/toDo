package etodo

type User struct {
	Id        int    `json:"-" db:"user_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
