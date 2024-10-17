package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan empat bilangan (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	perAC := permutation(a, c)
	comAC := combination(a, c)
	perBD := permutation(b, d)
	comBD := combination(b, d)

	fmt.Println(perAC, comAC)
	fmt.Println(perBD, comBD)
}
