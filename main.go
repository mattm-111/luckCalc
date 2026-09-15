package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		reset  = "\033[0m"
		red    = "\033[31m"
		green  = "\033[32m"
		yellow = "\033[33m"

		purple = "\033[35m"
	)

	replScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n\n\n\nWhat is your character's level?")
		fmt.Print(" => ")
		replScanner.Scan()
		if err := replScanner.Err(); err != nil {
			fmt.Printf("error in input => %v\n", replScanner.Err())
			continue
		}
		if len(replScanner.Text()) == 0 {
			continue
		}
		inputError := checkInput(replScanner.Text())
		if inputError != nil {
			fmt.Printf("Error: %v, please try again\n", inputError)
			continue
		}
		lvl, _ := strconv.ParseFloat(replScanner.Text(), 64)

		sunl := math.Ceil(lvl * 0.34)
		unl := math.Ceil(lvl * 0.67)
		lu := math.Ceil(lvl * 1.33)
		slu := math.Ceil(lvl * 1.66)

		fmt.Print("\n\n\n\n#############\nLuck Ranges:\n\n")
		fmt.Printf("%sSuper Unlucky: ≤ %v%s\n", red, sunl, reset)
		fmt.Printf("%sUnlucky: %v -- %v%s\n", yellow, sunl+1, unl, reset)
		fmt.Printf("Normal: %v -- %v\n", unl+1, lu-1)
		fmt.Printf("%sLucky: %v -- %v%s\n", green, lu, slu-1, reset)
		fmt.Printf("%sSuper Lucky: ≥ %v%s\n\n\n\n", purple, slu, reset)
		for {
			fmt.Print("Enter Y to calc again, or anything else to quit => ")
			replScanner.Scan()
			if err := replScanner.Err(); err != nil {
				fmt.Printf("error in input => %v\n", replScanner.Err())
				continue
			}
			if len(replScanner.Text()) == 0 {
				continue
			}
			if strings.ToLower(replScanner.Text()) == "y" {
				fmt.Print("\n\n\n\n")
				break
			} else {
				os.Exit(0)
			}
		}
		continue
	}

}

func checkInput(s string) error {
	if s == "0" {
		return errors.New("Level cannot be 0")
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return errors.New("Invalid level detected")
		}
	}
	return nil

}
