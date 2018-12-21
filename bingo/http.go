package main

import (
	"log"
	"net/http"
)

func request(url string) (*http.Response, error){
	resp, err := http.Get(url)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200{

	}

	return resp, nil
}
