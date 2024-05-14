package main

import "fmt"

// multi-dimensional slice
func main() {
	
	employees := []string{"Bob","John","Keneddy","Hazel"} // slice of strings
	fmt.Println(employees)
	
	countries := []string{"Japan","China","london","Brazil"}
	fmt.Println(countries)

	employeesAndCountries := [][]string{employees,countries} // this slice contains slice of strings 
	fmt.Println(employeesAndCountries)						 // so called multi-dimensional slice				


}
