package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	id := r.URL.Query().Get("id")
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

	var article map[string]string
	err = json.Unmarshal(body, &article)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(article)
}

func main() {
	http.HandleFunc("/articles", Articles)
	http.HandleFunc("/article", Article)
	log.Fatal(http.ListenAndServe(":80", nil))
}
