package main

import "fmt"

func main() {
	array := [...]string {"Muhammad", "Ridha", "Keren", "Ganteng", "Banged", "Kannnn"}

	//slice 1 := mengambil data ke 3 sampai 6
	slice1 := array[3:6]
	fmt.Printf("slice1 := array[3:6] adalah %v\n", slice1)


	//slice 2 := mengambil data pertama sampai ketiga
	slice2 := array[:3]
	fmt.Printf("slice2 := array[4:6] adalah %v\n", slice2)

	//slice 3 := mengambil data ke-empat sampai terkahir
	slice3 := array[3:]
	fmt.Printf("slice3 := array[4:6] adalah %v\n", slice3)


	//slice 4 := taking the whole array
	slice4 := array[:]
	fmt.Printf("slice4 := array[:] adalah %v\n", slice4)


	// membuat array baru yang isinya hari dalam seminggu
	days := [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}


	// membuat slice baru pada dua data terakhir dan mengubah isinya
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Printf("daySlice1 := %v\n", daySlice1)

	// membuat slice daySlice2 menggunakan append, disini melakukan append dengan 1 data baru dari daySlice1
	// diluar capasitas dari daySlice1 
	daySlice2 := append(daySlice1, "Hari Libur Baru euyyy") 
	daySlice2[0] = "ups"
	daySlice2[1] = "Minggu Lama"
	fmt.Printf("daySlice2 := %v\n", daySlice2) // ["ups", "Minggu Lama", "Hari Libur Baru euyyy"]
	fmt.Printf("daySlice1 := %v\n", daySlice1) // ["Sabtu Baru", "Minggu Baru"]
	fmt.Printf("days := %v\n", days)


	// make(array, length, capacity) method untuk membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Ridha"
  //newSlice[2] = "Keren" // error index [2] out of range because length is 2
	fmt.Printf("newSlice := %v\n", newSlice)


	// here we checking the lenght and capacity
	fmt.Printf("length of newSlice is %v\n", len(newSlice))
	fmt.Printf("capacity of newSlice is %v\n", cap(newSlice))

	// here we make new newSlice2 append with newSlice
	newSlice2 := append(newSlice, "Keren")
	fmt.Printf("newSlice2 := %v\n", newSlice2)
	fmt.Printf("len(newSlice2) := %v\n", len(newSlice2))
	fmt.Printf("cap(newSlice2) := %v\n", cap(newSlice2))

	newSlice2[0] = "Kan"
	fmt.Printf("%v\n", newSlice2)
	fmt.Printf("%v\n", newSlice)


	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	
//	fmt.Println(fromSlice)

	// copy isi dari fromSlice ke toSlice menggunakan method copy(destination slice, origin slice)
	copy(toSlice, fromSlice)
	fmt.Printf("fromSlice: %v\n", fromSlice)
	fmt.Printf("toSlice: %v\n", toSlice)

	iniArray := [...]int{1,2,3,4,5,6}
	iniSlice := []int{1,2,3,4,5,6}

	fmt.Printf("ini array %v\n", iniArray)
	fmt.Printf("ini slice %v\n", iniSlice)


}
