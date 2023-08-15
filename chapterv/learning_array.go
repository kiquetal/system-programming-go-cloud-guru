package main

import "fmt"

func main() {
	names := [3]string{"John", "Doe", "Smith"}
	for _, v := range names {
		fmt.Println(v)
	}
}
