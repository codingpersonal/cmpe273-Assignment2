package main

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"strings"
)

func main() {

	client := &http.Client{}
	var req LocationService;
	req.Name = "Saurabh1"
	req.Address = "1055 E Evelyn Ave"
	req.City = "Sunnyvale"
	req.State = "CA"
	req.Zip = "94086"
	
	reqBody, err := json.Marshal(req)
	
	req1, err := http.NewRequest("POST", "http://localhost:8080/locations", strings.NewReader(string(reqBody)))
	resp, err := client.Do(req1)

	fmt.Println(resp,err);
	if err != nil {
	fmt.Println("error!");	
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	var res LocationService
	err = json.Unmarshal(body, &res)
	
	fmt.Println("Body:",res);
}