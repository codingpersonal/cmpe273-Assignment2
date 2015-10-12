package main

import ("net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"strings"
)

type GetLocationReq struct {
	Name string `json:"name"`
	Address string  `json:"add"`
}

type GetLocationRes struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Lat string `json:"lat"`
	Lon string `json:"long"`
}

func main() {

	client := &http.Client{}
	var req1 GetLocationReq;
	req1.Name = "Saurabh1"
	req1.Address = "1600 Amphitheatre Parkway Mountain View CA"
	reqBody, err := json.Marshal(req1)
	
	req, err := http.NewRequest("PUT", "http://localhost:8080/locations/1234", strings.NewReader(string(reqBody)))
	resp, err := client.Do(req)

//	resp, err := http.Put("http://localhost:8080/locations/12345");
	fmt.Println(resp,err);
	if err != nil {
	fmt.Println("error!");	
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	var res GetLocationRes
	err = json.Unmarshal(body, &res)
	
	fmt.Println("Body:",res);
}