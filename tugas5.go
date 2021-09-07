package main

import "fmt"

func main() {

	var pilih, jumlah int

	fmt.Println("==============================")
	fmt.Println("1. Segitiga")
	fmt.Println("2. Segitiga Terbalik")
	fmt.Println("==============================")
	fmt.Print("Masukan jenis segitiga [1/2]: ")
	fmt.Scan(&pilih)
	fmt.Print("Masukan jumlah baris : ")
	fmt.Scan(&jumlah)

	if pilih == 1 {
		for i := 1; i <= jumlah; i++ {
			for j := jumlah; j >= i; j-- {
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ {
				fmt.Print("*")
			}
			for k := (i - 1); k >= 1; k-- {
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	} else {
		for i := 1; i <= jumlah; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print(" ")
			}
			for k := 0; k >= (i - jumlah); k-- {
				fmt.Print("*")
			}

			for k := 0; k < (jumlah - i); k++ {
				fmt.Print("*")
			}
			fmt.Print("\n")

		}
	}

}
