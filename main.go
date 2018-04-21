package main

import "fmt"

//No hacen falta extends
//Cualquier custom type que contenga una func
//Que se llame getGreeeting() y devuelva un string
//Se convierte automáticamente al tipo "bot"
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Very custom logic for generating english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//Very custom logic for generating spanish greeting
	return "Hola!"
}