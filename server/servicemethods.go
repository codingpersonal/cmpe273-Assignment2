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
	    success := setData(req.Id,req)
	    if !success {
	    	fmt.Println("Unable to create an entry in the database")
	    }
    }
    
    json.NewEncoder(w).Encode(req)
    
    if err != nil {
    	fmt.Println("some error");
    }  

}

func GetLocation(w http.ResponseWriter, r *http.Request) LocationService{
	fmt.Println("inside Get location fn");
	vars := mux.Vars(r)
	location_id := vars["id"]
	var res LocationService
	
	//Get the response from Mongo Db for this Location_Id
	res = getData(location_id)
	
	//change this res to the response which needs to be sent back
	
	//If error is nil, then fill the response structure
	res.Id = "Dummy" 
	res.Name = "Dummy"
	res.Address = "Dummy"
	res.City = "Dummy"
	res.State = "Dummy"
	res.Zip = "Dummy"
	res.Coordinate.Lat = "Dummy"
	res.Coordinate.Lng = "Dummy"
	res.ErrorMsg = nil
    
    json.NewEncoder(w).Encode(res)
    return res
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Delete location fn");
	vars := mux.Vars(r)
	location_id := vars["id"]
	
	//Delete this location id data from Mongo Db
	success := deleteData(location_id)
	if !success {
		fmt.Println("Unable to delete entry from Mongo Db")
	}
}


func PutLocation(w http.ResponseWriter, r *http.Request) LocationService {
	vars := mux.Vars(r)
	location_id := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	
	var req LocationService
	var res LocationService
	err = json.Unmarshal(body, &req)
	
	googleresp := getGoogleLocation(req.Address + "+" + req.City + "+" + req.State + "+" + req.Zip);
    fmt.Println("resp is: ", googleresp);
	
	if !ValidateResponseWithRequest(googleresp, req) {
		// modify the request object itself
		res.ErrorMsg = "Invalid Address, cannot update. No such address exists as per Google service";
	} else {
	    res.location_id = location_id
	    res.Address = req.Address
	    res.City = req.City
	    res.State = req.State
	    res.Zip = req.Zip
	    res.ErrorMsg = nil
	    res.Coordinate.Lat = req.Coordinate.Lat
	    res.Coordinate.Lng = req.Coordinate.Lng

	    // TODO: store the response in the mongo db for this Location ID 
	    success:=setData(location_id, res)
	    if !success {
	    	fmt.Println("Unable to update data in the database")
	    }
    }
    
    json.NewEncoder(w).Encode(res)
    return res
}
