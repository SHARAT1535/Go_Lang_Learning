package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		fmt.Println("div by zero error")
		return 0
	}
	return a / b
}

func computeexp(exp string) float64 {
	parts := strings.Fields(exp)

	if len(parts) == 0 {
		return 0
	}

	res, _ := strconv.ParseFloat(parts[0], 64)

	for i := 1; i < len(parts)-1; i += 2 {
		opr := parts[i]

		nn, _ := strconv.ParseFloat(parts[i+1], 64)

		if opr == "+" {
			res = add(res, nn)
		} else if opr == "-" {
			res = sub(res, nn)
		} else if opr == "*" {
			res = mul(res, nn)
		} else if opr == "/" {
			res = div(res, nn)
		} else {
			fmt.Println("innvalid oprt:", opr)
			return 0
		}
	}

	fmt.Println("ans is:", res)
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var res float64

	for {
		fmt.Println("enter expr end with =")
		exp, _ := reader.ReadString('\n')
		exp = strings.TrimSpace(exp)

		if exp == "=" {
			break
		}

		res = computeexp(exp)
	}

	fmt.Println("final ans is :", res)
}
