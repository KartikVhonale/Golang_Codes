package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the any random number you wand dumbass :")
	// comma Ok || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for your random contribution :)|-( ")
	fmt.Println("you give this input : ", input)
}
