package handler

import (
	"net/http"

	"github.com/jehincastic/go_vercel/utils"
)

func Json(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	resp["message"] = "Hello World from Go! 👋"
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	resp["github"] = "https://github.com/riccardogiorato/template-go-vercel/blob/main/api/json.go"
	err := utils.WriteJSON(w, http.StatusCreated, resp, nil)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
	}
}
