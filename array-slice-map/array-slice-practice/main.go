package main

import (
	"fmt"
)

// ^ create a product struct
type Product struct {
	id    int64
	title string
	price float64
}

func main() {
	var hobbies [3]string = [3]string{"Web Development", "Reading", "Playing Games"}
	firstHobby := hobbies[0]
	excludeFirstHobby := hobbies[1:]
	mainHobbies := hobbies[:2]
	//^ storing parent array from the child slice
	parentHobbiesSlice := mainHobbies[1:3]

	//^ create a slice
	goals := []string{"Create a Rest API", "Learn Go for creating Real time web application."}
	goals[1] = "Learn to create a chat application in backend using channels."
	goals = append(goals, "Learn about new data structures.")

	products := []Product{
		{id: 1, title: "Tea", price: 10.89},
		{id: 2, title: "Power Bank", price: 30.23},
	}
	products = append(products, Product{id: 3, title: "Headsets", price: 89.67})

	fmt.Printf("\nHobbies: %v", hobbies)
	fmt.Printf("\nFirst Hobby: %v", firstHobby)
	fmt.Printf("\nExcluding First Hobby: %v", excludeFirstHobby)
	fmt.Printf("\nMain Hobbies: %v", mainHobbies)
	fmt.Printf("\nPrinting parent array's slice from the child slice: %v", parentHobbiesSlice)
	fmt.Printf("\nMy Goals are: %v", goals)
	fmt.Printf("\nThird product name: %v", products[2].title)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
