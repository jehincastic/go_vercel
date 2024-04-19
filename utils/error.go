package utils

import (
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, status int, message any) {
	env := envelope{"error": message}
	err := WriteJSON(w, status, env, nil)

	if err != nil {
		fmt.Printf("error happened while writing json %s\n", err.Error())
		w.WriteHeader(500)
	}

}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("something went wrong %s\n", err.Error())
	message := "the server encountered a problem and could not process your request"
	errorResponse(w, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	errorResponse(w, http.StatusUnprocessableEntity, errors)
}

func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	errorResponse(w, http.StatusUnauthorized, message)
}

func NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	errorResponse(w, http.StatusForbidden, message)
}
