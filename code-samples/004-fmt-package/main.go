package main

import ("fmt")
var y = 42

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n",y) //%T print the type
	fmt.Printf("%b\n",y) // %b for binary 
	fmt.Printf("%x\n",y) // %x for hex
	fmt.Printf("%#x\n",y)  // %#x for binary

	y = 911
	fmt.Printf("%#x\n",y)
	fmt.Printf("%#x\t%b\t%x\n",y,y,y)

	s := fmt.Sprintf("%#x\t%b\t%x\n",y,y,y)
	fmt.Println(s)
	fmt.Printf("%v",y) // %v print the normal value
}

