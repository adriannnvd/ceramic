package views

import (
	"net/http"
	"strings"

	// "github.com/uadmin/uadmin"
)

// Handler is used to handle/load pages
func CompanyHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "Company Details | Page"

	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/company_details")

	// uadmin.RenderHTML(w, r, "templates/company.html", c)
	
	return c
}
