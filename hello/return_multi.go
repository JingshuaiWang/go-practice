package main

import "fmt"

func swap(x, y string)(n, m string){
	n = x
	m = y
	return
}

func main(){
	a, b := swap("hello", "world") 
	fmt.Println(a, b)
}
