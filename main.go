// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program shows what price you have to pay to get in the museum.

package main

import "fmt"

func main() {
  var firstNumber int
  var secondNumber int
  var addedNumber = 0
  var answer = 0

	fmt.Println("Enter the two numbers you want to multiply (Integers only)")
	fmt.Println()
	fmt.Print("First Number: ")
	fmt.Scanln(&firstNumber)
	fmt.Print("Second Number: ")
	fmt.Scanln(&secondNumber)

  if firstNumber > 0 && secondNumber > 0 {
		for addedNumber < secondNumber {
      addedNumber = addedNumber + 1;
      answer += firstNumber;
    }
	} else if firstNumber < 0 && secondNumber < 0 {
		for addedNumber > secondNumber {
      addedNumber--;
      answer -= firstNumber;
    }
  } else if firstNumber > 0 && secondNumber < 0 {
    for addedNumber > secondNumber {
      addedNumber--;
      answer -= firstNumber;
    }
  } else if firstNumber < 0 && secondNumber > 0 {
    for addedNumber < secondNumber {
      addedNumber++;
      answer += firstNumber;
    }
  }
	
	fmt.Println(firstNumber, " x ", addedNumber, " = ", answer);
}
