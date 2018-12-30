package main

/*
	ini contoh komentar multiline
	pada pemograman golang
*/

import "fmt"

func main(){
	// ini contoh komentar inline GO LANG

	// manifest typing
	var FirtName string = "Nur"

	var	LastName string
	LastName = "Fajar"

	//type inference
	MiddleName := "Muhammad"

	_ =  "error"

	fmt.Println("Hallo %s %s!\n", FirtName, MiddleName, LastName)

	//type data decimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//type data boolean
	var exist bool = true

	fmt.Printf("exist? %t \n", exist)

	//type data strings
	var message string = "pesan"
	fmt.Printf("message: %s \n", message)

	//type data strings menggunakan backticks
	var pesan string = `Nama saya "Nur Fajar".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Printf("Pesan : %s \n", pesan)

	/*
		Default nilai dari type data :
		string => "",
		bool 	 => false,
		numerik non-desimal => 0,
		numerik desimal => 0.0

		Nilai nil adalah kosong, yg bisa diset nil :
		> Pointer
		> tipe data fungsi
		> slice
		> map
		> Channel
		> Interface kosong atau interface{}		
	*/



}
