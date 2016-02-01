package main

import (
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type GetMeResult struct {
	Id			int				`json:"id"`
	FirstName	string			`json:"first_name"`
	UserName	string			`json:"username"`
}

type GetMeResponse struct {
	Ok			bool			`json:"ok"`
	Result		GetMeResult		`json:"result"`
}

func getMe() (getMeResponse GetMeResponse) {

	url := BASE_URL + "getMe"
	log.Printf("url: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%s", body)

	err = json.Unmarshal(body, &getMeResponse)
	if err != nil {
		log.Fatal(err)
	}

	return
}