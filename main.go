package main

import (
	"fmt"
	"strings"
)

type Pakaian struct {
	Tipe              string
	Kategori          string
	Warna             string
	tingkatFormalitas int
}

const MAX_PAKAIAN = 100

func main() {
	var daftarPakaian [MAX_PAKAIAN]Pakaian
	var jumlahPakaian int
	var pilihan string

	jumlahPakaian = 0

	for {
		tampilkanMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == "1" {
			jumlahPakaian = tambahPakaian(&daftarPakaian, jumlahPakaian)
		} else if pilihan == "2" {
			tampilkanDaftarPakaian(daftarPakaian, jumlahPakaian)
		} else if pilihan == "3" {
			cariPakaian(daftarPakaian, jumlahPakaian)
		} else if pilihan == "4" {
			urutkanPakaian(&daftarPakaian, jumlahPakaian)
		} else if pilihan == "5" {
			rekomendasiOutfit(daftarPakaian, jumlahPakaian)
		} else if pilihan == "0" || pilihan == "keluar" {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Procedure untuk menampilkan menu utama
func tampilkanMenu() {
	fmt.Println("\n=== AI Fashion Stylist ===")
	fmt.Println("1. Tambah Pakaian")
	fmt.Println("2. Lihat Daftar Pakaian")
	fmt.Println("3. Cari Pakaian")
	fmt.Println("4. Urutkan Pakaian")
	fmt.Println("5. Rekomendasi Outfit")
	fmt.Println("0. Keluar")
}

// Function untuk menambahkan pakaian ke dalam daftar
// Parameter: pointer array pakaian dan jumlah pakaian saat ini
// Return: jumlah pakaian setelah penambahan
func tambahPakaian(daftarPakaian *[MAX_PAKAIAN]Pakaian, jumlahPakaian int) int {
	var p Pakaian

	if jumlahPakaian >= MAX_PAKAIAN {
		fmt.Println("Daftar pakaian sudah penuh!")
		return jumlahPakaian
	}

	p = inputDataPakaian()

	if validasiPakaian(p) {
		daftarPakaian[jumlahPakaian] = p
		jumlahPakaian = jumlahPakaian + 1
		fmt.Println("Pakaian berhasil ditambahkan!")
	} else {
		fmt.Println("Data pakaian tidak valid!")
	}

	return jumlahPakaian
}

// Function untuk input data pakaian dari user
// Return: struct Pakaian dengan data yang diinput user
func inputDataPakaian() Pakaian {
	var p Pakaian

	fmt.Print("Tipe pakaian: ")
	fmt.Scanln(&p.Tipe)
	fmt.Print("Kategori (atas/bawah/luar/aksesoris): ")
	fmt.Scanln(&p.Kategori)
	fmt.Print("Warna: ")
	fmt.Scanln(&p.Warna)
	fmt.Print("Tingkat formalitas (1-5): ")
	fmt.Scanln(&p.tingkatFormalitas)

	return p
}

// Function untuk validasi data pakaian
// Parameter: struct Pakaian yang akan divalidasi
// Return: true jika valid, false jika tidak valid
func validasiPakaian(p Pakaian) bool {
	if p.Tipe == "" || p.Kategori == "" || p.Warna == "" {
		return false
	}

	if p.tingkatFormalitas < 1 || p.tingkatFormalitas > 5 {
		return false
	}

	kategori := strings.ToLower(p.Kategori)
	if kategori == "atas" || kategori == "bawah" || kategori == "luar" || kategori == "aksesoris" {
		return true
	}

	return false
}

// Procedure untuk menampilkan daftar pakaian
// Parameter: array pakaian dan jumlah pakaian
func tampilkanDaftarPakaian(daftarPakaian [MAX_PAKAIAN]Pakaian, jumlahPakaian int) {
	if jumlahPakaian == 0 {
		fmt.Println("Daftar pakaian kosong!")
		return
	}

	fmt.Println("\nDaftar Pakaian:")
	tampilkanListPakaian(daftarPakaian, jumlahPakaian)
}

// Procedure untuk menampilkan list pakaian dengan format tertentu
// Parameter: array pakaian dan jumlah pakaian
func tampilkanListPakaian(pakaianList [MAX_PAKAIAN]Pakaian, jumlah int) {
	var i int
	for i = 0; i < jumlah; i = i + 1 {
		p := pakaianList[i]
		fmt.Printf("%d. %s (%s, %s) - Formalitas: %d\n",
			i+1, p.Tipe, p.Kategori, p.Warna, p.tingkatFormalitas)
	}
}

// Procedure untuk mencari pakaian berdasarkan kriteria tertentu
// Parameter: array pakaian dan jumlah pakaian
func cariPakaian(daftarPakaian [MAX_PAKAIAN]Pakaian, jumlahPakaian int) {
	var hasil [MAX_PAKAIAN]Pakaian
	var jumlahHasil int

	if jumlahPakaian == 0 {
		fmt.Println("Daftar pakaian kosong!")
		return
	}

	kriteria, nilai := inputKriteriaPencarian()
	jumlahHasil = pencarianPakaian(daftarPakaian, jumlahPakaian, kriteria, nilai, &hasil)
	tampilkanHasilPencarian(hasil, jumlahHasil)
}

// Function untuk input kriteria pencarian dari user
// Return: kriteria pencarian dan nilai yang dicari
func inputKriteriaPencarian() (string, string) {
	var berdasarkan, nilai string

	fmt.Print("Cari berdasarkan (warna): ")
	fmt.Scanln(&berdasarkan)
	fmt.Print("Masukkan nilai (berdasarkan tingkat Formalitas): ")
	fmt.Scanln(&nilai)

	return berdasarkan, nilai
}

// Function untuk melakukan pencarian pakaian (Algoritma Sequential Search)
// Parameter: array pakaian, jumlah pakaian, kriteria, nilai, dan pointer array hasil
// Return: jumlah pakaian yang ditemukan
func pencarianPakaian(daftarPakaian [MAX_PAKAIAN]Pakaian, jumlahPakaian int, kriteria string, nilai string, hasil *[MAX_PAKAIAN]Pakaian) int {
	var jumlahHasil int
	var i int

	jumlahHasil = 0

	for i = 0; i < jumlahPakaian; i = i + 1 {
		pakaian := daftarPakaian[i]

		if kriteria == "warna" && strings.ToLower(pakaian.Warna) == strings.ToLower(nilai) {
			hasil[jumlahHasil] = pakaian
			jumlahHasil = jumlahHasil + 1
		} else if kriteria == "kategori" && strings.ToLower(pakaian.Kategori) == strings.ToLower(nilai) {
			hasil[jumlahHasil] = pakaian
			jumlahHasil = jumlahHasil + 1
		}
	}

	return jumlahHasil
}

// Procedure untuk menampilkan hasil pencarian
// Parameter: array hasil pencarian dan jumlah hasil
func tampilkanHasilPencarian(hasil [MAX_PAKAIAN]Pakaian, jumlahHasil int) {
	if jumlahHasil == 0 {
		fmt.Println("Tidak ditemukan pakaian dengan kriteria tersebut")
		return
	}

	fmt.Println("\nHasil Pencarian:")
	tampilkanListPakaian(hasil, jumlahHasil)
}

// Procedure untuk mengurutkan pakaian berdasarkan tingkat formalitas (Selection Sort)
// Parameter: pointer array pakaian dan jumlah pakaian
func urutkanPakaian(daftarPakaian *[MAX_PAKAIAN]Pakaian, jumlahPakaian int) {
	if jumlahPakaian <= 1 {
		fmt.Println("Tidak cukup pakaian untuk diurutkan")
		return
	}

	// Algoritma Selection Sort
	selectionSort(daftarPakaian, jumlahPakaian)

	fmt.Println("Pakaian telah diurutkan berdasarkan formalitas")
}

// Procedure untuk implementasi Selection Sort
// Parameter: pointer array pakaian dan jumlah pakaian
func selectionSort(pakaian *[MAX_PAKAIAN]Pakaian, n int) {
	var i, j, minIdx int

	for i = 0; i < n-1; i = i + 1 {
		minIdx = i

		for j = i + 1; j < n; j = j + 1 {
			if pakaian[j].tingkatFormalitas < pakaian[minIdx].tingkatFormalitas {
				minIdx = j
			}
		}

		// Swap elements
		temp := pakaian[i]
		pakaian[i] = pakaian[minIdx]
		pakaian[minIdx] = temp
	}
}

// Procedure untuk memberikan rekomendasi outfit
// Parameter: array pakaian dan jumlah pakaian
func rekomendasiOutfit(daftarPakaian [MAX_PAKAIAN]Pakaian, jumlahPakaian int) {
	var atas [MAX_PAKAIAN]Pakaian
	var bawah [MAX_PAKAIAN]Pakaian
	var jumlahAtas, jumlahBawah int

	if jumlahPakaian < 2 {
		fmt.Println("Tidak cukup pakaian untuk membuat rekomendasi")
		return
	}

	jumlahAtas, jumlahBawah = pisahkanKategoriPakaian(daftarPakaian, jumlahPakaian, &atas, &bawah)

	if jumlahAtas == 0 || jumlahBawah == 0 {
		fmt.Println("Tidak cukup pakaian untuk membuat rekomendasi")
		return
	}

	tampilkanRekomendasi(atas, bawah, jumlahAtas, jumlahBawah)
}

// Function untuk memisahkan pakaian berdasarkan kategori atas dan bawah
// Parameter: array pakaian lengkap, jumlah pakaian, pointer array atas, pointer array bawah
// Return: jumlah pakaian atas dan jumlah pakaian bawah
func pisahkanKategoriPakaian(daftarPakaian [MAX_PAKAIAN]Pakaian, jumlahPakaian int, atas *[MAX_PAKAIAN]Pakaian, bawah *[MAX_PAKAIAN]Pakaian) (int, int) {
	var jumlahAtas, jumlahBawah int
	var i int

	jumlahAtas = 0
	jumlahBawah = 0

	for i = 0; i < jumlahPakaian; i++ {
		p := daftarPakaian[i]
		if strings.ToLower(p.Kategori) == "atas" {
			atas[jumlahAtas] = p
			jumlahAtas = jumlahAtas + 1
		} else if strings.ToLower(p.Kategori) == "bawah" {
			bawah[jumlahBawah] = p
			jumlahBawah = jumlahBawah + 1
		}
	}

	return jumlahAtas, jumlahBawah
}

// Procedure untuk menampilkan rekomendasi outfit
// Parameter: array pakaian atas, array pakaian bawah, jumlah atas, jumlah bawah
func tampilkanRekomendasi(atas [MAX_PAKAIAN]Pakaian, bawah [MAX_PAKAIAN]Pakaian, jumlahAtas int, jumlahBawah int) {
	var maksimalRekomendasi int
	var i int

	fmt.Println("\nRekomendasi Outfit:")

	maksimalRekomendasi = hitungMaksimalRekomendasi(jumlahAtas, jumlahBawah)

	for i = 0; i < maksimalRekomendasi; i++ {
		fmt.Printf("%d. %s (%s) dengan %s (%s)\n",
			i+1, atas[i].Tipe, atas[i].Warna, bawah[i].Tipe, bawah[i].Warna)
	}
}

// Function untuk menghitung maksimal rekomendasi yang bisa ditampilkan
// Parameter: jumlah pakaian atas dan jumlah pakaian bawah
// Return: jumlah maksimal rekomendasi (maksimal 3)
func hitungMaksimalRekomendasi(jumlahAtas int, jumlahBawah int) int {
	var maksimal int

	maksimal = 3

	if jumlahAtas < maksimal {
		maksimal = jumlahAtas
	}

	if jumlahBawah < maksimal {
		maksimal = jumlahBawah
	}

	return maksimal
}
