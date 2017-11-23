package main

import (
	"fmt"
	//"golang_starter/stringutil"
)

func trace(s string) string{
	fmt.Println("entering: ", s)
	
	return "shit " + s
}

func un(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	
	defer un(trace("a"))
	fmt.Println("something in a")
}


func main() {
	a := "hello"
	fmt.Println(a)
	
}
