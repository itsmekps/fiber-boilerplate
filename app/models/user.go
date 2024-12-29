package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `jsom:"password"`
}

// Error implements error.
func (u *User) Error() string {
	panic("unimplemented")
}
