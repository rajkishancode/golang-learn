package main
import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

func main(){
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword( []byte(s), bcrypt.MinCost) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword := "password123"
	err = bcrypt.CompareHashAndPassword(bs,  []byte(loginPassword)) 

	if err != nil {
		fmt.Println("You Can't Login!")
	}

	fmt.Println("You logged in!")
	
}