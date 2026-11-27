package main
import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	x, y, z := 4, 5, 6
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)

	// Menganti Nilai
	var jumlah int = 10
	jumlah = 20
	fmt.Println("Jumlah:", jumlah)
}