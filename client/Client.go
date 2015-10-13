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
	var req1 LocationService;
	req1.Name = "Saurabh1"
	req1.Address = "1055 E Evelyn Ave Sunnyvale CA"
	reqBody, err := json.Marshal(req1)
	
	req, err := http.NewRequest("PUT", "http://localhost:8080/locations/1243", strings.NewReader(string(reqBody)))
	resp, err := client.Do(req)

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