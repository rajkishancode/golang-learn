package main
import "fmt"
//syntax for function ->
//func (r receiver) identifier(parameters) (return(s)) {code}
func main(){
	foo()
	bar("James")
	s:= woo("Moneypenny");
	fmt.Println(s);
	x,y := fullName("Ian","Flemings")
	fmt.Println(x)
	fmt.Println(y)
}

func foo(){
	fmt.Println("Hello from foo")
}
//EVERYTHING IN GO IS PASS BY VALUE
func bar(s string){ // takes an argument
  fmt.Println("Hello from",s)
}

// function return a string
func woo(s string)string{
	return fmt.Sprint("Hello from ,",s);  //Sprint is string print
}

//function with multiple returns
func fullName(first string,last string) (string, bool) {
	a := fmt.Sprint(first,"  ",last ,"  ",`say have a nice day!`);
	b := false
	return a,b
}

