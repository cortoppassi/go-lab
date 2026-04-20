package httpapi

import "net/http"

func NewRouter(registerRoutes func(*http.ServeMux)) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	registerRoutes(mux)

	return mux
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		WriteError(w, http.StatusNotFound, "rota nao encontrada")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{
		"message": "API de tarefas funcionando",
		"routes": []string{
			"GET /tasks",
			"GET /tasks/{id}",
			"POST /tasks",
			"PUT /tasks/{id}",
			"DELETE /tasks/{id}",
		},
	})
}
