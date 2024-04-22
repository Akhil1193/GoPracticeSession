package main

import "fmt"

func main() {

	str := "MADAM"
	fmt.Println("Golang program to check palindrome,\n Given Word =", str)

	Palindrome(str)

	fmt.Printf("'%s' is palindrome\n", str)
}
func Palindrome(str string) bool {
	lastIdx := len(str) - 1
	// using for loop
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}
	return true
}
