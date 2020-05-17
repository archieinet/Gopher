package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 42} //key string type, value are integer

	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 7  //map are NOT readonly
	fmt.Println(m["foo"])

	delete(m, "foo") //delete from m variable with key == "foo"
}
