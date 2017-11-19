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
	responses := [3] string {
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	//regexp strings to look for certain patterns
	reObjs := [6]string{
		//check for father regardless of lower/upper case
		"(?i).*\b[Ff]ather\b.*",
		//check for "I am", "I'm", "Iam", "i'm", "i am", "iam" in input string and capturing data
		"(?i)(I[' a]*m)([^.?!]*)[.?!]?",
		//check for "i think" and capturing data
		"(?i)(I think)([^.?!]*)[.?!]?",
		//check for you and ?
		"(?i)(you+)(.*)(\\?+)",
		//check for "i like" and capturing data
		"(?i)(I like)(\\s+)[.?!]?",
		//check for numbers
		"\\d",
	}
	//for loop iterating over regexp strings and compiling them
	for i:=0; i < len(reObjs); i++ {
		//test print to console
		//fmt.Println(reObjs[i])
		re := regexp.MustCompile(reObjs[i])
		match := re.MatchString(input)
		if (match == true) {
			switch i {
				case 0:
					//father response
					return "Why don’t you tell me more about your father?"
				case 1:
					//replacing start of response and adding captured data
					return re.ReplaceAllString(input, "How do you know you are $2?")
				case 2:
					return re.ReplaceAllString(input, "What makes you think that $2?")
				case 3:
					//talking about you not me response
					return "I thought we were talking about you, not me!"
				case 4:
					//I like response
					return re.ReplaceAllString(input, "Why is it that you like $2?")
				case 5:
					//numbers response
					return "I'm not a fan of numbers, can you simplify it in any way?"
				default:
					return "Error message"
			}
		} 
	}
	return responses[random(0, len(responses))]	
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