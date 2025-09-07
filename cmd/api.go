package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mir-asif01/students_api/config"
	"github.com/mir-asif01/students_api/rest"
)

func Api(){
	fmt.Println("API server running")
	
	cnf := config.GetEnv()
	addr :=":" + strconv.Itoa(cnf.Port)

	mux := http.NewServeMux()
	rest.Router(mux)
	fmt.Printf("Server running on http://localhost:%d\n\n",cnf.Port)

	err := http.ListenAndServe(addr,mux)
	if err != nil {
		log.Fatal("Error running the Server")
	}
}