package views

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maulibra/enigma_school/service"
	"github.com/maulibra/enigma_school/utils"
)

func AddMahasiswaView(db *sql.DB) {
	var jenisKelaminTemp, namaLengkap, alamat string
	fmt.Println("=====================================================")
	fmt.Println("=                 ADD MAHASISWA FORM                =")
	fmt.Println("=====================================================")
	for {
		fmt.Printf("Masukkan Nama Lengkap :")
		namaLengkap = utils.Scanner()
		if !utils.IsEmpty(namaLengkap) {
			break
		}
		fmt.Printf("Nama tidak boleh kosong\n")
	}

	for {
		fmt.Printf("Masukkan Jenis Kelamin ( 1 untuk laki-laki dan 0 untuk perempuan) :")
		jenisKelaminTemp = utils.Scanner()
		if utils.IsNumber(jenisKelaminTemp) {
			break
		}
		fmt.Printf("jenis kelamin harus berupa angka\n")
	}
	jenisKelamin, _ := strconv.Atoi(jenisKelaminTemp)
	for {
		fmt.Printf("Masukkan alamat :")
		alamat = utils.Scanner()
		if !utils.IsEmpty(alamat) {
			break
		}
		fmt.Printf("alamat tidak boleh kosong\n")
	}

	isRegister, err := service.AddMahasiswa(db, namaLengkap, alamat, jenisKelamin)
	if err != nil {
		fmt.Println(err)
	} else {
		if isRegister {
			fmt.Println("Registration Success")
		} else {
			fmt.Println("Registration gagal")
		}
	}
	utils.Scanner()
}

func ReadMahasiswa(db *sql.DB) {
	fmt.Println("=====================================================================================================================")
	fmt.Println("=                                               DETAIL MAHASISWA                                                    =")
	fmt.Println("=====================================================================================================================")
	fmt.Printf("%-25s \t%-20s \t%-10s \t%-20s\n", "NIM", "NAMA", "JENIS KELAMIN", "ALAMAT")
	listMahasiswa, err := service.ReadMahasiswa(db)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listMahasiswa {
			fmt.Printf("%-25s \t%-20s \t%-10d \t%-20s\n", val.Nim, val.NamaMahasiswa, val.JenisKelamin, val.Alamat)
		}
	}
}

func UpdateMahasiswa(db *sql.DB) {
	var nim, jenisKelaminTemp, namaLengkap, alamat string
	ReadMahasiswa(db)

	for {
		fmt.Printf("Masukkan Nim Mahasiswa yang ingin di update :")
		nim = utils.Scanner()
		if !utils.IsEmpty(nim) {
			break
		}
		fmt.Printf("Nim tidak boleh kosong\n")
	}

	for {
		fmt.Printf("Masukkan Nama Lengkap :")
		namaLengkap = utils.Scanner()
		if !utils.IsEmpty(namaLengkap) {
			break
		}
		fmt.Printf("Nama tidak boleh kosong\n")
	}

	for {
		fmt.Printf("Masukkan Jenis Kelamin ( 1 untuk laki-laki dan 0 untuk perempuan) :")
		jenisKelaminTemp = utils.Scanner()
		if utils.IsNumber(jenisKelaminTemp) {
			break
		}
		fmt.Printf("jenis kelamin harus berupa angka\n")
	}
	jenisKelamin, _ := strconv.Atoi(jenisKelaminTemp)

	for {
		fmt.Printf("Masukkan alamat :")
		alamat = utils.Scanner()
		if !utils.IsEmpty(alamat) {
			break
		}
		fmt.Printf("alamat tidak boleh kosong\n")
	}
	_, err := service.UpdateMahasiswa(db, nim, namaLengkap, alamat, jenisKelamin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Update Successfull")
	}
	utils.Scanner()
}

func DeleteMahasiswa(db *sql.DB) {
	var nim string
	ReadMahasiswa(db)
	for {
		fmt.Printf("Masukkan Nim Mahasiswa yang ingin di update :")
		nim = utils.Scanner()
		if !utils.IsEmpty(nim) {
			break
		}
		fmt.Printf("Nim tidak boleh kosong\n")
	}
	isDeleted, err := service.DeleteMahasiswa(db, nim)
	if err != nil {
		fmt.Println(err)
	} else {
		if isDeleted {
			fmt.Println("Delete Successfull")
		} else {
			fmt.Println("Delete Gagal")
		}
	}
	utils.Scanner()
}

func AmbilMatakuliah(db *sql.DB) {
	var nim, kodeMatkul string
	fmt.Println("=====================================================")
	fmt.Println("=              PENGAMBILAN MATA KULIAH FORM         =")
	fmt.Println("=====================================================")
	ReadMahasiswa(db)

	for {
		fmt.Printf("Masukkan nim mahasiswa yang ingin mengambil matkul :")
		nim = utils.Scanner()
		if !utils.IsEmpty(nim) {
			break
		}
		fmt.Printf("Nim tidak boleh kosong\n")
	}

	ReadMatkulView(db)

	for {
		fmt.Printf("masukkan kode matkul yang ingin di ambil :")
		kodeMatkul = utils.Scanner()
		if !utils.IsEmpty(kodeMatkul) {
			break
		}
		fmt.Printf("Nim tidak boleh kosong\n")
	}

	nilai := ""
	isAdded, err := service.AmbilMatakuliah(db, nim, kodeMatkul, nilai)
	if err != nil {
		fmt.Println(err)
	} else {
		if isAdded {
			fmt.Println("Pembambilan Mata kuliah berhasil")
		} else {
			fmt.Println("Pembambilan Mata kuliah berhasil")
		}
	}
	utils.Scanner()
}

func ListMatkulByNim(db *sql.DB) {
	ReadMahasiswa(db)
	var nim string
	for {
		fmt.Printf("Masukkan nim mahasiswa yang ingin di cari :")
		nim = utils.Scanner()
		if !utils.IsEmpty(nim) {
			break
		}
		fmt.Printf("Nim tidak boleh kosong\n")
	}

	fmt.Println("=====================================================================================================================")
	fmt.Println("=                                         PENGAMBILAN MATA KULIAH FORM                                              =")
	fmt.Println("=====================================================================================================================")
	fmt.Printf("%-25s \t%-20s \t%-25s \t%-5s\n", "KODE MATKUL", "NAMA MATKUL", "KODE DOSEN", "NILAI")
	listMatkul, err := service.ListMataKuliahByNim(db, nim)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listMatkul {
			fmt.Printf("%-25s \t%-20s \t%-25s \t%-5s\n", val.MataKuliah.KodeMatkul, val.MataKuliah.NamaMatkul, val.MataKuliah.Dosen.KodeDosen, val.Nilai)
		}
	}
	utils.Scanner()
}
