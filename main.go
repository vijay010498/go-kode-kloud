package main

import (
	"fmt"
	"reflect"
)

func main() {
	//fmt.Println("Hello World")
	//name := "Vijay"
	//fmt.Println(name)
	//var greeting string = "Hello World"
	//fmt.Println(greeting)
	//fmt.Print("Hello World")
	//var city string = "Chennai"
	//fmt.Print(city)
	//fmt.Println("Welcome to ", city, ", ", city)
	////\n new line
	//fmt.Println("Welcome to ", city, ",\n ", city)
	//// Format specifier
	//// %v
	//fmt.Printf("Welcome to, %v", city)
	//// %d - decimal Integers
	//var grades int = 42
	//fmt.Printf("Marks: %d", grades)
	//var name string = "vijay"
	//var i int = 78
	//fmt.Printf("Hey, %v! you have scored %d/100 in maths", name, i)
	//name := "vijay"
	//name = "pablo"
	//fmt.Println(name)
	//city := "London"
	//{
	//	country := "UK"
	//	fmt.Println(country)
	//	fmt.Println(city)
	//}
	//fmt.Println(city)
	//var name string
	//var is_muggle bool
	//fmt.Print("Enter your name and are you a muggle: ")
	////fmt.Scanf("%s", &name)
	//fmt.Scanf("%s %t", &name, &is_muggle)
	//fmt.Println(name, is_muggle)
	var grades int = 42
	fmt.Printf("Variable grades=%v is of type %T \n", grades, grades)
	fmt.Printf("Type: %v \n", reflect.TypeOf("Vijay"))
}
