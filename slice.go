package main

import "fmt"

func main() {
	// names := [...]string{"Aris", "Annelies", "Mahmudi", "Mellema", "Oppenheimer", "Christopher", "Nolan"}
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	var nSlice []string = make([]string, 2, 4)
	nSlice[0] = "Aris"
	nSlice[1] = "Mahmudi"
	nSlice2 := append(nSlice, "Mellema")
	nSlice2[0] = "Mavis"

	// slice1 := names[3:6]
	// slice2 := names[:3]
	// slice3 := names[4:]
	// slice4 := names[:]

	dSlice1 := days[5:]
	dSlice1[0] = "Pahing"
	dSlice2 := append(dSlice1, "Kliwon")
	dSlice2[1] = "Wage"

	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(slice4)
	fmt.Println(dSlice1)
	fmt.Println(dSlice2)
	fmt.Println(days)
	fmt.Println(nSlice)
	fmt.Println(len(nSlice))
	fmt.Println(cap(nSlice))

	fmt.Println(nSlice2)
	fmt.Println(len(nSlice2))
	fmt.Println(cap(nSlice2))
	fmt.Println(nSlice)
	fmt.Println(nSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	Arr := [...]int{1, 2, 3}
	Sli := []int{1, 2, 3}
	toSlice = append(toSlice, "Kliwon")

	fmt.Println(toSlice)
	fmt.Println(Arr[0])
	fmt.Println(Sli[0])
}
