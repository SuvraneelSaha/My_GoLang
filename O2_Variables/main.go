package main

import (
	"fmt"
)

// numberOfUsers1 := 30000;
// fmt.Println(numberOfUsers1);
// fmt.Printf("Variable is of type : %T \n",numberOfUsers1);

// here we cant declare like this outside of a function
// the shorthand initialization only works inside a function

// so we can initialize it using the var only outside of function


	const JwtToken_Example string = "sdfghjklpoiuytrewqasdfghjklmnbvcxz";
	// it is a constant variable 
	// Capital letter means it is public variable -- the starting letter of the variable





func main(){
	fmt.Println("Hello, World from O2_Variables");
	var username string = "Meow";
	// if you created or initialized a variable and not used it, it will throw an error
	fmt.Println("the name is ",username);
	fmt.Printf("Variable is of type : %T \n",username);
	// %T is used to print the type of the variable 

	var isLoggedIn bool = true; 
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type : %T \n",isLoggedIn);

	// 
	var smallVal uint8 = 255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type : %T \n",smallVal);

	// we can basically use the int for more simplicity 

	var smallVal1 int = 255;
	fmt.Println(smallVal1);
	fmt.Printf("Variable is of type : %T \n",smallVal1);
	// this is int 

	var smallfloat float32 = 255.4555555;
	fmt.Println(smallfloat);
	fmt.Printf("Variable is of type : %T \n",smallfloat);

	// float64 
	var smallfloat1 float64 = 255.69696969;
	fmt.Println(smallfloat1);
	fmt.Printf("Variable is of type : %T \n",smallfloat1);

	// depends on the precision 

	var avr int ; 
	fmt.Println(avr);
	fmt.Printf("Variable is of type : %T \n",avr);

	var avr1 float64;
	fmt.Println(avr1);
	fmt.Printf("Variable is of type : %T \n",avr1);

	var avr2 string;
	fmt.Println(avr2);
	// default value is empty string 
	fmt.Printf("Variable is of type : %T \n",avr2);

	var website = "www.google.com";
	fmt.Println(website); 
	// website = 3 ; 
	// we cannot change the type of variable
	// once assigned we cant change the type of that var 

	// no var style 

	numberOfUsers := 30000;
	fmt.Println(numberOfUsers);
	fmt.Printf("Variable is of type : %T \n",numberOfUsers);


	// := it is walrus operator 

	// jwttoken example 
	fmt.Println(JwtToken_Example);
	fmt.Printf("Variable is of type : %T \n",JwtToken_Example);













}