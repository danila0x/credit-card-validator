package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func main() {
	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Printf("Не удалось загрузить банки: %v\n", err)
		return
	} else {
		fmt.Printf("Загружено банков: %d\n", len(banks))
	}
	cardNumber := "4000123456789017"
	valid := LuhnCheck(cardNumber)
	fmt.Printf("Валиден по Луне: %t\n", valid)
	if res := DetectBank(cardNumber, banks); res == nil {
		fmt.Println("Банк: не определён")
	} else {
		fmt.Printf("Банк: %s\n", res.Name)
	}
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var banks []Bank
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("неверный формат строки: %q", line)
		}
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		newBank := Bank{
			Name:   parts[0],
			Prefix: parts[1],
		}
		banks = append(banks, newBank)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении: %w", err)
	}
	return banks, nil
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
