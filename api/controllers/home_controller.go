package controllers

import (
	"net/http"

	"github.com/MAlifiahM/Vutura_Test_Code/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}