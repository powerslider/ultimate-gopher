package foo

type User struct {
	FirstName string
	LastName  string
	password  string
}

func NewUser(firstName string, lastName string, password string) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		password:  password,
	}
}
