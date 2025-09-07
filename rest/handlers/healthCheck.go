package rest

import (
	"net/http"

	"github.com/mir-asif01/students_api/utils"
)

func Health(w http.ResponseWriter, r* http.Request){
	utils.SendResponse(w,"API running successfully",200)
}