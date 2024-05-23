package main

import (
	"fmt"
	"runtime"
)
//Hands-on-Exercise #6

/*
Create a Program that Prints out your OS and ARCH . Use the following command to run it
	- go run
	- go build - will create a executable file same folder
	- go install - will move exe file to C:\Users\Raj\Documents\goworkspace\bin
				 - go to the bin path from vscode terminal 
				 - just type file name ex- 06-os-arch > it will run .
*/
func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
}
