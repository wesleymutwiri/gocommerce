package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/wesleymutwiri/gocommerce/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *httprouter.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", DbUser, DbPassword, DbName, DbPort)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("Successfully connected to %s database", Dbdriver)
	}
	server.DB.Debug().AutoMigrate(&models.User{}) // database migration
	server.Router = httprouter.New()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
