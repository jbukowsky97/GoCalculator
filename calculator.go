package main

import (
	"math"
	"strconv"
	"fmt"
	"os"
)

func main() {
	var total float64 = 0.0
	for true {
		var operator string
		fmt.Println("Enter an operator (+ - * / ^)")
		_, err := fmt.Scanf("%s", &operator)
		fmt.Println("")
		if err != nil {
			fmt.Println("Error occurred")
			os.Exit(5)
		}
		if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "^" {
			if operator == "quit" || operator == "exit" {
				fmt.Println("Quitting...")
				os.Exit(0)
			}
			fmt.Println("Unkown operator\n")
			continue
		}
		var num float64
		fmt.Println("Enter a number")
		_, err = fmt.Scanf("%f", &num)
		fmt.Println("")
		if err != nil {
			fmt.Println("Error occurred")
			os.Exit(5)
		}
		if operator == "+" {
			total += num
		}else if operator == "-" {
			total -= num
		}else if operator == "*" {
			total *= num
		}else if operator == "/" {
			if num == 0.0 {
				fmt.Println("Cannot divide by 0\n")
				continue
			}
			total /= num
		}else if operator == "^" {
			total = math.Pow(total, num)
		}
		fmt.Println("Current Total:\t" + strconv.FormatFloat(total, 'f', -1, 64) + "\n")
	}
}
