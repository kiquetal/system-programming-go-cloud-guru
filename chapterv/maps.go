package main

import "fmt"

func main() {

	ages := map[string]int{}
	ages["Kevin"] = 18
	ages["John"] = 20
	ages["Smith"] = 30
	for k, v := range ages {
		switch v {
		case 18, 20:
			fmt.Println(fmt.Sprintf("key: %s, value: %d", k, v))
		default:
			fmt.Println(fmt.Sprintf("The key %s is not allowed", k))

		}

	}

}

func greeting(name, message string) (salution string) {
	salution = fmt.Sprintf("%s, %s", name, message)
	return
}
