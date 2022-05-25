package main

import "fmt"

//var myName string

func main() {
	fmt.Println("hello world")

	var whatToSay string
	var i int

	whatToSay = "Amit"
	fmt.Println(whatToSay)

	i = 23

	fmt.Println("i value is ", i)

	whatWasSay, theOtherThing, nu := saySomething()

	fmt.Println(whatWasSay, theOtherThing, nu)

}

func saySomething() (string, string, int) {
	return "something", "else", 23
}
