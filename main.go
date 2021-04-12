package main

import (
	"fmt"
	"net/url"
)

func main() {
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	u := url.URL{}
	u.Scheme = "postgres"
	u.Host = "localhost:5433"
	u.User = url.UserPassword("user_ivan", "1234")
	u.Path = "db3"
	u.RawQuery = "sslmode=disable"
	fmt.Println(u.String())
}

