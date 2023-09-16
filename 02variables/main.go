package main

import "fmt"

/**
* ^ In PUBLIC we cannot define a no var style variables. You cannot write a variable without a var declaration like this
* & someVariable := validValue
* ^ But you can write a variable in PUBLIC using var.
 */

//^ PUBLIC
var someVariable = "A valid variable value"

/**
* ^ You can also define a constant variable by using the const keyword and it has it own way to create a constant variable.

* ^ You can see the example below,
 */

const ConstantVariable = "String variable"

func main() {
	var username string = "Sadiqhasan Rupani"

	fmt.Printf("\nMy name is: %s \nAnd The variable '%s' type is '%T'.\n", username, username, username)

	var isLoggedIn bool = true
	fmt.Printf("\nVariable of this type '%t' is '%T'.\n", isLoggedIn, isLoggedIn)

	//^ This variable data type is only available for 0 to 255 number.
	var smallInt uint8 = 255
	fmt.Printf("\nVariable of this type '%d' is '%T'.\n", smallInt, smallInt)

	//^ Now the floating values
	var smallFloat float32 = 255.234344
	fmt.Printf("\nVariable of this type '%f' is '%T'.\n", smallFloat, smallFloat)

	//^ default values and aliases
	var simpleInt int
	fmt.Printf("\nVariable of this type '%d' is '%T'.\n", simpleInt, simpleInt)

	//^ Implicit type
	var stringValue = "Sadiqhasan Rupani"
	fmt.Printf("\nVariable of this type '%s' is '%T'.\n", stringValue, stringValue)

	/**
	* ^ After declaring a variable without type but declare some value then it treats as it initial datatype value. And that datatype will never changes because "GO" is a static type language.

	* ^ Their is another way to declare variable by using this syntax
	* & variableName := value
	 */

	//^ no var style
	anotherNumber := 1000
	fmt.Printf("\nVariable of this type '%d' is '%T'.\n", anotherNumber, anotherNumber)

	//^ global constant variable print line.
	fmt.Printf("\nVariable of this type '%s' is '%T'.\n", ConstantVariable, ConstantVariable)

}
