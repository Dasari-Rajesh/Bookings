package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Cross-site request forgery
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	fmt.Println("MiddleWare Started")

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Indevelopment,
		SameSite: http.SameSiteLaxMode,
	})
	fmt.Println("MiddleWare ended")
	return csrfHandler
}
// func SessionLoad(next http.Handler) http.Handler {
// 	fmt.Println("Session Started")
// 	return se ssion.LoadAndSave(next)
// }
