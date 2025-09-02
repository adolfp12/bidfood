package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var users = make(map[string]string)

func createUser(name string) error {
	log.Printf("Data {%s}", name)
	if _, exists := users[name]; exists {
		/* Issue #2 need to check if the input exists in the map
		if it does, return an error. If it doesn’t, insert it into the map.
		*/
		return errors.New("user already exist!!")
	}
	users[name] = time.Now().String()
	log.Println("Inserted!!")
	return nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" { // Issue #1 need to validate the input such as empty value.
		fmt.Fprintln(w, "Error : `name` can't empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// go createUser(name)
	err := createUser(name)

	if err != nil {
		fmt.Fprintln(w, "Error : ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "User inserted! ")
	log.Println("Finished")
	w.WriteHeader(http.StatusOK)
}
