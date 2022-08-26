package main

import "fmt"

func main() {
	fishArr := [10]fish{
		{"Benih Lele", 50000},
		{"Pakan lele cap menara", 25000},
		{"Probiotik A", 75000},
		{"Probiotik Nilai B", 10000},
		{"Pakan Nila", 20000},
		{"Benih Nila", 20000},
		{"Cupang", 5000},
		{"Benih Nila Super", 30000},
		{"Benih Cupang", 10000},
		{"Probiotik B", 10000},
	}
	//Harga termurah dan termahal
	outputMin, outputMax := findMinMax(fishArr[:])
	fmt.Println("Termurah :", outputMin.name, "-", outputMin.price, "\nTermahal :", outputMax.name, "-", outputMax.price)
	fmt.Println("====================================================")

	//Ikan dengan harga 10000
	result := findFishWithPrice(fishArr[:], 10000)
	fmt.Println("Ikan dengan harga tertentu :")
	for _, v := range result {
		fmt.Println("- "+v)
	}
	fmt.Println("====================================================")

	//Maksimum ikan dengan uang sekian
	output := findTotalFish(fishArr[:], 100000)
	fmt.Println("Jumlah maksimum ikan dengan jumlah uang tertentu :")
	fmt.Println(output)
}

type fish struct{
	name string
	price int
}

//Harga termurah dan termahal
//function accepts array of fish
func findMinMax (arr []fish) (fish, fish){
	//Sort Array then return the first and the last element
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var temp fish
			if (arr[i].price > arr[j].price) {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr[0], arr[len(arr)-1]
}

//Ikan dengan harga 10000
//function accepts array of fish and price
func findFishWithPrice (arr[] fish, price int) []string {
	//Find matching value(s)
	var result[]string
	for i := 0; i < len(arr); i++ {
		if (arr[i].price == price){
			result = append(result, arr[i].name)
		}
	}
	return result
}

//Maksimum ikan dengan uang 10000
//function accepts array of fish and price
func findTotalFish (arr[] fish, price int) int {
	//Sort Array
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var temp fish
			if (arr[i].price > arr[j].price) {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	//Count fish
	count, totPrice := 0, 0
	for i := 0; i < len(arr); i++ {
		if (totPrice < price) {
			totPrice+=arr[i].price
			count++
		}else{
			break
		}
	}
	return count
}