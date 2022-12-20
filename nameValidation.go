package main

import (
	"fmt"
	"strings"
)

func isValidName(name string ) bool {
	name = strings.TrimSpace(name)
	if len(name)>=5 && len(name)<=15{
		return true
	}
	return false
}

func main() {
	name := "vasu chauhan"
	if isValidName(name){
		fmt.Println(fmt.Sprintf("%s is a valid name",name))
	}else{
		fmt.Println(fmt.Sprintf("%s is not a valid name",name))
	}
}