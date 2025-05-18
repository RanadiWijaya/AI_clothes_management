package main

import (
	"fmt"
	"strings"
)

type Pakaian struct {
	Nama              string
	Kategori          string
	Warna             string
	tingkatFormalitas int
}

var daftarPakaian []Pakaian

func main() {
	var pilihan string
	for {
		fmt.Println("\n=== AI Fashion Stylist ===") //Tampilan Awal
		fmt.Println("1. Tambah Pakaian")
		fmt.Println("2. Lihat Daftar Pakaian")
		fmt.Println("3. Cari Pakaian")
		fmt.Println("4. Urutkan Pakaian")
		fmt.Println("5. Rekomendasi Outfit")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		// Daftar Pilihan menu
		if pilihan == "1" {
			tambahPakaian()
		} else if pilihan == "2" {
			tampilkanDaftarPakaian()
		} else if pilihan == "3" {
			cariPakaian()
		} else if pilihan == "4" {
			urutkanPakaian()
		} else if pilihan == "5" {
			rekomendasiOutfit()
		} else if pilihan == "0" || pilihan == "keluar" {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Fungsi untuk menambahkan pakaian ke dalam daftar
// Fungsi ini meminta input dari pengguna untuk nama, kategori, warna, dan tingkat formalitas pakaian
func tambahPakaian() {
	var p Pakaian
	fmt.Print("Nama pakaian: ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Kategori (atas/bawah/luar/aksesoris): ")
	fmt.Scanln(&p.Kategori)
	fmt.Print("Warna: ")
	fmt.Scanln(&p.Warna)
	fmt.Print("Tingkat formalitas (1-5): ")
	fmt.Scanln(&p.tingkatFormalitas)

	daftarPakaian = append(daftarPakaian, p)
	fmt.Println("Pakaian berhasil ditambahkan!")
}

// Fungsi untuk menampilkan daftar pakaian
// Fungsi ini akan menampilkan semua pakaian yang ada dalam daftar
func tampilkanDaftarPakaian() {
	if len(daftarPakaian) == 0 {
		fmt.Println("Daftar pakaian kosong!")
		return
	}

	fmt.Println("\nDaftar Pakaian:")
	for i := 0; i < len(daftarPakaian); i++ {
		p := daftarPakaian[i]
		fmt.Printf("%d. %s (%s, %s) - Formalitas: %d\n",
			i+1, p.Nama, p.Kategori, p.Warna, p.tingkatFormalitas)
	}
}

// Fungsi untuk mencari pakaian berdasarkan warna atau kategori
// Fungsi ini meminta input dari pengguna untuk kriteria pencarian
func cariPakaian() {
	var berdasarkan, nilai string
	fmt.Print("Cari berdasarkan (warna/kategori): ")
	fmt.Scanln(&berdasarkan)
	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	var hasil []Pakaian

	// Pencarian berdasarkan warna
	if berdasarkan == "warna" {
		for i := 0; i < len(daftarPakaian); i++ {
			if strings.ToLower(daftarPakaian[i].Warna) == strings.ToLower(nilai) {
				hasil = append(hasil, daftarPakaian[i])
			}
		}
		// Pencarian berdasarkan kategori
	} else if berdasarkan == "kategori" {
		for i := 0; i < len(daftarPakaian); i++ {
			if strings.ToLower(daftarPakaian[i].Kategori) == strings.ToLower(nilai) {
				hasil = append(hasil, daftarPakaian[i])
			}
		}
		// selain "warna" dan "kategori" maka pilihan tidak valid
	} else {
		fmt.Println("Pilihan pencarian tidak valid!")
		return
	}

	// Jika tidak ada pakaian yang ditemukan, tampilkan pesan
	if len(hasil) == 0 {
		fmt.Println("Tidak ditemukan pakaian dengan kriteria tersebut")
		return
	}
	// jika ada pakaian yang ditemukan, tampilkan hasil pencarian
	fmt.Println("\nHasil Pencarian:")
	for i := 0; i < len(hasil); i++ {
		p := hasil[i]
		fmt.Printf("%d. %s (%s, %s) - Formalitas: %d\n",
			i+1, p.Nama, p.Kategori, p.Warna, p.tingkatFormalitas)
	}
}

// Fungsi untuk mengurutkan pakaian berdasarkan tingkat formalitas
// Fungsi ini menggunakan algoritma selection sort untuk mengurutkan pakaian
// berdasarkan tingkat formalitas dari yang paling rendah ke yang paling tinggi
func urutkanPakaian() {
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarPakaian[j].tingkatFormalitas < daftarPakaian[minIdx].tingkatFormalitas {
				minIdx = j
			}
		}
		daftarPakaian[i], daftarPakaian[minIdx] = daftarPakaian[minIdx], daftarPakaian[i]
	}
	fmt.Println("Pakaian telah diurutkan berdasarkan formalitas")
}

// Fungsi untuk memberikan rekomendasi outfit
// Fungsi ini akan merekomendasikan kombinasi pakaian atas dan bawah
func rekomendasiOutfit() {
	if len(daftarPakaian) < 2 {
		fmt.Println("Tidak cukup pakaian untuk membuat rekomendasi") // jika daftar pakaian kurang dari 2, maka tidak bisa membuat rekomendasi
		return
	}

	// Pisahkan pakaian atas dan bawah berdasarkan kategori
	// Kategori pakaian atas dan bawah
	var atas, bawah []Pakaian
	for i := 0; i < len(daftarPakaian); i++ {
		p := daftarPakaian[i]
		if p.Kategori == "atas" {
			atas = append(atas, p)
		} else if p.Kategori == "bawah" {
			bawah = append(bawah, p)
		}
	}

	// Jika tidak ada pakaian atas atau bawah, maka tidak bisa membuat rekomendasi
	if len(atas) == 0 || len(bawah) == 0 {
		fmt.Println("Tidak cukup pakaian untuk membuat rekomendasi")
		return
	}
	// Tampilkan rekomendasi outfit
	// Menggunakan loop untuk menampilkan 3 rekomendasi outfit
	fmt.Println("\nRekomendasi Outfit:")
	for i := 0; i < 3 && i < len(atas) && i < len(bawah); i++ {
		fmt.Printf("%d. %s (%s) dengan %s (%s)\n",
			i+1, atas[i].Nama, atas[i].Warna, bawah[i].Nama, bawah[i].Warna)
	}
}
