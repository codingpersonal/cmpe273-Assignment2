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
	req.Name = "Somu"
	req.Address = "1 hacker way"
	req.City = "Mountainview"
	req.State = "CA"
	req.Zip = "94025"

	reqBody, err := json.Marshal(req)
	req1, err := http.NewRequest("PUT", "http://localhost:8081/locations/5940", strings.NewReader(string(reqBody)))
	resp, err := client.Do(req1)

	fmt.Println("resp code is : " , resp);
	if err != nil {
		fmt.Println("error!");	
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var res LocationService
	err = json.Unmarshal(body, &res)

	fmt.Println("Body:",res);

//	url := "http://localhost:8081/locations/" + "5940"
//	fmt.Println("url is : " + url);
//	req2, err2 := http.NewRequest("GET", url, strings.NewReader(""))
//	res2, err2 := client.Do(req2)
//
//	fmt.Println(res2,err2);
//	if err2 != nil {
//	fmt.Println("error!");	
//	}
//
//	defer res2.Body.Close()
//	body1, err2 := ioutil.ReadAll(res2.Body)
//	
//	var res1 LocationService
//	err2 = json.Unmarshal(body1, &res1)
//	
//	fmt.Println("Body:",res1);
}
