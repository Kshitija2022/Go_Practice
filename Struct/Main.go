package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	FirstName string
	LastName  string
	contact   contactInfo
}

func main() {
	//alex := person{FirstName: "Kshitija", LastName: "Gawande"}
	//fmt.Println(alex)
	//var alex person
	//alex.FirstName = "Kshitija"
	//alex.LastName = "Gawande"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)
	xyz := person{
		FirstName: "Kshitija",
		LastName:  "Gawande",
		contact: contactInfo{
			email:   "ksh@gmail.com",
			pincode: 4040405,
		},
	}
	//fmt.Printf("%+v", xyz)

	//xyzPointer := &xyz
	xyz.Updatename("ksh")
	xyz.print()

}
func (pointerToPerson *person) Updatename(NewFirstname string) {
	pointerToPerson.FirstName = NewFirstname
}
func (p person) print() {

	fmt.Printf("%+v", p)
}
