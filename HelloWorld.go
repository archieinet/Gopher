package main

import ("fmt")

const (
  one = iota
  
  pi int = 3 
  two = iota
  name string = "Archie 🧑 "
  
)
/*
  Go have invisible ; func should have { right after declaration.
  line 9 it'll will add ; at runtime
*/
func main() {
	fmt.Println("Hello, Go!")

	var ptOne *string = new(string) //create pointer using * and new to create instance
	*ptOne = name             //assign value to ptOne location
	fmt.Println(ptOne, *ptOne)

	*ptOne = "🧑?"
	ptTwo := &*ptOne           //& - address operator
	fmt.Println(ptTwo, *ptTwo) //*de-reference

	
	fmt.Println("🌠 It's not a pi: ", pi)
	fmt.Println("🥧:", float32(pi)+0.1415)  


  fmt.Println(one,2 << two)
  
}
