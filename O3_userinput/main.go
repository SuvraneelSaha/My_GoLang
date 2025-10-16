package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// taking user input
	welcome:= "Welcome to this exercise";
	fmt.Println(welcome);

	// buffio package 
	reader:= bufio.NewReader(os.Stdin);
	fmt.Println("Enter the food ratings");
	// now we want to store the input of the user inside of a variable 

	// comma ok  || error err
	input, _ := reader.ReadString('\n');
	fmt.Print("Thanks for the rating ",input);
	fmt.Printf("Type of the rating is %T ",input);
	// print , printf nd println 
	// comma error syntax 
	// _ , err := reader.ReadString('\n'); -- like this also works
	


	

}
