//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change
package main

import "fmt"

type securityTag bool

type Item struct {
	name string
	tag  securityTag
}

func activate(tag *securityTag) {
	*tag = true
}

func deactivate(tag *securityTag) {
	*tag = false
}

func checkout(items []Item) {
	for i := range items {
		deactivate(&items[i].tag)
	}
}

func main() {
	book := Item{"Book", true}
	pen := Item{"Pen", true}
	watch := Item{"Watch", true}
	laptop := Item{"Laptop", true}

	items := []Item{book, pen, watch, laptop}

	deactivate(&items[1].tag)

	fmt.Println(items)

	checkout(items)

	fmt.Println(items)
}
