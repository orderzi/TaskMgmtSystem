package utils

import (
	// "bufio"
	"fmt"
	"math"
	"regexp"
	"strings"

	// "os"
	// "regexp"
	"strconv"
	// "strings"
)



func ValidateAge(age string) (int, error) {
	a, err := strconv.Atoi(age)
	if err != nil {
		return 0, fmt.Errorf("field age cannot convert string '%v' to a number", age)
	}
	if a < 16 {
		return 0, fmt.Errorf("age '%d' is less than 16", a)
	}
	return a, nil
}

func RoundWeight(weight string) (int, error) {
	w, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert weight '%v' to a float", weight)
	}
	if math.Floor(w) == w {
		return int(w), nil
	}
	return int(math.Round(w)), nil
}

func FormatName(name string) (string, error) {
	name = strings.TrimSpace(name)
	reg, err := regexp.Compile("^[a-zA-Z]+$")
	if err != nil {
		return "", fmt.Errorf("error compiling regex", err)
	}
	if reg.MatchString(name) {
		return strings.Title(name), nil
	} else {
		return "", fmt.Errorf("name '%v' should contain only letters", name)
	}
}

// func GetUserName() string {
// 	var name string
// 	var reg *regexp.Regexp
// 	var err error
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("Enter a user name: ")
// 		name, _ = reader.ReadString('\n')
// 		name = strings.TrimSpace(name)
// 		reg, err = regexp.Compile("^[a-zA-Z]+$")
// 		if err != nil {
// 			fmt.Println("Error compiling regular expression:", err)
// 			continue
// 		}
// 		if reg.MatchString(name) {
// 			break
// 		} else {s
// 			fmt.Println("The name should only contain letters.")
// 		}
// 	}
// 	return strings.Title(name)
// }

// func GetUserAge() int {
// 	var age int
// 	var err error
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Println("Enter your age: ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		age, err = strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. please enter a number!")
// 		} else if age < 16 {
// 			fmt.Println("You must be at least 16 years old to use the app!")
// 		} else {
// 			break
// 		}
// 	}
// 	return age

// }

// func GetUserWeight() int {
// 	var weight int
// 	var err error
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Println("Enter your Weight: ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		weight, err = strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. please enter your weight")
// 		} else {
// 			break
// 		}

// 	}
// 	return weight

// }
