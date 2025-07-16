package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"golang_course_demo/internal/api"
	"golang_course_demo/internal/config"
	"golang_course_demo/internal/db"
	"golang_course_demo/internal/migrations"
	"golang_course_demo/internal/reader"
	"golang_course_demo/internal/storage"
	"golang_course_demo/internal/worker"
	"log"
	"net/http"
)

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
	sourceRepo := storage.NewSource(pool)
	channelRepo := storage.NewChannel(pool)
	recordsRepo := storage.NewRecords(pool)

	w := worker.New(ctx, r, channelRepo, recordsRepo, sourceRepo)
	router := api.New(":8080", sourceRepo, channelRepo, recordsRepo)

	eg := &errgroup.Group{}
	eg.Go(router.Start)
	eg.Go(w.Scan)

	log.Println("> Scanning")

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
