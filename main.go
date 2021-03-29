package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Conn struct {
	ConnTest *sql.DB
}

func (db *Conn) Connect(user string, password string, dbname string, port int, host string) (err error) {
	defer func() { err = errors.Wrap(err, "Conn.Connect") }()
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "test",
	}
	// v := url.Values{}
	v, err := url.ParseQuery("sslmode=disable")
	if err != nil {
		err = errors.Wrap(err, "Failed ParseQuery")
		return
	}
	u.RawQuery = v.Encode()
	u.User = url.UserPassword(user, password)
	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%d", user, password, dbname, port)
	fmt.Println(u.String())
	db.ConnTest, err = sql.Open("postgres", u.String())
	if err != nil {
		err = errors.Wrap(err, "Failed connection to DB")
		return
	}
	return
}

func (db *Conn) Close() (err error) {
	defer func() { err = errors.Wrap(err, "Conn.Close") }()
	err = db.ConnTest.Close()
	if err != nil {
		err = errors.Wrap(err, "Failed close DB")
		return
	}
	return
}

func main() {
	var conTest Conn
	err := conTest.Connect("ivan_user", "password", "test", 5432, "localhost")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = conTest.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
