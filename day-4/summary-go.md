# GO

## Variable Declaration
    // var <var_name> <data_type>
	// var decimalNumber float64 = 2.62

	// var <var_name> --> var decimalNumber = 2.62
	
	//shorthand --> decimalNumber := 2.62

	//multivariables --> var x, y, z float64 | x = ... ; y = ... ; z = 9.02 ^float64 applies to all variables

	//multi assignment --> varA, varB := "nilai A", "nilai B" bisa juga untuk switch tanpa buat variabel baru varA, varB = "nilai B", "nilai A"

	// fmt.Println(decimalNumber)
	// fmt.Printf("Decimal number menggunakan formatting : %f\n", decimalNumber)
	//%f --> float
	//%t --> boolean

	// isExist := true

	// fmt.Printf("isExist : %t\n", isExist)

	// const constant string = "Prophet"
	
	//another way

	// const (
	// 	errorMessage = "error"
	// 	infoMessage = "info"
	// )

## Loops
	//===FOR RANGE===

	var buah = [2]string{"apel", "nanas"}

	for i, value := range buah {
		fmt.Println("index :", i, "value :", value)
	}

	//===FOR LOOP===

	for i := 1; i < len(buah)+1; i++ {
		if(i % 2 == 0){
			break
		}else{
			fmt.Println("index :", i, "value :", buah[i])
		}
	}

	fmt.Println(test(3,2))

	//Multiple return variables
	a, b := swap("A", "B")

	fmt.Println("This is A :", a)
	fmt.Println("This is B :", b)

	//Only one variable use underscore --> a, _ := swal("A", "B")

## Variadic

	//Variadic
	output := hello("A", "B", "C")
	fmt.Println(output)

	//Struct
	var school_1 school
	school_1.name = "SHS"
	school_1.address = "Malang"

	var student_1 student
	student_1.name = "Joe"
	student_1.age = 20
	student_1.school.name = school_1.name
	student_1.school.address = school_1.address


	fmt.Println("Name :", student_1.name)
	fmt.Println("Age :", student_1.age)
	fmt.Println("School :", student_1.school.name)

	//Combine Struct with slice
	students := []student{
		{name:"Joe", age:20, school: school{name:"JHS", address: "Malang"}},
		{name:"Black", age:22, school: school{name:"SHS", address: "Malang"}},
	}

	for _, v := range students {
		fmt.Println("Name :", v.name)
		fmt.Println("Age :", v.age)
		fmt.Println("School Name :", v.school.name)
		fmt.Println("School Address :", v.school.address)
	}
### Map
	//Map similar to object but each key or element needs to be unique
	var studentMap map[string]int
	studentMap = map[string]int{}

	studentMap["Student A"] = 100
	studentMap["Student B"] = 200
	studentMap["Student C"] = 300

	fmt.Println(studentMap)
### Array
	//Array --> the three dots indicate that this array is flexible in length (variadic)
	var animals = [...]string{"Cat", "Dog", "Dolphin", "Shark"}

	fmt.Println(animals[0])
### Slice
	//Slice --> similar to array but with pointer, length and capacity

	var fruits = []string{"Orange", "Mango", "Grape"} //--> length = 3, capacity = 3

	// partOfFruits := fruits[0:2]

	fruits = append(fruits, "Coconut") //--> similar to push in js

	fmt.Println("New Length :", len(fruits), "New Capacity :", cap(fruits)) //--> length = 4, capacity = 6 (3x2)

	//DEFER --> something that has "defer" will be executed last
	//multiple defer? the last defer will be executed first and so on until the first defer
	
	defer fmt.Println("This is defer")
	fmt.Println("This is not defer")


    func test(a int, b int) (int) {
        return a+b
    }

    func swap(a string, b string) (string, string) {
        return b, a
    }

## Variadic Function
    //Variadic parameter (spreading multiple parameter??)
    func hello(name ...string) (string) {
        msg := "Hello, "
        for _, v := range name {
            msg+=v
        }
        return msg
    }

## Struct
    //STRUCT short for STRUCTURE
    type student struct{
        name string
        age int
        school school
    }

    type school struct{
        name string
        address string
    }

## Library
    //Most used standard libraries (stdlib)
    //fmt, time, strings, strconv