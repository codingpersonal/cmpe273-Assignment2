package main

import (
    "fmt"
	"github.com/gorilla/mux"
    "net/http"
    "encoding/json"
    "log"
    "time"
    "io/ioutil"
)

func main() {
	router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation",
        Completed: false,
        Due: time.Now()},
        Todo{Name: "Host meetup",
        Completed: true,
        Due: time.Now()},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

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


func PutLocation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    Id := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	
	var req GetLocationReq
	err = json.Unmarshal(body, &req)
	
	googleresp := getGoogleLocation(req.Address);
	fmt.Println(googleresp);
	
    // get data for that id from mongo
    var res GetLocationRes;
    res.Id = Id
    res.Name = req.Name
    res.Lat = ""
    res.Lon = ""
    
    json.NewEncoder(w).Encode(res)
    
    if err != nil {
    	fmt.Println("some error");
    }  
//    fmt.Fprintln(w, "Showing location for id :", Id)
}
