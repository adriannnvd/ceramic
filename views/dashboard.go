package views

import (
	"net/http"
	"strings"
	"time"

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

	now := time.Now()
	year, month, _ := now.Date()
	// For the uploads column graph
	t := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	currentMonth := t.Month()
	c["CurrentMonth"] = currentMonth

	// uadmin.Trail(uadmin.DEBUG, "CurrentMonth", currentMonth)

	previousMonth := currentMonth - 1
	if currentMonth == time.January {
		previousMonth = time.December
	}
	c["PreviousMonth"] = previousMonth
	// uadmin.Trail(uadmin.DEBUG, "PreviousMonth", previousMonth)

	secPreviousMonth := currentMonth - 2
	if currentMonth == time.February {
		previousMonth = time.December
	}
	c["SecPreviousMonth"] = secPreviousMonth

	thirdPreviousMonth := currentMonth - 3
	if currentMonth == time.March {
		thirdPreviousMonth = time.December
	}
	c["ThirdPreviousMonth"] = thirdPreviousMonth

	// Get the number of images uploaded in the month requested

	goodTotalCurrentMonth := uadmin.Count(ceramic, "classification_num == 1 AND month_uploaded = ?", currentMonth.String())
	c["GoodTotalCurrentMonth"] = goodTotalCurrentMonth
	defectTotalCurrentMonth := uadmin.Count(ceramic, "classification_num == 2 AND month_uploaded = ?", currentMonth.String())
	c["DefectTotalCurrentMonth"] = defectTotalCurrentMonth

	goodTotalPrevMonth := uadmin.Count(ceramic, "classification_num == 1 AND month_uploaded = ?", previousMonth.String())
	c["GoodTotalPrevMonth"] = goodTotalPrevMonth
	defectTotalPrevMonth := uadmin.Count(ceramic, "classification_num == 2 AND month_uploaded = ?", previousMonth.String())
	c["DefectTotalPrevMonth"] = defectTotalPrevMonth

	goodTotalSecMonth := uadmin.Count(ceramic, "classification_num == 1 AND month_uploaded = ?", secPreviousMonth.String())
	c["GoodTotalSecMonth"] = goodTotalSecMonth
	defectTotalSecMonth := uadmin.Count(ceramic, "classification_num == 2 AND month_uploaded = ?", secPreviousMonth.String())
	c["DefectTotalSecMonth"] = defectTotalSecMonth

	goodTotalThirdMonth := uadmin.Count(ceramic, "classification_num == 1 AND month_uploaded = ?", thirdPreviousMonth.String())
	c["GoodTotalThirdMonth"] = goodTotalThirdMonth
	defectTotalThirdMonth := uadmin.Count(ceramic, "classification_num == 2 AND month_uploaded = ?", thirdPreviousMonth.String())
	c["DefectTotalThirdMonth"] = defectTotalThirdMonth

	// Initialize Title and mapped it on the html file (you can check it if you want :)
	c["Title"] = "Dashboard"

	// tells the handler that it will load to this path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dashboard")

	// uadmin.RenderHTML(w, r, "templates/dashboard.html", c)

	return c
}
