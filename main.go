package main

import (
	"net/http"

	"github.com/adriannnvd/ceramic/api"
	"github.com/adriannnvd/ceramic/models"
	"github.com/adriannnvd/ceramic/views"

	"github.com/uadmin/uadmin"
)

func main() {

	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Ceramics Classification"

	uadmin.Register(
		models.Ceramic{},
	)
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))
	http.HandleFunc("/api/images/", uadmin.Handler(api.ImagesAPIHandler))
	http.HandleFunc("/api/addimage/", uadmin.Handler(api.AddImageAPIHandler))

	uadmin.StartServer()
}
