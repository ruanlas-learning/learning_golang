package main

import (
	"fmt"
	"log"
	"io/ioutil"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	// fmt.Println(r)
	// fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println()
	fmt.Println("Method:", r.Method)
	fmt.Println("Url:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("RequestURI:", r.RequestURI)
	fmt.Println("Form:", r.Form)
	fmt.Println()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Body")
	fmt.Println(string(body))
	fmt.Println()
	fmt.Println(r.Header)
	fmt.Println(r.Header["Teste"])
	fmt.Println()

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Location", "patati patatá")
	w.WriteHeader(http.StatusBadRequest)
	http.Error(w, `{"error": "fodeu"}`, http.StatusBadRequest)
	// http.Error(w, "{\"error\": \"fodeu\"}", http.StatusBadRequest)
	// http.Error(w, "{\"error\": \"fodeu\"}", 400)
	// fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
    handleRequests()
}