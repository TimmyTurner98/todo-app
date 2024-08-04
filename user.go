package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" blinding:"required"`
	Username string `json:"username" blinding:"required"`
	Password string `json:"password" blinding:"required"`
}
