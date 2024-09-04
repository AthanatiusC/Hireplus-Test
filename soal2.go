package main

import (
	"fmt"
	"math"
)

func hitungKembalian(totalBelanja, uangDibayar float64) (bool, int, map[int]int) {
	if uangDibayar < totalBelanja {
		return false, 0, nil
	}

	kembalian := int(math.Floor(uangDibayar-totalBelanja)/100) * 100

	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	pecahanMap := make(map[int]int)

	for _, p := range pecahan {
		for kembalian >= p {
			pecahanMap[p]++
			kembalian -= p
		}
	}

	return true, kembalian, pecahanMap
}

func main() {
	var totalBelanja, uangDibayar float64
	fmt.Println("Masukkan total belanja : ")
	fmt.Scanln(&totalBelanja)
	fmt.Println("Masukkan uang yang dibayarkan : ")
	fmt.Scanln(&uangDibayar)

	sukses, kembalian, pecahan := hitungKembalian(totalBelanja, uangDibayar)

	if !sukses {
		fmt.Println("False, kurang bayar")
	} else {
		fmt.Printf("Kembalian yang harus diberikan kasir: Rp %d\n", kembalian)
		fmt.Println("Pecahan uang:")
		for pecahan, jumlah := range pecahan {
			if jumlah > 0 {
				fmt.Printf("%d lembar %d\n", jumlah, pecahan)
			}
		}
	}
}
