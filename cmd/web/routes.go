package main

import (
	"gocourse/pkg/config"
	"gocourse/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(App *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	// mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/rooms/room1", http.HandlerFunc(handlers.Repo.Room1))

	mux.Get("/rooms/room2", http.HandlerFunc(handlers.Repo.Room2))
	mux.Get("/book", http.HandlerFunc(handlers.Repo.Book))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))

	mux.Get("/makeReserve", http.HandlerFunc(handlers.Repo.MakeReserve))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
