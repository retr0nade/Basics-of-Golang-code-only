package main

import (
	"fmt"
	"errors"
)

func check (sum int) (string, error){
	if sum < 0 {
		return "", fmt.Errorf("negative sum: %d", sum)
	}else{
		return "num is positive", nil
	}
}

func read(filename string) error{
	if filename == ""{
		return errors.New("filename cannot be empty")
	}
	return nil
}

func divide(x,y int) int{
	defer func (){
		if r := recover(); r != nil{
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if y == 0{
		panic ("div by 0")
	}
	return x/y
}

func main(){
	s, e := check(-1)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Success:", s)
	}

	err := read("")
	if err != nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File read successfully")
	}

	fmt.Println("Division result:", divide(10, 2)) // Normal case
	fmt.Println("Division result:", divide(10, 0)) // This will cause a panic
	fmt.Println("This line will not be executed due to panic") // will cause a panic
}