package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

//Planet is a planet type
type Planet struct {
	Name string `json:"name"`
	Population string `json:"population"`
	Terrain string `json:"terrain"`
}

//Person is a person type
type Person struct {
	Name string `json:"name"`
	HomeWorldURL string `json:"homeworld"`
	HomeWorld Planet  
}

// All people is a collection of Person types
type AllPeople struct {
	People []Person `json:"results"`

}

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeWorldURL)

	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err !=nil {
		log.Print("Error log response body", err)
	}

	json.Unmarshal(bytes, &p.HomeWorld)
}

func getPeople(w http.ResponseWriter, r *http.Request){
	 res, err := http.Get(BaseURL + "people")
	 if err!=nil {
		http.Error(w, err.Error(),http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	 }
 

	 bytes, err := ioutil.ReadAll(res.Body)

	 if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse the request body")
	 }

	// fmt.Println(string(bytes))
	 var people AllPeople
	 if err := json.Unmarshal(bytes, &people); err !=nil {
		fmt.Println("Error parsing json", err)
	 }

	 for _, pers := range people.People {
		pers.getHomeWorld()
		fmt.Println(pers)
		
	 }

	
}

func main() {
	 
	http.HandleFunc("/people", getPeople)
	fmt.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	

}
