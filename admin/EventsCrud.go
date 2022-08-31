package admin

import (
	"net/http"

	utils "msebaktir.com/LCV/utils"
	views "msebaktir.com/LCV/views"
)

func EventsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db, _ := utils.Connect()
	tmpl := views.BuildThemplate("Admin.template", "Admin/Events.html")
	tmp, _ := tmpl.Builder()
	allEvents, _ := utils.GetAllEvents(db)
	tmp.ExecuteTemplate(w, "base", allEvents)
}
