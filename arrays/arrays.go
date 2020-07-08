package main

import (
	"fmt"
)

func main(){
	// for dongusu
	for i:=0; i<5;i++ {
		fmt.Println(i);
	}

	/***Arrays***/
	// array definition
	myArray1 := [3] int {}
	var colors [3] string;
	
	myArray1[0] = 1;
	myArray1[1] = 2;
	myArray1[2] = 3;

	fmt.Println(myArray1)

	//color array
	colors[0] = "blue";
	colors[1] = "green";
	colors[2] = "yellow";

	fmt.Println(colors)
} 