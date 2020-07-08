package main

import(
	"fmt"
)

type Human struct{
	firstName string
	lastName string
	age int
}

// senin olusturmanLazim varsayilan bos contructor olusturma
func newHuman() *Human{
	h:= new(Human);
	return h;
}

//constructor with params
// Golang method overload desteklemiyor farkli isimlerle consttructuru tanimlamalisin
func newHumanParams(fname,lname string, age int) *Human{
	human:= new(Human);
	human.firstName = fname;
	human.lastName = lname;
	human.age = age;
	return human;
}

// CLASS YAPAISNA BENZER YAPIDA
func main(){
	
	// constructirsiz atama yapma
	x:= Human{firstName:"Sebahattin",lastName:"OZDEMIR"}

	fmt.Println(x.firstName, x.lastName)

	fmt.Println("/*************************************/")

	x.firstName = "Ali";
	x.lastName = "Ozdemir";

	fmt.Println(x.firstName,x.lastName)

	fmt.Println("/*************************************/")

	// construnctorla
	y:= newHuman();
	y.age = 25;
	y.firstName = "Sebahattin";
	y.lastName = "Ozdemir";

	fmt.Println(y.firstName,y.lastName,y.age)

	fmt.Println("/*************************************/")

	w:=newHumanParams("Melek","Sila",8);
	fmt.Println(w.firstName,w.lastName,w.age);
}