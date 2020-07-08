package main
import (
	 "fmt"
	_ "strconv"
)
func main(){
	/*
	//TYPE CONVERSION IN GO

	var str string = "Sebahattin";
	var num int = 1;
	var f float32 = 2.5;
	fmt.Println(str,num,f);

	// string to int donusumu
	 str = "1"
	fmt.Println(str);
	 var number,_ = strconv.Atoi(str);
	 fmt.Println(number+2)

	 // intten stringe donusum
		 
	 fmt.Println("Result :" + strconv.Itoa(number))

	 /**********CASTING**********/
	/*
	 var i int = 55;
	 var fs float64 = float64(i);
	 var u uint = uint(fs);

	 fmt.Println(i,fs,u);
	 */

	 /*******Functions********/

	 var sum int= add(3,5);

	 fmt.Println("Sum :" , sum)
	
	 a:=3;
	 b:=5;
	 fmt.Println("a: ", a, "b: ", b)
	 a,b = swap(a,b);
	 fmt.Println("a: ", a, "b: ", b)

	//pass by reference vs pass by value

	var mes  = "Hello";
	fmt.Println(mes);
	sayHello(&mes);
	fmt.Println(mes);

}

func add(num1 int,num2 int) int{
	return num1+num2;
}

func swap(num1 int, num2 int) (int,int) {
	return num2 , num1;
}

func sayHello(message *string){
	fmt.Println(*message);
	//oroginal value changed
	*message = "Sebahattin"
}