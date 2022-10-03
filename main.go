package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {

	nr_license := os.Args[1]

	fmt.Println("nr_license:" + nr_license)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("GO_API_SAMPLE_3"),
		newrelic.ConfigLicense(nr_license),
	)

	if err != nil {
		log.Fatal(err)
	}

	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", usersHandler))

	// http.HandleFunc("/", homePage)
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", homePage))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
