package main

// Every go program starts with package declaration.
// Package can be shared as a library
// import to include code from other packages.
// "fmt" -Input and output
// Main is a special package
// main func is the stating point inside main package
// Every pgm should contain a main and main func

// Data Types
// staticT Typed - Compiler throws an error wen types are used incorrectly. Ex. C++, Java
// Adv :  Better performance, Data integrity
// Dynamic Typed - JS, Python - Compiler does not enforce the type system
// ADV. Faster to write code, Shorter learning curve

// Golang -  Either explicitly declaration ot inferred by compiler,
// - Fast, statically typed compiled language feels like dynamically typed
// Shot variable declaration // name:= "Vijay"

// Kinds of Data Types
// 1. Integers - int, uint - only positive include zero
// 2. Float - float32, float64
// 3. String - string - "abc" - 16 bytes
// 4. Boolean - bool
// 5. Array, slices, maps -

// Variables, Syntax and declaration
// Declaring - Explicitly or implicitly
// 1st way : var s string = "Hello World"
// var i int = 100
// var b bool = false
// var f float64 = 44.45

// Printing Variables
// fmt.print
// fmt.println
// new line character - \n
// fmt.printf - %s Object
// Format specifier - %v - default format // fmt.Printf("Welcome to, %v", city)
// %d - format decimal integers // fmt.Printf("Marks: %d", grades)
// %T - type of the value
// %c - character
// %q - quoted string/character
// %s- plain string
// %t -true or false
// %f-floating numbers
// %.2f-floating numbers upto 2 decimal places

// Declaring Variables
// Shorthand// var s,t string = "foo", "bar"
// shorthand different data types// var(
//s string = "foo"
//i int = 5)
// short variable declaration // s:= "Hello World" // no need to use var or data types

// Variable scope
// Block  { //outer block  { //inner block }  }
// Local and Global Variables
// Local - Inside a function of block
// Global- Outside a function or block -accessed from any part of the program
