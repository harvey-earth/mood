package main

import (
	"database/sql"
	"flag"
	"image/color"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/harvey-earth/mood/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	teams    *models.TeamModel
}

var palette1 = []color.Color{color.RGBA{0, 0xff, 0, 0xff}, color.Black}
var palette2 = []color.Color{color.RGBA{0xff, 0, 0, 0xff}, color.Black}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbString := dbUser + ":" + dbPass + "@/mood?parseTime=true"
	dsn := flag.String("dsn", dbString, "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		teams:    &models.TeamModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
