package main

import (
	"fmt"
	"os"
	"time"
)

var (
	Username = "Kiarra"
	Password = "2406357072"
	Histori  []string
)

type Pengguna struct {
	Nama        string
	NPM         string
	Universitas string
	Umur        string
	Alamat      string
	Genre       string
	Buku        string
}

type Buku struct {
	Nama   string
	Jumlah int
}

var DaftarBuku = []Buku{
	{Nama: "Pemograman", Jumlah: 10},
	{Nama: "Film", Jumlah: 5},
	{Nama: "Printing", Jumlah: 20},
}

func lihatInformasiPenggunaProgram() {
	data := Pengguna{
		Nama:   "Kiarra Seruni",
		NPM:    "2406357072",
		Umur:   "18",
		Alamat: "Depok",
		Genre:  "Action, Comedy, Psychological Horror",
		Buku:   "Auf wiedersehen, Sweetheart by GeorgeValier",
	}

	fmt.Println("======================================")
	fmt.Println("----- INFORMASI PENGGUNA PROGRAM -----")
	fmt.Println("======================================")
	fmt.Println("Nama: ", data.Nama)
	fmt.Println("NPM: ", data.NPM)
	fmt.Println("Umur: ", data.Umur)
	fmt.Println("Alamat: ", data.Alamat)
	fmt.Println("Genre buku yang disukai: ", data.Genre)
	fmt.Println("Buku favorit: ", data.Buku)
	fmt.Println("======================================")
}

func LihatDaftarBuku() {
	fmt.Println("========================================")
	fmt.Println("-- DAFTAR BUKU PERPUSTAKAAN VOKASI UI --")
	fmt.Println("======================================")
	for i, buku := range DaftarBuku {
		fmt.Printf("%v. Judul : %s - Jumlah : %v \n", i+1, buku.Nama, buku.Jumlah)
	}
	fmt.Println("========================================")
}

func TambahDaftarBuku() {
	fmt.Println("=====================================")
	fmt.Println("-------- TAMBAH DAFTAR BUKU ---------")
	fmt.Println("======================================")

	newBuku := Buku{}
	fmt.Println("Masukkan Nama Buku yang ingin ditambahkan: ")
	fmt.Scanln(&newBuku.Nama)
	fmt.Println("Masukkan Jumlah Buku yang ingin ditambahkan: ")
	fmt.Scanln(&newBuku.Jumlah)

	DaftarBuku = append(DaftarBuku, newBuku)
	fmt.Println("-----Buku berhasil ditambahkan!------")
	fmt.Println("=====================================")
}

func TambahPeminjamanBuku() {
	fmt.Println("=======================================")
	fmt.Println("------- TAMBAH PEMINJAMAN BUKU --------")
	fmt.Println("======================================")
	LihatDaftarBuku()

	var PilBuku, JmlBuku int
	fmt.Println("Masukkan Nomor Buku yang ingin dipinjam: ")
	fmt.Scanln(&PilBuku)
	fmt.Println("Masukkan Jumlah Buku yang ingin dipinjam: ")
	fmt.Scanln(&JmlBuku)

	if PilBuku < 1 || PilBuku > len(DaftarBuku) {
		fmt.Println("---Buku yang ingin dipinjam tidak ditemukan.---")
		fmt.Println("===============================================")
		return
	}

	if JmlBuku < 0 || JmlBuku == 0 {
		fmt.Println("----- Jumlah buku tidak bisa negatif/nol. -----")
		fmt.Println("===============================================")
		return
	}

	if JmlBuku > DaftarBuku[PilBuku-1].Jumlah {
		fmt.Println("Jumlah buku yang ingin dipinjam tidak mencukupi.")
		fmt.Println("===============================================")
		return
	}

	DaftarBuku[PilBuku-1].Jumlah -= JmlBuku
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	Histori = append(Histori, fmt.Sprintf("Meminjam buku berjudul %s sejumlah %v pcs pada tanggal %s", DaftarBuku[PilBuku-1].Nama, JmlBuku, currentTime))
	fmt.Println("-- Selamat! Peminjaman anda berhasil! --")
	fmt.Println("-Jangan lupa untuk dikembalikan ya! ^^-")
	fmt.Println("=======================================")
}

func HistoriPeminjamanBuku() {
	fmt.Println("======================================")
	fmt.Println("------ HISTORI PEMINJAMAN BUKU -------")
	fmt.Println("======================================")

	if len(Histori) == 0 {
		fmt.Println("----Kamu belum meminjam buku disini----")
	} else {
		for _, history := range Histori {
			fmt.Println(history)
		}
	}
	fmt.Println("======================================")
}

func main() {
	var inputUsername, inputPassword string

	fmt.Println("===========================================")
	fmt.Println("---SELAMAT DATANG DI PERPUSTAKAAN VOKASI---")

	fmt.Print("Silahkan input username anda: ")
	fmt.Scanf("%s\n", &inputUsername)

	fmt.Print("Silahkan input password anda: ")
	fmt.Scanf("%s\n", &inputPassword)

	fmt.Println("===========================================")

	if inputUsername != Username && inputPassword != Password {
		fmt.Println("password atau username anda, salah.")
		return
	}

	for {
		fmt.Println("Menu: ")
		fmt.Println("1. Lihat Informasi pengguna program")
		fmt.Println("2. Lihat daftar buku")
		fmt.Println("3. Tambah daftar buku")
		fmt.Println("4. Tambah peminjaman buku")
		fmt.Println("5. Histori peminjaman buku")
		fmt.Println("6. Keluar dari program")

		var choice int
		fmt.Println("===========================================================================")
		fmt.Print("---- Pilih menu:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			lihatInformasiPenggunaProgram()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			fmt.Println("Terima kasih telah menggunakan program peminjaman buku perpustakaan vokasi!")
			fmt.Println("-------Perbanyaklah membaca buku, karena buku adalah jendela dunia--------")
			fmt.Println("===========================================================================")
			os.Exit(0)
		default:
			fmt.Println("Pilihan anda tidak valid.")
		}
	}
}
