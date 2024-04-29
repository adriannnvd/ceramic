package api

import (
	"net/http"
	
	"github.com/adriannnvd/ceramic/models"
	"github.com/uadmin/uadmin"
)

func ImagesAPIHandler(w http.ResponseWriter, r *http.Request) {
	ceramic := []models.Ceramic{}

	results := []map[string]interface{}{}

	uadmin.AdminPage("id", false, 0, -1, &ceramic, "")

	for i := range ceramic {
		uadmin.Preload(&ceramic[i])

		results = append(results, map[string]interface{}{
			"ID":              ceramic[i].ID,
			"Name":            ceramic[i].Name,
			"Images":          ceramic[i].Image,
			"ClassificationNum": ceramic[i].ClassificationNum,
			"Classification":  ceramic[i].Classification,
			"Date":            ceramic[i].Date,
		})
	}

	// uadmin.Trail(uadmin.DEBUG, "results:" ,results)

	uadmin.ReturnJSON(w, r, results)
}