// pass_fail to check if student passed or failed

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	userinput := bufio.NewReader(os.Stdin)
	//returns multiple value, one return value and the other error value, can be ignored with _
	// or handled.
	input, err := userinput.ReadString('\n')
	// we trim space \n from the input
	input = strings.TrimSpace(input)
	// assign the converted string to float64 and assign it grade var
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)

	}
	// declaring a var out of scope of if the statements helps call them out of scope.
	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)

	//err is not named error cause error is an type in go and it would take precedence over error
	//type, it would same as declaring a variable fmt

}
