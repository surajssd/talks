package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

type Name struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

type Contact struct {
	Twitter string `json:"twitter,omitempty"`
	IRC     string `json:"irc,omitempty"`
	Slack   string `json:"slack,omitempty"`
	Mail    string `json:"mail,omitempty"`
}

type About struct {
	Name     Name     `json:"name,omitempty"`
	Company  string   `json:"company,omitempty"`
	ReachOut Contact  `json:"reach_out,omitempty"`
	Projects []string `json:"projects,omitempty"`
}

func main() {

	introContents, err := ioutil.ReadFile("intro_example.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	var aboutMe About
	err = json.Unmarshal(introContents, &aboutMe)
	if err != nil {
		logrus.Fatalln(err)
	}
	fmt.Printf("scanned: %#v", aboutMe)
}
