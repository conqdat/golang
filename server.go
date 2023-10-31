package main

import (
	"fmt"
	"myapp/db"
	log "myapp/log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("init package !!!")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	fmt.Println("Main !!!")

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "password",
		DbName:   "postgres",
	}

	sql.Connect()
	defer sql.Close()

	// var email string
	// err := sql.Db.GetConte	xt(
	// 	context.Background(),
	// 	&email,
	// 	"SELECT email FROM users WHERE email=$1", "acb@gmail.com")
	// if err != nil {
	// 	log.Error(err.Error())
	// }

	e := echo.New()

	e.Logger.Fatal(e.Start(":3000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
