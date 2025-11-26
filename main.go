/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will demonstrate if nested statements 
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Variables
	var userName string = ""
	var student string = ""
	var userAge string = ""
	var userAgeNumber int = 0

	reader := bufio.NewReader(os.Stdin)

	// INPUT
	fmt.Print("Hello! Please state your name: ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Print("Okay, now are you a student? ")
	student, _ = reader.ReadString('\n')
	student = strings.TrimSpace(student)

	if student == "yes" || student == "Yes" {
		fmt.Print("Okay, now please state your age: ")
		userAge, _ = reader.ReadString('\n')
		userAge = strings.TrimSpace(userAge)
		userAgeNumber, _ = strconv.Atoi(userAge)

		if userAgeNumber >= 13 && userAgeNumber <= 19 {
			fmt.Println("Student " + userName + " is a teenager")
		} else if userAgeNumber <= 12 && userAgeNumber >= 5 {
			fmt.Println("Student " + userName + " is a child")
		}
	}

	if student == "no" || student == "No" {
		fmt.Print("Okay, now please state your age: ")
		userAge, _ = reader.ReadString('\n')
		userAge = strings.TrimSpace(userAge)
		userAgeNumber, _ = strconv.Atoi(userAge)

		if userAgeNumber >= 20 && userAgeNumber <= 30 {
			fmt.Println(userName + " is a young adult")
		} else {
			fmt.Println(userName + " is in a different life stage")
		}
	}

	fmt.Println("\nDone.")
}
