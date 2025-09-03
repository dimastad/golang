package main

import (
	"errors"
	"fmt"
	// "golang/maps/mapuse"
	// "golang/maps/mymap"
)

type User struct {
	Name  string
	Email string
}

func UpdateEmail(users map[int]*User, id int, newEmail string) error {
	fmt.Println("1", users[id])
	_, ok := users[id]
	
	fmt.Println(ok)
	if ok {
		users[id].Email = newEmail
	} else {
		return errors.New("user not found")
	}
	return nil
}

func main() {
	// users := map[int]*User{
	// 	1: {Name: "Alice", Email: "alice@example.com"},
	// 	2: {Name: "Bob", Email: "bob@example.com"},
	// }

	// fmt.Println(UpdateEmail(users, 1, "alice@newmail.com"))
	// settings := map[string]map[string]string{
	// 	"alice": {
	// 		"theme": "dark",
	// 		"lang":  "en",
	// 	},
	// }

	// mymap.SetUserSet(settings, "alice", "lang", "ru")
	// mymap.SetUserSet(settings, "bob", "theme", "light")
	// users := map[string]string{
	// 	"Alice": "Go",
	// 	"Bob":   "Python",
	// 	"Tom":   "Go",
	// 	"Kate":  "Java",
	// }
	// mapuse.CountLanguages(users)
}