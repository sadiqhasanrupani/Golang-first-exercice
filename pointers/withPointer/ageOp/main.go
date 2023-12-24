package ageop

//^ avoid copy of variable using pointer example
func AddAge(age *int, value int) int {
	return *age + value
}

//^ mutate a variable using pointer example
func SubstractAge(age *int, value int) {
	*age = *age - value
}
