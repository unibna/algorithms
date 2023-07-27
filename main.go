package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const CHUNKSIZE = 1_000_000_000

type LargeNumber struct {
	foo  []string
	sign int // 1 is possitive, 0 is negative
}

func (ln *LargeNumber) fromString(input string) {
	ln.sign = 1
	if string(input[0]) == "-" {
		ln.sign = 0
		input = input[1:]
	}

	for i := 0; i < len(input); i++ {
		// digit, _ := strconv.Atoi(string(input[i]))
		// ln.value = append(ln.value, digit)
		ln.foo = append(ln.foo, string(input[i]))
	}
}

func (ln *LargeNumber) toString() string {
	var output string

	if !ln.isZero() {
		var removeZero bool = true
		for _, val := range ln.foo {

			if removeZero {
				if val != "0" {
					removeZero = false
				} else {
					continue
				}
			}

			output += val
		}

		if ln.sign == 0 {
			output = "-" + output
		}
	} else {
		output = "0"
	}

	return output
}

func (ln *LargeNumber) append(input int) {
	if input == 0 {
		for i := 0; i < 9; i++ {
			// ln.value = append([]int{0}, ln.value...)
			ln.foo = append([]string{"0"}, ln.foo...)
		}
	} else {
		for input > 0 {
			carry := input % 10
			input = input / 10
			// ln.value = append([]int{carry}, ln.value...)
			ln.foo = append([]string{fmt.Sprintf("%d", carry)}, ln.foo...)
		}
	}
}

func (ln *LargeNumber) random(size int) {
	var (
		input string
	)
	for i := 0; i < size; i++ {
		if i == 0 {
			input = fmt.Sprintf("%d", rand.Intn(8)+1) // 1-9
		} else {
			input += fmt.Sprintf("%d", rand.Intn(9)) // 0-9
		}
	}

	if rand.Intn(2) == 0 {
		input = "-" + input
	}

	ln.fromString(input)
}

func (ln LargeNumber) isZero() bool {
	var count int
	for _, x := range ln.foo {
		if x == "0" {
			count++
		}
	}

	if count == len(ln.foo) {
		return true
	}
	return false
}

func (ln LargeNumber) print() {
	fmt.Println(ln.toString())
}

func (ln LargeNumber) gtAbs(num LargeNumber) bool {
	// true meaning ln is greater, else is num or even
	if len(ln.foo) > len(num.foo) {
		return true
	} else if len(ln.foo) < len(num.foo) {
		return false
	}

	for i := 0; i < len(ln.foo); i++ {
		if ln.foo[i] > num.foo[i] {
			return true
		} else if ln.foo[i] > num.foo[i] {
			return false
		}
	}

	return false
}

func (num1 LargeNumber) add(num2 LargeNumber) (result LargeNumber) {
	if num1.sign != num2.sign {
		if num1.gtAbs(num2) {
			return subtract(num1, num2)
		}
		return subtract(num2, num1)
	}

	return handleAdd(num1, num2)
}

func handleAdd(num1 LargeNumber, num2 LargeNumber) (result LargeNumber) {
	var leftNum1, leftNum2, carry int
	result.sign = num1.sign
	rightNum1 := len(num1.foo)
	rightNum2 := len(num2.foo)

	for rightNum1 > 0 || rightNum2 > 0 {
		if rightNum1-9 > 0 {
			leftNum1 = rightNum1 - 9
		} else {
			leftNum1 = 0
		}

		if rightNum2-9 > 0 {
			leftNum2 = rightNum2 - 9
		} else {
			leftNum2 = 0
		}

		val1, _ := strconv.Atoi(strings.Join(num1.foo[leftNum1:rightNum1], ""))
		val2, _ := strconv.Atoi(strings.Join(num2.foo[leftNum2:rightNum2], ""))

		sum := val1 + val2 + carry
		val := sum % CHUNKSIZE
		carry = sum / CHUNKSIZE
		result.append(val)

		rightNum1 = leftNum1
		rightNum2 = leftNum2
	}

	return
}

func subtract(num1 LargeNumber, num2 LargeNumber) (result LargeNumber) {
	var leftNum1, leftNum2, borrow int
	result.sign = num1.sign
	rightNum1 := len(num1.foo)
	rightNum2 := len(num2.foo)

	for rightNum1 > 0 || rightNum2 > 0 {
		if rightNum1-9 > 0 {
			leftNum1 = rightNum1 - 9
		} else {
			leftNum1 = 0
		}

		if rightNum2-9 > 0 {
			leftNum2 = rightNum2 - 9
		} else {
			leftNum2 = 0
		}

		val1, _ := strconv.Atoi(strings.Join(num1.foo[leftNum1:rightNum1], ""))
		val2, _ := strconv.Atoi(strings.Join(num2.foo[leftNum2:rightNum2], ""))

		diff := val1 - val2 - borrow
		if diff < 0 {
			borrow = 1
			diff += CHUNKSIZE
		} else {
			borrow = 0
		}

		result.append(diff)

		rightNum1 = leftNum1
		rightNum2 = leftNum2
	}

	return
}

func main() {
	// Init numbers
	var number1, number2 LargeNumber

	// Set custom values
	number1.fromString("0")
	number2.fromString("0")

	// Random values
	// number1.random(1_000)
	// number2.random(1_000)

	// Caculate
	startTime := time.Now()
	result := number1.add(number2)
	endTime := time.Now()

	fmt.Println("Number 1:", number1.toString())
	fmt.Println("Number 2:", number2.toString())
	fmt.Println("Result:", result.toString())
	fmt.Println("Total time 2:", endTime.Sub(startTime))
}
