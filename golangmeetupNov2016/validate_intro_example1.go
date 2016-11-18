package main

import (
	"encoding/json"
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

	// Limit on name length
	if len(aboutMe.Name.FirstName) > 30 {
		logrus.Fatalln("Firstname length is more than 30")
	}

	if len(aboutMe.Name.LastName) > 30 {
		logrus.Fatalln("Lastname length is more than 30")
	}

	allowed_projects := []string{"kompose", "openshift", "kubernetes"}
	isprojectallowed := func(project string) bool {
		for _, validProject := range allowed_projects {
			if project == validProject {
				return true
			}
		}
		return false
	}
	for _, project := range aboutMe.Projects {
		if !isprojectallowed(project) {
			logrus.Fatalf("Invalid project value detected %q", project)
		}
	}

}
