package main

import "fmt"

/*
  Go have invisible ; func should have { right after declaration.
  line 9 it'll will add ; at runtime
*/
func main() {
	fmt.Println("Hello, Go!")

	var ptOne *string = new(string) //create pointer using * and new to create instance
	*ptOne = "~Archie"              //assign value to ptOne location
	fmt.Println(ptOne, *ptOne)

	*ptOne = "Me?"
	ptTwo := &*ptOne           //& - address operator
	fmt.Println(ptTwo, *ptTwo) //*de-reference

	const pi int = 3 
	fmt.Println("🌠 It's not a pi: ", pi)
	fmt.Println("🥧:", float32(pi)+0.1415)  

}
