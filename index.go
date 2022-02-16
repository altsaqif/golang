package main

import "fmt"

func main() {
	fmt.Println("halo")
	
	var nama5 = [5]string{
		"AlTsaqif", 
		"Philips", 
		"Ucup",
		"Nugi",
		"Indra",
	}
	nama5[2] = "Azka"
	fmt.Println(nama5[4])
	fmt.Println(nama5[2] == "Azka")
	fmt.Println(nama5[2])
	fmt.Println(nama5[0:4])
	fmt.Println(nama5[0:])
	fmt.Println(nama5[:4])
	fmt.Println(nama5[:])
}
