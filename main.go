package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"},
		{Name: "Venus Express", Prefix: "6000"},
		{Name: "Saturn Ring", Prefix: "7000"},
	}
	cardNumber := "4000123456789017"
	if res := DetectBank(cardNumber, banks); res == nil {
		fmt.Println("не определён")
	} else {
		fmt.Println(res)
	}
	fmt.Println(LuhnCheck(cardNumber))

}

func LuhnCheck(cardNumber string) bool {
	if cardNumber == "" {
		return false
	}
	sum := 0
	isSecond := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isSecond = !isSecond
	}
	return sum%10 == 0
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if cardNumber == "" {
		return nil
	}
	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}
	return nil
}
