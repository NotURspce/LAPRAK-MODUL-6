package main
import "fmt"

func main() {
	var a,b int
	var j int

	fmt.Scan(&a, &b)
	for j = a; j <= b; j+=1 {
		fmt.Print(j, " ")
	}
}

//program cetak_rentang
//kamus
   //a, b, j : int
//Algoritma
   //input (a, b)
   //j <- a
   //while (j <= b) do
      //output (j)
      //j <- j + 1
   //endwhile
//endprogram