package main

import (
	"fmt"
)

func runvariables(){
	// Long hand way  - declare and then assign
	var firstNumber int
	//assigning
	firstNumber = 2
	fmt.Println(firstNumber)

	//We can assign multiple variables of
	//same dataype in single line
	var secondNumber,thirdNumber int
	// By default, number 0 is assigned
	//if we do not initialize a variable if
	//it is a string variable "" blank string is assigned

	thirdNumber=firstNumber+secondNumber

	//var 1variable
	// variable must being with a alphabet

	//One more way of declare variable
	//Go will set the type based on the right side
	secondNumber = 5
	fmt.Println(thirdNumber)

	var number =6
	fmt.Println(secondNumber)
	fmt.Println(number)

	//One step variable
	//This is shortest form where we do not even use var key word
	thirdVariable:=7
	fmt.Println(thirdVariable)

}
