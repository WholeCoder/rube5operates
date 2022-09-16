package main

import (
	"fmt"
)

func main() {

	memory := []BitsetByte{}

	for i := 0; i < 10; i++ {
		memory = append(memory, InitNewByteset([]byte{}))
		fmt.Printf("#%V\n", memory[i])
	}

	fmt.Println("vim-go")
}
