package main

import (
)

type LocationService struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Coordinate struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}
	ErrorMsg string `json:"error"`			// to send any error back from server to client
}
