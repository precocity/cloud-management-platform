package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo"

	"cloud-management-platform/pkg/resources/request"
)

func main() {

	//migrate database
	m, err := migrate.New(
		"file:///migrations",
		"postgres://request_svc:request_svc@requests.db:5432/request_svc?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	repo := request.CreateRepository()

	service := request.CreateService(repo)

	handler := request.CreateHandler(service)

	e := echo.New()
	e.GET("/", handler.HandleCreateRequest)

	e.Logger.Fatal(e.Start(":3000"))
}
