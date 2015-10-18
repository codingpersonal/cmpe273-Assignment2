package main

import (
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {  
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost:8081")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}

func getData(location_id string) LocationService{
	var a LocationService
	return a
}

func setData(location_id string, loc LocationService) bool{
	return true
}

func deleteData(location_id string) bool {
	return true
}
