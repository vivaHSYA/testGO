package main

import (
	"fmt"

	"github.com/vivaHSYA/learngo/account"
)

func main(){
	// const name string = "hello";

	result := account.NewAccount("jiwan");
	
	fmt.Println(result);

}