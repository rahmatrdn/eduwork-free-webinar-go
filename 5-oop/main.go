package main

import "fmt"

type user struct {
	name string
}

func newUser(name string) user {
	return user{name: name}
}

func (u user) getName() string {
	return u.name
}

func main() {
	user := newUser("John Lennon")
	fmt.Println(user.getName()) // John Lennon
}
