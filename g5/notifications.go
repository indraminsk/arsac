package main

import "fmt"

func sayHello(output ...string) {
	fmt.Printf("\n\n")

	fmt.Println(output[0])
	fmt.Printf("\n\n")

	for i := 1; i < len(output); i++ {
		fmt.Println(output[i])
	}

	fmt.Printf("\n\n")
}

func sayBye(output ...string) {
	fmt.Printf("\n\n")

	for i := 1; i < len(output); i++ {
		fmt.Println(output[i])
	}

	fmt.Printf("\n\n")
	fmt.Println(output[0])

	fmt.Printf("\n\n")
}
