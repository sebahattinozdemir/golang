package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	str1:="The Quick"
	str2:="Melek"
	str3:="Sila"

	aNumber:= 42;
	isTrue := true;
	
	fmt.Println(aNumber,isTrue)

	strlen,err := fmt.Println(str1,str2,str3);
	
	if err == nil{
		fmt.Println("str len",strlen)
	}

	fmt.Printf("isTrue %v \n", isTrue)
	fmt.Printf("Number %v\n",aNumber)
	fmt.Printf("Value as float %.2f\n",float64(aNumber))

	// data types place holder yuzde t ile gosterilir
	fmt.Printf("Data Types %T %T %T %T %T\n",str1,str2,str3,isTrue,aNumber)


	reader:= bufio.NewReader(os.Stdin);	
	fmt.Print("Enter a text: ");
	str,_ := reader.ReadString('\n');
	
	//string to int donusumu
	num,err := strconv.ParseFloat(strings.TrimSpace(str),64);
	if err == nil{
		fmt.Println("Value of Number",num);
	} else{

		panic(err)
	}

	//fmt.Println(str);

}