package main

import "fmt"

func main() {
	var n, k, prioritas int

	fmt.Scan(&n)
	for k = 1; k <= n; k++ {
		prioritas = n - k + 1
		fmt.Println("Request", k, "diproses dengan prioritas", prioritas)
	}
}