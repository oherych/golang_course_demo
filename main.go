package main

import (
	"context"
	"golang_course_demo/internal/config"
	"golang_course_demo/internal/db"
	"golang_course_demo/internal/migrations"
	"golang_course_demo/internal/reader"
	"golang_course_demo/internal/storage"
	"golang_course_demo/internal/worker"
	"log"
	"net/http"
)

var urls = []string{
	"http://podcast.dou.ua/rss",
}

func main() {
	// TODO: read from env
	cfg := config.Config{
		Database: "postgres://postgres:password@localhost:5432/postgres?sslmode=disable",
	}

	ctx := context.Background()

	log.Println("> Start migration")

	if err := migrations.Run(cfg.Database); err != nil {
		log.Fatal(err)
	}

	log.Println("> Connect to DB")

	pool, err := db.Connect(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	r := reader.New(http.Client{})
	feedsRepo := storage.NewFeeds(pool)

	w := worker.New(r, feedsRepo)

	log.Println("> Scanning")

	if err := w.Scan(ctx, urls); err != nil {
		log.Fatal(err)
	}

}
