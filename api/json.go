package handler

import (
	"net/http"

	"github.com/jehincastic/go_vercel/utils"
)

func Json(w http.ResponseWriter, r *http.Request) {
	resp := make(utils.Envelope)
	resp["message"] = "Hello World from Go! 👋"
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	resp["github"] = "https://github.com/jehincastic"
	err := utils.SuccessRespone(w, http.StatusCreated, resp, nil)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
	}
}
