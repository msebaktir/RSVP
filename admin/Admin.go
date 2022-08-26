package admin

import (
	"net/http"

	"msebaktir.com/LCV/utils"
	views "msebaktir.com/LCV/views"
)

func AdminPage(w http.ResponseWriter, r *http.Request) {
	// get current dir

	tmpl := views.BuildThemplate("Admin.template", "Admin/index.html")

	tmp, err := tmpl.Builder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := utils.NewPageData("Admin")

	tmp.ExecuteTemplate(w, "base", data)

}
