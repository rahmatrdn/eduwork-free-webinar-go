package main

import "fmt"

// TUJUAN:
// Struct digunakan untuk mendefinisikan objek data.
// Mempermudah dalam mengelompokkan data berdasarkan objeknya.

// Definisikan sebuah struktur bernama Person
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Membuat objek dari struktur Person
	person1 := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main Street",
	}

	// Mengakses properti dari objek Person
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Address:", person1.Address)

	// Mengubah nilai properti pada objek Person
	person1.Name = "Jane Smith"
	person1.Age = 25
	person1.Address = "456 Elm Street"

	// Menampilkan hasil perubahan
	fmt.Println("\nUpdated Information:")
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Address:", person1.Address)
}
