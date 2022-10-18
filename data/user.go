package data

type User interface {
	createTableUser()
	createUser()
}

type user struct {
	Username string `json:username`
	Password string `json:password`
}
