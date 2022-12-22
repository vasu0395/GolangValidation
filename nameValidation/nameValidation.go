package main

import (
	"fmt"
	"strings"

	"github.com/forPelevin/gomoji"
)

// This function checks name validation based on lengths. ( here we're using name length [3,15])
func isValidNameBasedOnLength(name string ) bool {
	name = strings.TrimSpace(name)
	if len(name)>=3 && len(name)<=15{
		return true
	}
	return false
}

// This function return bool value i.e True when name contains emoji else False
func isEmojiPresent(name string) bool {
	res := gomoji.ContainsEmoji(name)
	return res
}

// This Function stops reserved name from request
func isNameContainsReservedWords(name string,mp map[string]int) bool {
	name = strings.TrimSpace(name)
	nameArray := strings.Split(name," ")
	for _,value := range nameArray{
		if _ , ok := mp[value]; ok {
			return true
		}
	}
	return false
}

func main() {
	// Samples for length checks..
	worstCaseSampleNameForLengthChecks := "qweritheweoewfjobewfvwoeveib"
	validSampleNameForLengthChecks := "vasu chauhan"
	
	if !isValidNameBasedOnLength(worstCaseSampleNameForLengthChecks){
		fmt.Println(worstCaseSampleNameForLengthChecks + " is not valid name based on length 😔.")
	}
	
	if isValidNameBasedOnLength(validSampleNameForLengthChecks){
		fmt.Println(validSampleNameForLengthChecks + " is valid name based on length 😊.")
	}

	// Samples for emoji checks..
	worstCaseSampleNameForEmojiChecks := "vasu 😊 chauhan"
	validSampleNameForEmojiChecks := "vasu"

	if isEmojiPresent(worstCaseSampleNameForEmojiChecks){
		fmt.Println(worstCaseSampleNameForEmojiChecks + " is not valid name as it contains emoji 😔.")
	}
	
	if !isEmojiPresent(validSampleNameForEmojiChecks){
		fmt.Println(validSampleNameForEmojiChecks + " is valid name 😊.")
	}

	reservedWords := map[string]int{
		"badWord" : 1,
		"foul" : 1,
	}

	name := "badWord foul animal"
	if isNameContainsReservedWords(name,reservedWords){
		fmt.Println(name + " is not valid name as it contains reserved words 😔.")

	}
}

// Sample outPut :- 
// qweritheweoewfjobewfvwoeveib is not valid name based on length 😔.
// vasu chauhan is valid name based on length 😊.
// vasu 😊 chauhan is not valid name as it contains emoji 😔.
// vasu is valid name 😊.