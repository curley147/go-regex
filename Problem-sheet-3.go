//Go fundamentals Problemn sheet 3
//Author: Micheal Curley
//A program to return one of a pool of saved responses to user input
package main
//import packages
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
//response method 
func ElizaResponse(input string)(string){
	//saved responses
	str1 := "I’m not sure what you’re trying to say. Could you explain it to me?"
	str2 := "How does that make you feel?"
	str3 := "Why do you say that?"
	//using contains method to look for certain patterns
	if (strings.Contains(input, "father")==true){
		return str3
	} else if(strings.Contains(input, "mother")==true) {
		return str2
	} else {
		return str1
	}
}
//main function
func main() {
	fmt.Println("Say somethig to Eliza..")
	//initiate new scanner to take in user input
	scanner := bufio.NewScanner(os.Stdin)
	// use this to keep reading (including white space)
	scanner.Scan() 
    input := scanner.Text()
	fmt.Println("You entered", input)
	//calling response method
    fmt.Println(ElizaResponse(input))
}