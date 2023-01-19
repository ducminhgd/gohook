package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleGitlabHook(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	hookID := chi.URLParam(r, "hookID")
	fmt.Println(hookID)
}
