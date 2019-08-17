package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const api = "https://api.hurriyet.com.tr/v1"

var apikey = os.Getenv("HURRIYET_API_KEY")

// Articles fetches newest articles from Hürriyet API
func Articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	url := api + "/articles?$select=Id,Title"

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("apikey", apikey)
	req.Header.Set("host", "lite.hurriyet.com.tr")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, string(body))
}

// Article fetches article detail by article ID
func Article(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	url := api + "/articles/" + id + "?$select=Title,Text,CreatedDate,Editor"

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("apikey", apikey)
	req.Header.Set("host", "lite.hurriyet.com.tr")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, string(body))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/articles", Articles).Methods("GET")
	mux.HandleFunc("/articles/{id}", Article).Methods("GET")
	log.Fatal(http.ListenAndServe(":8091", mux))
}
