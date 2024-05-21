package main
import ("fmt"
		 "encoding/json")

//marshal ths []user to JSON.

type user struct{
	First string
	Age int
}

func main(){
	u1 := user{
		First:"James",
		Age:32,
	}
	u2 := user{
		First:"Moneypenny",
		Age:27,
	}

	people := []user{u1,u2}

	bs , err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}