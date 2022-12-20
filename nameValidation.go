package main

import (
	"fmt"
	"strings"
)

// this function checks name validation based on lengths.
func isValidName(name string ) bool {
	name = strings.TrimSpace(name)
	if len(name)>=5 && len(name)<=15{
		return true
	}
	return false
}

func print(name string) {
	if isValidName(name){
		fmt.Println(fmt.Sprintf("%s is a valid name",name))
	}else{
		fmt.Println(fmt.Sprintf("%s is not a valid name",name))
	}
}

func main() {
	print("vasu chauhan")
	print("invalidnameinvalidnameinvalidname")
}