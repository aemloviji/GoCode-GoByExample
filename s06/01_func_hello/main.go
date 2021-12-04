package main

import "fmt"

func main() {
	sayHi()
	sayGreetings("ESP")
	sayGreetings("GER")
	sayGreetings("ENG")

	fmt.Println()
	fmt.Printf("Type of sayHi: %T\n", sayHi)
	fmt.Printf("Type of sayGreetings: %T\n", sayGreetings)
}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayGreetings(lang string) {
	//if lang == "ESP" {
	//	fmt.Println("Hola")
	//} else if lang == "GER" {
	//	fmt.Println("Hallo")
	//} else {
	//	fmt.Println("Hello")
	//}

	switch lang {
	case "ESP":
		fmt.Println("Hola")
	case "GER":
		fmt.Println("Hallo")
	default:
		fmt.Println("Hello")
	}
}
