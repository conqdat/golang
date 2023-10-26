package db

import (
	"fmt"

	"github.com/labstack/gommon/log"
)

type Sql struct {
	// Db *sqlx.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	dataSource := fmt.Printf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.UserName, s.Password, s.DbName)
	s.Db = sqlx.MustConnect("postgress", dataSource)

	if err := s.Db.ping(); err != nil {
		log.Error(err.Error())
	}
}

func (s *Sql) Close() {
	s.Db.Close()
}
