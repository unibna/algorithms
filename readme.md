## Problem
Given two numbers a and b, where:
	-10^1000 < a,b < 10^1000
Calculate a + b

## How to run
Access directory
```
go run main.go
```

To set custom numbers, let's use function fromString()
```
var number1, number2 LargeNumber
number1.fromString("123")
number2.fromString("-456")
```

To create random numbers with specific length, let's use function random()
```
var number LargeNumber
number.random(1_000) // a number have 1,000 digits
```
To add two numbers:
```
var number1, number2 LargeNumber
number1.fromString("123")
number2.fromString("-456")
result := number1.add(number2) // 123 + (-456)
```