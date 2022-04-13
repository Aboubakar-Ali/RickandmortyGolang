package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Info    string      `json:"info"`
	Results []Character `json:"results"`
}

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	URL     string   `json:"url"`
	Created string   `json:"created"`
}

var Characters = []Character{}

func homePage(w http.ResponseWriter, r *http.Request) {
	url := "https://rickandmortyapi.com/api/character"

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Info)
	fmt.Println(len(responseObject.Results))

	for i := 0; i < len(responseObject.Results); i++ {
		fmt.Println(responseObject.Results[i].Name)
	}

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	tmpl.Execute(w, responseObject)

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8088", nil)

}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	handleRequests()
}
