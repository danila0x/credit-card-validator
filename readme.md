# Credit Card Validator

## Description

This training project helps me practice the core concepts of the Go programming language. One of its key features is the implementation of the Luhn algorithm. The program reads data from a text file and checks whether the provided numbers (such as credit card numbers) pass the Luhn check, ensuring their validity according to the algorithm's rules.

## Opportunities

- Reading and processing data from a text file

- Validation of card numbers using the Luhn algorithm

- Error handling for file operations and data processing

- User input and input validation

## Installation and launch

1. Clone the repository:
`git clone https://github.com/danila0x/credit-card-validator.git`

2. Go to the project directory:
cd your-project

3. Prepare a text file with data (e.g., banks.txt) in the project folder

4. Run project:
`go run main.go`

## Usage examples

1. Add the txt file to the project directory, with the correct data
for the example:
```go
Lunar Bank,4000
Mars Credit Union,5000
Venus Express Bank,6000
```
2. Launch the program and start entering the card number

Luhn algorithm
```go
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
```

3. To exit the program, press Enter or insert an empty line.

## Project structure

| File | Description |
|------|-------------|
| `main.go` | The main program flow: asks for a card number, loads banks, validates input, and displays the result |
| `banks.txt` | A text file containing bank names and their BIN prefixes (e.g., `Venus Express Bank,4000`) |