//Go fundamentals Problemn sheet 3
//Author: Micheal Curley
//A program to return one of a pool of saved responses to user input
package main
//import packages
import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"math/rand"
	"time"
)
//random number generator method
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
//response method 
func ElizaResponse(input string)(string){
	//saved responses array
	responses := [4] string {
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		"Why don’t you tell me more about your father?"}
	//using regex methods to look for certain patterns
	match, err := regexp.MatchString("(?i).*\b[Ff]ather\b.*", input)
	//certain match returns certain response or error message
	if (match == true && err == nil){
		return responses[3]
	} else if(match == false && err == nil){#
		//check for "I am" in input string
		re := regexp.MustCompile("(?i)I am ([^.?!]*)[.?!]?")
		if matched := re.MatchString(input); matched {
			return re.ReplaceAllString(input, "How do you know you are $1?")
		} else {
			return responses[random(0, len(responses))]
		} 
	} else {
		return "Error message"
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