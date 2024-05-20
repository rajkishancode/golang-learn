package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	fmt.Println("Hello world.") // this Println call below function to print internally
	fmt.Fprintln(os.Stdout,"Hello world.")
	io.WriteString(os.Stdout,"Hello world.")
}