package api

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/adriannnvd/ceramic/models"
	"github.com/uadmin/uadmin"
)

func AddImageAPIHandler(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		f, fh, err := r.FormFile("photo") //Name of the input
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Form file error, %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		if fh.Size <= 0 {
			w.WriteHeader(400)
			w.Write([]byte("Got file length <= 0"))
			return
		}

		outFile, err := getFile(fh.Filename)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Get file error %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		written, err := io.Copy(outFile, f)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Copy error. %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		uadmin.Trail(uadmin.INFO, "Written. %#v", written)

		ceramic := models.Ceramic{}

		// Create a path in media folders for the directory
		ceramic.Image = "/media/uploads/" + fh.Filename
		ceramic.Save()

		uadmin.Save(&ceramic)

		// Redirect so  it will not stay on api/images
		http.Redirect(w, r, "/images", http.StatusSeeOther)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getFile(filename string) (*os.File, error) {
	file, err := os.Create("media/uploads/" + filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
