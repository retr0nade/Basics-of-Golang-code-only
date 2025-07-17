package main

import "fmt"

func main(){

	//loops
	sum := 0
	for i:=0; i<5; i++{
		sum += i
	}
	fmt.Println("Sum is:", sum)

	j := 5
	sum = 0
	for j<= 10{
		sum += j
		j++
	}
	fmt.Println("Sum is:", sum)

	num := []int{1, 2, 3, 4, 5}
	for i, v := range num{
		fmt.Println("Index: ", i, "Value: ", v,)
	}


	//if else
	a := 10
	if a > 5{
		fmt.Println("a is greater than 5")
	}else if a == 5{
		fmt.Println("a is equal to 5")
	}else{
		fmt.Println("a is less than 5")
	}

	//switch case
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

	//type switch
	var x interface{} = "Hello"
	switch x.(type){
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of unknown type")
	}

	//multiple cases
	grade := 85
	switch grade{
	case 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100:
		fmt.Println("Grade is A")
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
		fmt.Println("Grade is B")
	case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79:
		fmt.Println("Grade is C")
	case 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
		fmt.Println("Grade is D")
	default:
		fmt.Println("Grade is F")
	}

}