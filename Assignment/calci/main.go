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
	var opr string
	//	fmt.Println("printed from function", exp)
	//	parts = strings.Fields(exp)
	res, _ := strconv.ParseFloat(parts[0], 64)
	//var res float64
	for i := 1; i < len(parts) - 1; i=i+2 {
		opr = parts[i]
		nn, _ := strconv.ParseFloat(parts[i+1], 64)
		switch opr {
		case "+":
			res = res + nn
		case "-":
			res = res - nn
		case "*":
			res = res * nn
		case "/":
			res = res / nn
		default:
			fmt.Println("invalid operator")
		}

	}
	//fmt.Println("ans is :", res)
	return res

}

func main() {
	var res float64
	var exp string
	var lastchar string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter an expression (end with '='): ")
	exp, _ = reader.ReadString('\n')
	//exp = strings.TrimSpace(exp)
	lastchar=string(exp[len(exp)-1])
	if lastchar=="=" {
		res = computeexp(exp)
		fmt.Println(res)
	}

	
	res = computeexp(exp)
		//fmt.Println(exp)

		fmt.Println(res)

}
