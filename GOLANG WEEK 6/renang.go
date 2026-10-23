package main
import "fmt"

func main() {
	var n, j, x, totalJarak int

	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		fmt.Scan(&x)
		totalJarak = x * 5 
		fmt.Println(totalJarak)
	}
}