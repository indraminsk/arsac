package main

import "fmt"

func sayHello(output ...string) {
	fmt.Printf("\n\n")

	if len(output) == 0 {
		output = append(output, "let's have a fun!")
	}

	fmt.Println(output[0])
	fmt.Printf("\n\n")

	for i := 1; i < len(output); i++ {
		fmt.Println(output[i])
	}

	fmt.Printf("\n\n")
}

func sayBye(output ...string) {
	fmt.Printf("\n\n")

	if len(output) == 0 {
		output = append(output, "bye! see you later")
	}

	for i := 1; i < len(output); i++ {
		fmt.Println(output[i])
	}

	fmt.Printf("\n\n")
	fmt.Println(output[0])

	fmt.Printf("\n\n")
}
