package views

import (
	"net/http"
	"strings"

	"github.com/adriannnvd/ceramic/models"
	"github.com/uadmin/uadmin"
	// "github.com/uadmin/uadmin"
)

// Handler is used to handle/load pages
func ImagesHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	ceramic := []models.Ceramic{}

	uadmin.All(&ceramic)
	for x := range ceramic {
		uadmin.Preload(&ceramic[x])
	}
	
	c["Ceramic"] = ceramic
	
	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "Images"

	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/images")

	// uadmin.RenderHTML(w, r, "templates/images.html", c)
	
	return c
}
