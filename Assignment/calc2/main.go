package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			res += nn
		} else if opr == "-" {
			res -= nn
		} else if opr == "*" {
			res *= nn
		} else if opr == "/" {
			if nn == 0 {
				fmt.Println("div by 0 error")
				return 0
			}
			res /= nn
		} else {
			fmt.Println("invalid opr", opr)
			return 0
		}
	}

	fmt.Println("ans is:", res)
	return res
}

func main() {
	var res float64
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("enter expr end with = ")
		exp, _ := reader.ReadString('\n')
		exp = strings.TrimSpace(exp)

		if exp == "=" {
			break
		}

		res = computeexp(exp)
	}

	fmt.Println("final ans is ", res)
}
