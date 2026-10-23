package main
import "fmt"

func main() {
	var n, hasil, j int

	fmt.Scan(&n)
	hasil = 1
	for j = n; j >= 1; j-- {
		hasil = hasil * j
	}
	fmt.Println(hasil)
}