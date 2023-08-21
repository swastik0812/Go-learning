package main

import (
	"fmt"
)

var deckSize int

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastNAme  string
	contactInfo
}

type name string

type bot interface {
	getGreeting() string
}

type logWriter struct{}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

// above lines means that any type inside over program that implements the function get
// greeting that returns the string is also an honorary member of bot/\.

// -------***********

//we can expect multiple function there return type an parameter they except inside the curly braces

// type bot interface {
// 	getGreeting(string, int) (string, error)
// 	getBoatVersion() float64
// 	respondToUser(user) string
// }

// _______*****
// there are to type of Types
// 1) concrete type eg:-map, struct, int, string, englishBot --> these are those by using which we can directly create value
// 2) bot -> we can not create a value with directly value type boat here as it is of type interface
// 3) interface are not generic types in Go
// 4) interface are implicit we do not need to write type englishBot implements bot struct{} , we do not manual
//  have to say that  our custom type satisfies some interface

type englishBot struct{}
type spanishBot struct{}

func main() {

	// var card string = "deded"
	//  height := "5"

	//  height =5

	// height = "9"
	// deckSize = 1

	// fmt.Println(card)
	// fmt.Println(height)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards1 := deck1{2, 1}

	// cards1.print()
	// fmt.Println("$$$$$$$$$$")

	// cards = append(cards, "six of spade")
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// *******----slice
	// cards := newDeck()
	// fmt.Println("00000", cards[0])
	// cards[0] = "99999"
	// cards = append(cards, "5678")
	// fmt.Println("-------------")
	// fmt.Println(cards[0])
	// fmt.Println("-------------")
	// fmt.Println(cards[0:2])
	// fmt.Println("-------------")
	// fmt.Println(cards[2:])
	// fmt.Println("-------------")
	// cards.print()

	// fmt.Println("-------------*******")

	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// fmt.Println("-------------*******")
	// remainingCards.print()

	//data type conversion
	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))

	// fmt.Println("-------------*******")

	// fmt.Println(reflect.TypeOf(cards.toString()))

	// fmt.Println("-------------*******")
	// cards.saveToFile("my_card")
	// fmt.Println("-------------*******")
	// newCards := newDeckFromFile("my_cards")

	// newCards.print()

	// x := [5]int{10, 20, 30, 40, 50}

	// fmt.Println(x)
	// fmt.Println("-------------*******xxxx")
	// y := [...]int{10, 20, 30}

	// z := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(z)
	// fmt.Println("-------------*******xxxxyyyyyyy")
	// z[len(z)-1] = 9
	// y[len(y)] = 10
	// n := []int{}
	// for i, val := range y {

	// 	n = append(n, )

	// }

	// fmt.Print("nnnnnnnn---->", reflect.TypeOf(z))

	// fmt.Println(y)
	// fmt.Println("-------------*******xxxx")
	// cards := newDeck()
	// cards.print()
	// fmt.Println("-------------*******")

	// cards.shuffle()
	// cards.print()

	// struct,struct,struct,struct,struct,struct,struct,struct,struct,struct,struct,struct,struct,struct

	//one way to assign value to struct

	// alex := person{"Swastika","Pal"}

	//another way to assign value to struct

	// alex := person{
	// 	firstName: "Swastik",
	// 	lastNAme:  "Pal",
	// 	contactInfo: contactInfo{
	// 		email:   "Swastik@gmail.com",
	// 		zipCode: 10009,
	// 	}}

	// raj :=
	// 	person{
	// 		firstName: "RAj",
	// 		lastNAme:  "RRR",
	// 		contactInfo: contactInfo{
	// 			email:   "Raj@gmail.com",
	// 			zipCode: 1000,
	// 		},
	// 	}

	// var val = person{}
	// val.firstName = "Robin"
	// val.lastNAme = "Pal"

	// fmt.Println("alex--->", alex)
	// fmt.Println("alex--->", val)
	// fmt.Println("raj--->", raj)
	// raj.print()

	// rajPointer := &raj
	// rajPointer.updateName("tushar")
	// fmt.Println("-------------*******")
	// raj.print()
	// fmt.Println("-------------*******")
	// another syntax in that go automatically pass the memory address by value
	// alex.updateName("Robin")
	// alex.print()

	// var myName name = "swastik"
	// fmt.Println(myName)
	// myNamePointer := &myName
	// myNamePointer.updateName()
	// &myNamePointer
	// myName.updateName()
	// fmt.Println("--------")

	// fmt.Println(myName)

	// ----------------------------MAP-------------------

	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"white": "#ffffff",
	// }

	// // var color map[string]string

	// color := make(map[int]string)
	// printMap(color)
	// fmt.Println("map------->", color)

	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)

	// *******-------Http
	// resp, err := http.Get("http://google.com")

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("bsssss--->", string(bs))
	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	tri := triangle{
		height: 10.0,
		base:   10.0,
	}

	sqr := square{
		sideLength: 10.0,
	}

	printArea(tri)
	printArea(sqr)

}

// struct with receiver

// &variable = it gives us the memory address of the value this variable is pointing at
// *pointer = it gives us the value on which the memory address is pointed at
//***important---> in below code the *person is a type description it means we are working with the pointer to a person
// & sign use to get the memory address by value
// * sign is use to get the value by memory address
// slices,maps,channels,pointers,function are reference type in go
// int , float, string,bool,struct are value types in go

// func (pointerPerson *person) updateName(newFirstName string) {
// 	(*pointerPerson).firstName = newFirstName
// }

func (pointerName *name) updateName() {
	fmt.Println("memory allocation-->", pointerName)
	(*pointerName) = "Vishal"
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for " + color + " is " + hex)
	}
}

func (englishBot) getGreeting() string {

	return "hi there"

}

func (spanishBot) getGreeting() string {

	return "Hola!"

}

// bellow two printGreeting function without using interface

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// bellow printGreeting with interface

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we need to declare the return type at the time of declaration of func if we want to return something
// func newCard() string {
// 	return "Ace of spade"
// }

// {The range keyword is used to iterate over an array, slice, string, map, or channel,
//  providing both the index and the corresponding value in each iteration.}

//Go is not a dynamic types as js but it is a static types

// package main

// func main() {
//     printState()
// }

// two basic data structure in Go use is Array  an another is slice

// Array-> Fixed  length of records
// Slice -> Slice is like an array that can grow or reduce
// every element in array or slice must me of same data type
// func (p person) print() {

// 	fmt.Printf("%+v", p)
// }

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func (tr triangle) getArea() float64 {

	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {

	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
