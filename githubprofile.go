package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const api_endpoint string = "https://api.github.com/users/"

// Profile describes a Github profile
type Profile struct {
	Login    string
	Name     string
	Company  string
	Location string
	Email    string
}

func (p Profile) String() string {
	return fmt.Sprintf("Name: %s (%s)\nCompany: %s\nEmail: %s\nLocation: %s", p.Name, p.Login, p.Company, p.Email, p.Location)
}

// Main function to call
func main() {
	username := flag.String("login", "", "The Github login")
	flag.Parse()

	if *username == "" {
		log.Fatal("Usage: githubprofile --login LOGIN")
	}

	url := fmt.Sprintf("%s%s", api_endpoint, *username)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusNotFound {
		log.Fatal(fmt.Sprintf("%s: user does not exist", *username))
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	r := new(Profile)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
