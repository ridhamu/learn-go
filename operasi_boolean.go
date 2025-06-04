package main

import "fmt"

func main() {
	var nilaiAkhir int32 = 90
	var absensi int32 = 81

	var apakahLulusNilaiAkir bool = nilaiAkhir > 80
	var apakahLulusAbsensi bool = absensi > 80

	var apakahLulus bool = apakahLulusNilaiAkir && apakahLulusAbsensi

	fmt.Println("apakah lulu? ", apakahLulus)
}
