package handlers

import (
	"fmt"
	"gocourse/pkg/config"
	"gocourse/pkg/models"
	"gocourse/pkg/renders"
	"net/http"
)

var Repo *Repository
var stringData = map[string]string{}
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Running!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	remoteIP := r.RemoteAddr
	// fmt.Println(remoteIP)
	stringData["remote_ip"] = remoteIP
	
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// fmt.Println(remoteIP)

	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About Running!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	
	stringData["Test"] = "HELLO WOrLD"
	// remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	// stringData["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringData: stringData,
	})
}
