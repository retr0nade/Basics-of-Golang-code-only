package main

import ("github.com/google/uuid"
	"fmt"
)

func main(){
	fmt.Println(uuid.New().String())
}