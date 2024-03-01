package main

import (
	"fmt"
	"os"
	"strconv"
)

type biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var allBio = []biodata{
	{nama: "Rizky", alamat: "Jakarta", pekerjaan: "asd", alasan: "alasan rizky"},
	{nama: "Rama", alamat: "Jakarta", pekerjaan: "asd", alasan: "alasan rama"},
	{nama: "Dhani", alamat: "Jakarta", pekerjaan: "asd", alasan: "alasan dhani"},
}

func main() {
	input := os.Args

	index, err := strconv.Atoi(input[1])
	if err != nil || index < 1 || index > len(allBio) {
		fmt.Println("Invalid argument. Please provide a valid index (1 to", len(allBio), ") to show biodata.")
		return
	}

	showBiodata(index - 1)
}

func showBiodata(index int) {
	bio := allBio[index]
	fmt.Println("Nama:", bio.nama)
	fmt.Println("Alamat:", bio.alamat)
	fmt.Println("Pekerjaan:", bio.pekerjaan)
	fmt.Println("Alasan:", bio.alasan)
}
