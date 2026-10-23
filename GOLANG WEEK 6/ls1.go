package main
import "fmt"

func main() {
	var a, b, hasil, i int

	fmt.Scan(&a, &b)
	hasil = 1
	for i = 1; i <= b; i++ {
		hasil = hasil * a
	}
	fmt.Println(hasil)
}











//program pangkat
//kamus
   //a, b, hasil, i : int
//Algoritma
   //input (a, b)
   //hasil <- 1
   //i <- 1
   //for (i <= b) do
      //hasil <- hasil * a
      //i <- i + 1
   //endfor
   //output (hasil)
//endprogram