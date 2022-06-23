package main

import (
	"fmt"

	"github.com/vivaHSYA/learngo/mydict"
)

func main(){
	// const name string = "hello";
	dictionary := mydict.Dictionary{}
	word := "hello"
	def := "greeting"
	err := dictionary.Add(word,def)

	if err != nil{
		fmt.Println(err)
	}
	
	wd, err1 := dictionary.Update("test","test!")

	wd2, err2 := dictionary.Update("hello","test2");

	
	fmt.Println(wd, err1, wd2, err2);
	
	result, _ := dictionary.Search("test");
	result2, _ := dictionary.Search("hello");
	
	fmt.Println(result, result2);
	
	
}