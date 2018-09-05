package main

import (
	"fmt"
	"study/6_interface/notification"
)

type user struct {
	email string
	name  string
}

func (u *user) Notify() {
	fmt.Printf("%s notified.", u.email)
}

// const adminLevel := [...]string{"super", "administrator", "moderator"}

type admin struct {
	user
	level string
}

func (a *admin) Notify() {
	fmt.Printf("%s notified. (admin)", a.email+", "+a.user.email)
}

func main() {
	ad := admin{
		level: "level1",
		user: user{
			name:  "john",
			email: "john@gmail.com",
		},
	}

	fmt.Println(ad, ad.user.email, ad.email)
	ad.Notify() // ad.user.notify -> ad.notify promoted
	notification.NotifyExternalFunc(&ad)
}
