package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/url"
)

type GetCoordinateReq struct {
	URL string `json:"name"`
}


type GoogleResponse struct {
	Results []GoogleResult
}

type GoogleResult struct {Address      string               `json:"formatted_address"`
	AddressParts []GoogleAddressPart `json:"address_components"`
	Geometry     Geometry
	Types        []string
}

type GoogleAddressPart struct {
	Name      string `json:"long_name"`
	ShortName string `json:"short_name"`
	Types     []string
}

type Geometry struct {
	Bounds   Bounds
	Location Point
	Type     string
	Viewport Bounds
}
type Bounds struct {
	NorthEast, SouthWest Point
}

type Point struct {
	Lat, Lng float64
}

func getGoogleLocation(Address string) GoogleResponse{
	
	client := &http.Client{}
	
	var req1 GetCoordinateReq;
	req1.URL = "http://maps.google.com/maps/api/geocode/json?address="
	req1.URL += url.QueryEscape(Address)
	req1.URL += "&sensor=false";
	fmt.Println("URL formed: "+ req1.URL)
	
//	reqBody, err := json.Marshal(req1)
//	fmt.Println(err);
//	req, err := http.NewRequest("POST",req1.URL , strings.NewReader(string(reqBody)))
	req, err := http.NewRequest("GET",req1.URL , nil)
	
	resp, err := client.Do(req)
	fmt.Println(resp,err);

	if err != nil {
	fmt.Println("error!");	
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	var res GoogleResponse
	err = json.Unmarshal(body, &res)
	
	fmt.Println("Body:",res);
	
	return res;

}