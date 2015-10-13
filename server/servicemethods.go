package main

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"github.com/gorilla/mux"
)

// defines all the REST APIs handlers

// kind of stubs
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside create location fn");

}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Get location fn");

}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Delete location fn");
}


func PutLocation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    Id := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	
	var req LocationService
	err = json.Unmarshal(body, &req)
	
	googleresp := getGoogleLocation(req.Address);
	fmt.Println(googleresp);
	
    // get data for that id from mongo
    var res LocationService;
    res.Id = Id
    res.Name = req.Name
    
    json.NewEncoder(w).Encode(res)
    
    if err != nil {
    	fmt.Println("some error");
    }  
}
