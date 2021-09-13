package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}
type Social struct {
	XMLName xml.Name `xml:"social"`
	Youtube string   `xml:"youtube"`
}

func unwrap(errorMessage string, err error) {
	if err != nil {
		log.Fatalf("error: %s.\n %v", errorMessage, err)
	}
}

func init() {
	fmt.Println("this is the start of an xml reader")
}

func main() {
	var users Users

	xmlFile, err := os.Open("demoxml.xml")
	unwrap("opening xmlFile", err)

	fmt.Println("successfullly opened xml file.")

	// unmarshal the xml
	bv, err := ioutil.ReadAll(xmlFile)
	unwrap("readall error at line 40", err)

	xml.Unmarshal(bv, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("user type: ", users.Users[i].Type)
	}

	defer xmlFile.Close()
}
