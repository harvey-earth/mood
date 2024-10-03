package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/harvey-earth/mood/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	teams    models.TeamModelInterface
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	dbType := flag.String("database", "sqlite3", "Database type")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// This holds the database connection string
	var dsn string

	if *dbType == "sqlite3" {
		dsn = "./mood.db"
	} else if *dbType == "mysql" {
		dsn = fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	} else if *dbType == "psql" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, 5432, dbUser, dbPass, dbName)
	} else {
		errorLog.Fatal("Only mysql, psql, and sqlite3 are supported database values.")
	}

	db, err := openDB(*dbType, dsn)
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
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dbType string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
