package main

import "fmt"

func main() {
	age := 17 // variable with a value

	fmt.Println(age) // printing the value

	fmt.Println(&age) // printing address of the value

	ptr := &age       // when you assign the address of a value to another variable
	fmt.Println(*ptr) // and print it using '*' it prints pack the original value
	fmt.Println(ptr)  // else it still prints the address of the value

	// you can also change the value of a variable using a pointer

	*ptr = 25        // use a pointer operator to chave the value of the variable and since 'ptr' already points to 'age' it the new variable will be printed
	fmt.Println(age) //Changing what the pointer points to changes the original variable.

	Count(age) // the value does not change becaue GO just copied the value and printed it
	fmt.Println(age)

	Change(&age)
	fmt.Println(age)
}

//using a pointer in a function

func Count(x int) {
	x = 100 // a value is assigned to a varriable but GO does not see it
}

func Change(x *int) {
	*x = 100 // but when a pointer is introduced the value changes because the function has the address instead of a copy.
}

// '&[value]' means give me the address of [value].
// '*[value]' means Go to the address stored in [value], and give me the value there.
