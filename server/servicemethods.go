package main

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
//		"github.com/gorilla/mux"
		"math/rand"
		"strconv"
)

// defines all the REST APIs handlers

// this will vlaidate the response from google with the req asked by user
func ValidateResponseWithRequest(res LocationService, req LocationService) bool {
	return res.City == req.City && res.Zip == req.Zip
}

// kind of stubs
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	
	var req LocationService
	err = json.Unmarshal(body, &req)
	
	googleresp := getGoogleLocation(req.Address + "+" + req.City + "+" + req.State + "+" + req.Zip);
    fmt.Println("resp is: ", googleresp);
	
	if !ValidateResponseWithRequest(googleresp, req) {
		// modify the request object itself
		req.ErrorMsg = "Invalid Address. No such address exists as per Google service";
	} else {
	    req.Id = strconv.Itoa(rand.Intn(10000))
	    req.Coordinate.Lat = googleresp.Coordinate.Lat
	    req.Coordinate.Lng = googleresp.Coordinate.Lng

	    // TODO: store the response in the mongo db
    }
    
    json.NewEncoder(w).Encode(req)
    
    if err != nil {
    	fmt.Println("some error");
    }  

}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Get location fn");
//	    vars := mux.Vars(r)
	

}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Delete location fn");
}


func PutLocation(w http.ResponseWriter, r *http.Request) {
}
