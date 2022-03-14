package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video ")
	//PerformGetRequest()
	//PerformPostJsonResquest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	var responseString strings.Builder
	//fmt.Println(string(content))
	//fmt.Println(content)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCpunt is :", byteCount)
	fmt.Println(responseString.String())
}
func PerformPostJsonResquest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
	"coursename":"Let's go wit golang",
	"price":0,
	"platform":"learncode.com"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname", "XYZ")
	data.Add("Lastname", "ABC")
	data.Add("email", "XYZ@gmail.com")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
