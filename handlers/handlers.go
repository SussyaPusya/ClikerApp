package handlers

import "fmt"

type Point struct {
	Score int
}

func (p *Point) AddPoint() {
	p.Score++
	fmt.Println("+1")
}

type User struct {
	Username     string
	password     string
	autheficated bool
}

func (u User) SetNameAndPass(username string, password string) {
	u.Username = username
	u.password = password
}

var Db = map[string]User{
	"bob": User{"bob", "nazik", false},
}
var IpBd = map[string]bool{}

func ValidAuth(ip string) bool {
	_, ok := IpBd[ip]
	if !ok {
		fmt.Println(false)
		return false
	}
	fmt.Println(true)
	return true
}
