package main

import "fmt"

type Person struct {
	Name, Country, JobDescription string
}

func main() {
	rochman := Person{
		Name:           "Rochman Ramadhani",
		Country:        "Indonesia",
		JobDescription: "Google - Tech Lead",
	}
	ramadhani := &rochman
	ramadhani.Name = "Ramadhani Rochman"
	fmt.Println(rochman)
	fmt.Println(ramadhani)

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)    // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	fmt.Println("numberB (value) :", *numberB)   // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
}
