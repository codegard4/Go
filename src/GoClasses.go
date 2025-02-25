package main

import (
	"fmt"
)

type ExampleClass struct {
	value   int
	aString string
}

func (m *ExampleClass) Increment() {
	m.value++
}

func (m *ExampleClass) Decrement() {
	m.value--
}

// declare
func (WhateverNameYouWant *ExampleClass) ChangeString(b string) {
	WhateverNameYouWant.aString = b
	fmt.Println("Hey! you changed aString to: ", WhateverNameYouWant.aString)
}

func (m *ExampleClass) PrintValue() {
	fmt.Println("Current value:", m.value)
}

// in Go we can only have one function with the same name in the same package
// so in HelloWorld.go we can't have a main method if it is also in the src folder
func main() {
	mock := ExampleClass{value: 0}
	mock.Increment()
	mock.PrintValue()
	mock.Decrement()
	mock.PrintValue()
	// if we declared mock string & never used it later on, we will get an error
	mockString := ExampleClass{aString: "Hello, World!"}
	changeString := "Hello, Go!"
	mockString.ChangeString(changeString)

}
