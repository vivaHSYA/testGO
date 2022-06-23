package main

import "fmt"


func main(){
	// const name string = "hello";

	names := [5]string{"1","2","3","4","5"}
	names[3] = "ala"
	fmt.Println(names)

}