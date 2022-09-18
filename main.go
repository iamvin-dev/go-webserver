package main;

import (
	"fmt"
	"log"
	"net/http"
)

func whoisvinHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/whoisvin" {
		fmt.Fprintf(res, "Vin is a 15 year old fullstack web developer from Germany. https://iamvin.dev | https://vinlikes.tech | contact@vinlikes.tech")
	} else {
		http.Error(res, "404", http.StatusNotFound)
		return 
	}

	if req.Method != "GET" {
		http.Error(res, "405", http.StatusMethodNotAllowed)
		return 
	}
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "405", http.StatusMethodNotAllowed)
		return 
	}

	if err := req.ParseForm(); err != nil {
		http.Error(res, "400", http.StatusBadRequest)
		return
	}

	if req.URL.Path == "/api/form" {
		req.ParseForm()
		fmt.Fprintf(res, "Your name is %s and the reason of the visit is %s", req.FormValue("name"), req.FormValue("info"))
	} else {
		http.Error(res, "404", http.StatusNotFound)
		return 
	}

}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	
	http.Handle("/", fileServer)
	http.HandleFunc("/api/form", formHandler)
	http.HandleFunc("/whoisvin", whoisvinHandler)

	fmt.Println("Webserver listening on port 5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}