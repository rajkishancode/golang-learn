//Handson exercise -3
//code - https://go.dev/play/p/BVRZTdlUZ_
// encode to JSON the []user sending the results to Stdout. Hint: you will
//need to use json.NewEncoder(os.Stdout).encode(v interface{})

//Reference 
/*
type Encoder
func NewEncoder(w io.Writer) *Encoder
func (enc *Encoder) Encode(v any) error
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("test",users) // it prints both go data and converted json data 

	// your code goes here
	//json.NewEncoder(os.Stdout) - this is *encoder and once we have *encoder we have Encode method [reference on top]
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Something went wrong and here's the error",err)
	}

	//it encoded our data(go data) straight to Stdout/straight from wire to json
	// go to json 
		
}


//Notes and Explaination
// -link-https://pkg.go.dev/os
// -link-https://pkg.go.dev/io
// -link-https://pkg.go.dev/encoding/json

/*

  1. when we go to /encoding/json

     - we have type Encoder down there 

       - type Encoder
		 func NewEncoder(w io.Writer) *Encoder 	 -->> and type encoder we have pointer to the type Encoder
		 func (enc *Encoder) Encode(v any) error -->> v means encode any type of data and return err if any



  2. we learn that os Stdout is pointer to a file because any value pointer to file has a method attach to it
     which make it a writer .
	 
	 for ex-
	 go to os link - there we get pointer to a file and multiple methods

		- func (f *File) Write(b []byte) (n int, err error) - take slice of byte and return int and err if any
		- and when we have this method we go to https://pkg.go.dev/io
		- we have type Writer
				type Writer interface {
					Write(p []byte) (n int, err error) - also take slice of byte and return int and err if any
				}

		- So if we have something type *file/pointer to file then that implements the writer interface

		- so wherever a writer is asked for, for example in NewEncoder method ,io.Writer we can put that in.

					type Encoder
						func NewEncoder(w io.Writer) *Encoder  

		- so in os Stdout

				Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")  				
				above code --> a Newfile and new file gives pointer to a file and that implements a writer interface

		- so we can take os Stdout where writer is asked for

				type Encoder
				func NewEncoder(w io.Writer) *Encoder -- like here io.writer is asked

				*/