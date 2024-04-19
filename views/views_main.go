package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {

	// 	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	// 	path := strings.TrimSuffix(r.URL.Path, "/")

	// 	uadmin.Trail(uadmin.DEBUG, "Path: %s", path)
	// 	// session := uadmin.IsAuthenticated(r)

	// 	switch path {
	// 	case "login":
	// 		LoginHandler(w, r)
	// 	case "dashboard":
	// 		DashboardHandler(w, r)
	// 	case "company_details":
	// 		CompanyHandler(w, r)
	// 	case "images":
	// 		ImagesHandler(w, r)
	// 	case "user_management":
	// 		UserManagementHandler(w, r)
	// 	case "logout":
	// 		LogoutHandler(w, r)
	// 	default:
	// 		w.Write([]byte("Not Found"))
	// 	}
	// }
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	page := strings.TrimPrefix(r.URL.Path, "/")
	session := uadmin.IsAuthenticated(r)

	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	}

	c := map[string]interface{}{}
	switch page {

	case "dashboard":
		c = DashboardHandler(w, r)
	case "company":
		c = CompanyHandler(w, r)
	case "images":
		c = ImagesHandler(w, r)
	case "user":
		c = UserManagementHandler(w, r)
	default:
		page = "dashboard"
		// w.Write([]byte("Not Found"))
	}
	c["Page"] = page

	Rendering(w, r, page, c)

}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/base.html")
	// newPage := strings.TrimSuffix(r.URL.Path, "/")

	path := "./templates/" + page + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}

func ElementHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	return c
}
