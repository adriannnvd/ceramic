package views

import (
	"net/http"
	"strings"

	"github.com/adriannnvd/ceramic/models"
	"github.com/uadmin/uadmin"
)

// Handler is used to handle/load pages
func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	users := []uadmin.User{}
	ceramic := []models.Ceramic{}

	uadmin.All(&users)
	for x := range users {
		uadmin.Preload(&users[x])
	}
	c["Users"] = users

	uadmin.All(&ceramic)
	uadmin.AdminPage("id", false, -1, 1, &ceramic, "")	
	for x := range ceramic {
		uadmin.Preload(&ceramic[x])
	}
	c["Ceramic"] = ceramic

	total := uadmin.Count(ceramic, "id > 0")
	c["Total"] = total

	totalGood := uadmin.Count(ceramic, "classification_num == 1")
	c["TotalGood"] = totalGood
	totalDefect := uadmin.Count(ceramic, "classification_num == 2")
	c["TotalDefect"] = totalDefect


	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "Dashboard"

	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dashboard")

	// uadmin.RenderHTML(w, r, "templates/dashboard.html", c)

	return c
}
