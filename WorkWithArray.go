 package main
import(
	"fmt"


)



func main(){
	//array fixed size collection
	arr := [3]string{"👍","😋","😇"} //at intial
	var ary [3]int //declar 
	ary[0] = 1
	ary[1] = 2
	ary[2] = 3
	fmt.Println("🚆",ary[0])
	fmt.Println("👾👾",arr[1])

	//slice
	slice1 := arr[0:2] // : all, start:end

	fmt.Println("slice:",slice1)

	slice2 :=  []int{4,5,6,7,7,9,10}
	fmt.Println("slice:",slice2)

	//append
	slice3 := arr[:] //type and value 
	slice3 = append(slice3,"🐱")
	fmt.Println(slice3)

	slice4 := slice3[1:] //from 1-end
	fmt.Println(slice4)
	
	slice5 := slice3[:2] // from 0-2
	slice5 = append(slice5,"🐶") //add new element
	fmt.Println(slice5)



}