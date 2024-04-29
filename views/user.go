package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// Handler is used to handle/load pages
func UserManagementHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	users := []uadmin.User{}

	uadmin.All(&users)
	for x := range users {
		uadmin.Preload(&users[x])
	}
	c["Users"] = users

	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "User Management"

	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/user_management")

	// uadmin.RenderHTML(w, r, "templates/user.html", c)
	
	return c
}
