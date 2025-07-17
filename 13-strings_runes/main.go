package main

import "fmt"

func main(){
	 // strings and runes
	// A string is a sequence of bytes, while a rune is a sequence of Unicode code points.
	// In Go, a string is immutable, meaning it cannot be changed after creation.
	// A rune is an alias for int32 and represents a single Unicode code point.
	// Runes are used to represent characters in a string, especially for non-ASCII characters
	 s := "hello world 👋"
	 r := []rune(s) // Convert string to rune slice
	 fmt.Println("String:", s)
	 fmt.Println("Runes:", r)
	 fmt.Println("Length of string:", len(s)) // Length in bytes
	 fmt.Println("Length of runes:", len(r)) // Length in runes (characters


}