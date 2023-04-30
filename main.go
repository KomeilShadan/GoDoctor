package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	doctorDialogue := doctor.Intro()

	fmt.Println(doctorDialogue)

	for {
		fmt.Print("---> ")
		input, _ := reader.ReadString('\n')
		input = stripCRLF(input)

		if input == "quit" {
			break
		} else {
			fmt.Print("Doctor: ")
			fmt.Println(doctor.Response(input))
		}
	}
}

func stripCRLF(input string) string {

	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}
