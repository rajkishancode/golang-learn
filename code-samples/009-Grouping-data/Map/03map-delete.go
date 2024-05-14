package main

import "fmt"


// delete(<map name>,"key")

func main(){

	m := map[string]int{   
		"James":32,
		"Miss Moneypenny":27,
	}

	fmt.Println(m)

	delete(m,"James") // deleting 
	fmt.Println(m)

	delete(m,"xyz") // can also delete something which is not present
	fmt.Println(m)
   

	//check before delete it exist or not
	if v,ok := m["Miss Moneypenny"];ok{ // this is comma okay idiom , as v= value comes in first 
			fmt.Println("value:",v)    // And if that value is ok it will go to below code block and run
			delete(m,"Miss Moneypenny")										
}
			fmt.Println(m) // now we have a empty map as we deleted Miss Moneypenny
}