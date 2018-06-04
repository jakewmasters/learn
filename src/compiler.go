package main

// I want to try and compile and evaluate strict mathematical expressions

import (
	"bufio"
	"fmt"
	"os"
)

func lexer(input string){
	fmt.Println(input)
	for n := 0; n < len(input); n++ {
		// lexical analysis
		fmt.Println(input[n])
	}
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    lexer(text)
}