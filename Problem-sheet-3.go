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
	"strings"
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
	re := regexp.MustCompile("(?i).*\b[Ff]ather\b.*")
	if (re.MatchString(input) == true) {
		return responses[3]
	} else if (re.MatchString(input) == false) {
		//check for "I am", "I'm", "Iam", "i'm", "i am", "iam" in input string
		re := regexp.MustCompile("(?i)(I[' a]*m)([^.?!]*)[.?!]?")
		if (re.MatchString(input) == true) {
			return re.ReplaceAllString(input, "How do you know you are $2?")
		} else {
			return responses[random(0, len(responses))]
		} 
	} else {
		return "Error message"
	}
}
//Adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//Pronouns reflection method
func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	// List the reflections.
	reflections := [][]string{
		{"your", "my"},
		{"you", "I"},
		{"me", "you"},
	}
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	// Put the tokens back together.
	return strings.Join(tokens, ``)
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
	//calling reflect and elizaresponse  method
	response := Reflect(input)
	fmt.Println(ElizaResponse(response))
}