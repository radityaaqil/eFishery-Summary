package main

import (
	"efishery-summary/lib"
	"fmt"
)

func main() {
	//REMEMBER TO USE CAPITAL LETTER TO MAKE SOMETHING PUBLIC and vice versa

	//Pointer --> by reference - value will change based on the referenced variable
	var x int = 100
	var y *int = &x
	fmt.Println(y) //--> will display 0xc00001e098 as the memory address
	fmt.Println(*y) //--> will display 100 as the value inside the memory address
	x = 50
	fmt.Println(*y) //--> will display 50 as the value changed from 100 to 50 inside the memory address
	//==============================================================================================================================


	//Method --> function inside object or struct (unlike function method will always bound to an object or struct or whatever)
	//Immutable by default to change the value we need to provide pointer
	car1 := car{"Porsche", 2022}
	car1.sayHello()
	car1.changeBrand("Ferrari") //this will print Hello Ferrari (if the pointer is not provided on the method below, this would print Hello Porsche)
	car1.sayHello()
	//==============================================================================================================================


	//Create Library
	//go mod init on your main.go directory
	//create lib folder, and lib.go with package lib basically just use the same name for the file
	hello := lib.SayHello("Ferrari")
	fmt.Println(hello)
	//==============================================================================================================================


	//Interface (IDK MAN)
	petani1 := Petani{"Joe", 35}
	resultPetani := petani1.Hello()
	fmt.Println(resultPetani)
	//==============================================================================================================================

	//Go routines --> ASYNCRONOUS
	//Sequential, Concurrent, Parallel
}

type car struct{
	brand string
	year int
}

func (newCar car) sayHello() {
	fmt.Println("Hello", newCar.brand)
}

//We need to assign pointer to really change the struct's value
func (newCar *car) changeBrand(name string) {
	newCar.brand = name
}

//Interface
type Repository interface{
	GetDataPetani() string
	GetDataPetaniByID() string
}

type Petani struct{
	name string
	age int
}

func (p Petani) Hello() string {
	return "Hello, " + p.name
}