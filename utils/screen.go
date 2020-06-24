package utils

import "fmt"

func MainScreen() {
	fmt.Println("===========================================")
	fmt.Println("===============ENIGMA SCHOOL===============")
	fmt.Println("===========================================")
	fmt.Println("==     1. Mahasiswa                      ==")
	fmt.Println("==     2. Mata Kuliah                    ==")
	fmt.Println("==     3. Dosen                          ==")
	fmt.Println("==     4. Input Nilai                    ==")
	fmt.Println("==     5. Laporan                        ==")
	fmt.Println("==     6. Exit                           ==")
	fmt.Println("===========================================")
	fmt.Printf("Masukkan Menu yang di pilih : ")
}

func MahasiswaScreen() {
	fmt.Println("==============================================")
	fmt.Println("===============ENIGMA SCHOOL==================")
	fmt.Println("==============================================")
	fmt.Println("==     1. Add Mahasiswa                     ==")
	fmt.Println("==     2. Delete Mahasiswa                  ==")
	fmt.Println("==     3. Detail Mahasiswa                  ==")
	fmt.Println("==     4. Ambil Mata Kuliah                 ==")
	fmt.Println("==     5. Detail Mata Kuliah yang di ambil  ==")
	fmt.Println("==     6. Update Mahasiswa                  ==")
	fmt.Println("==============================================")
	fmt.Printf("Masukkan Menu yang di pilih : ")
}

func DosenScreen() {
	fmt.Println("===========================================")
	fmt.Println("===============ENIGMA SCHOOL===============")
	fmt.Println("===========================================")
	fmt.Println("==     1. Add Dosen                      ==")
	fmt.Println("==     2. Delete Dosen                   ==")
	fmt.Println("==     3. Detail Dosen                   ==")
	fmt.Println("==     4. Update Dosen                   ==")
	fmt.Println("===========================================")
	fmt.Printf("Masukkan Menu yang di pilih : ")
}

func MatkulScreen() {
	fmt.Println("===========================================")
	fmt.Println("===============ENIGMA SCHOOL===============")
	fmt.Println("===========================================")
	fmt.Println("==     1. Add Mata Kuliah                ==")
	fmt.Println("==     2. Delete Mata Kuliah             ==")
	fmt.Println("==     3. Detail Mata Kuliah             ==")
	fmt.Println("==     4. Update Pengajar Mata Kuliah    ==")
	fmt.Println("===========================================")
	fmt.Printf("Masukkan Menu yang di pilih : ")
}

func LaporanScreen() {
	fmt.Println("===========================================")
	fmt.Println("===============ENIGMA SCHOOL===============")
	fmt.Println("===========================================")
	fmt.Println("==     1. Laporan Siswa Per Matkul       ==")
	fmt.Println("==     2. Laporan Rata-Rata Nilai        ==")
	fmt.Println("===========================================")
	fmt.Printf("Masukkan Menu yang di pilih : ")
}
