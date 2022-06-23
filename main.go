package main

import "fmt"


type person struct{
	name string 
	age int
	favFood []string
}

func main(){
	// const name string = "hello";


	favFood := []string{"test1","test2"}
	test := person{"jiwan",18,favFood}
	fmt.Println(test);
	

}