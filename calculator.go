package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) oparate(entry string, operator string) int {
	cleaned_entry := strings.Split(entry, operator)
	operator1 := parsear(cleaned_entry[0])
	operator2 := parsear(cleaned_entry[1])

	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println("Not supported")
		return 0

	}
}

func parsear(entry string) int {
	operator, _ := strconv.Atoi(entry)
	return operator
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entry := readEntry()
	operator := readEntry()
	c := calc{}
	fmt.Println(c.oparate(entry, operator))
}
