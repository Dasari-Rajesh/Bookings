package main

import (
	// "errors"
	"fmt"
	"gocourse/pkg/config"
	"gocourse/pkg/handlers"
	"gocourse/pkg/renders"
	"time"

	// "gocourse/pkg/route"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// Sprintf concat any two different types

func main() {

	tc, err := renders.CreateCacheTemplate()

	app.Indevelopment = false

	//session time, persist(if browser close cookie ends)
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Indevelopment
	app.Session = session

	if err != nil {
		log.Fatal("Cannot create a Template", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	renders.NewTemplate(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("This is running port %s\n", portNum)
	// _ = http.ListenAndServe(portNum, nil)
	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
