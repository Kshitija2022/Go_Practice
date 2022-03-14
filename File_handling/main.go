package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("WQelcome to files in golang")
	content := "This needs to go in file"
	file, err := os.Create("./myGoFile.txt")
	//if err != nil {
	//	panic(err)
	//}
	checkNilError(err)
	length, err := io.WriteString(file, content)
	//if err != nil {
	//	panic(err)
	//}
	checkNilError(err)
	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//}
	checkNilError(err)
	fmt.Println("Text data inside the file is\n ", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
