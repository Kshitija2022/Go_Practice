package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON ")
	//EncodeJson()
	Decodedata()
}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJS Bootcamp", 200, "WWw.google.com", "abc@21", []string{"web-dev", "js"}},
		{"Angular", 299, "WWw.google.com", "pqr@21", nil},
		{"Javascript Bootcamp", 300, "WWw.google.com", "xyz@12", []string{"full stack", "js"}},
	}
	//pacakage this data as JSON data
	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func Decodedata() {
	Jsondatafromweb := []byte(`
	{
                "CourseName": "ReactJS Bootcamp",
                "Price": 200,
                "Platform": "WWw.google.com",
                "tags": ["web-dev","js"]
    }
	`)
	var lcocourse course
	checkValid := json.Valid(Jsondatafromweb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(Jsondatafromweb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
      var mydata map[string]interface{}
	    json.Unmarshal(Jsondatafromweb, &mydata)
	    fmt.Printf("%#v\n", mydata)

}

