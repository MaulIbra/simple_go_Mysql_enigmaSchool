package views

import (
	"database/sql"
	"fmt"

	"github.com/maulibra/enigma_school/service"
	"github.com/maulibra/enigma_school/utils"
)

func AddMatkulView(db *sql.DB) {
	fmt.Println("=====================================================")
	fmt.Println("=                 ADD MATA KULIAH FORM              =")
	fmt.Println("=====================================================")
	fmt.Printf("Masukkan Nama Matakuliah :")
	namaMatkul := utils.Scanner()
	ReadDosenView(db)
	fmt.Printf("Masukkan Kode Dosen :")
	kodeDosen := utils.Scanner()
	isRegister, err := service.AddMatkul(db, namaMatkul, kodeDosen)
	if err != nil {
		fmt.Println(err)
	} else {
		if isRegister {
			fmt.Println("Add matkul Success")
		} else {
			fmt.Println("Add Matkul gagal")
		}
	}
	utils.Scanner()
}

func ReadMatkulView(db *sql.DB) {
	fmt.Println("================================================================================================")
	fmt.Println("=                                      DETAIL MATKUL                                           =")
	fmt.Println("================================================================================================")
	fmt.Printf("%-25s \t%-25s \t%-25s \t%-25s\n", "KODE MATKUL", "NAMA MATKUL", "KODE DOSEN", "NAMA DOSEN")
	listMatkul, err := service.ReadMatkul(db)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listMatkul {
			fmt.Printf("%-25s \t%-25s \t%-25s \t%-25s\n", val.KodeMatkul, val.NamaMatkul, val.Dosen.KodeDosen, val.Dosen.NamaDosen)
		}
	}
}

func ReadMatkulViewByKodeMatkul(db *sql.DB, kodeMatkul string) {
	fmt.Println("================================================================================================")
	fmt.Println("=                                       DETAIL MATKUL                                          =")
	fmt.Println("================================================================================================")
	fmt.Printf("%-25s \t%-25s \t%-25s \t%-25s \t%-10s\n", "NIM", "KODE MATKUL", "NAMA MATKUL", "KODE DOSEN", "NILAI")
	listMatkul, err := service.ReadMahasiswaMatkulByKode(db, kodeMatkul)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listMatkul {
			fmt.Printf("%-25s \t%-25s \t%-25s \t%-25s \t%-10s\n", val.Mahasiswa.Nim, val.MataKuliah.KodeMatkul, val.MataKuliah.NamaMatkul, val.MataKuliah.Dosen.KodeDosen, val.Nilai)
		}
	}
}

func UpdateMatkulView(db *sql.DB) {
	ReadMatkulView(db)
	fmt.Printf("Masukkan Kode Mata Kuliah :")
	kodeMatkul := utils.Scanner()
	ReadDosenView(db)
	fmt.Printf("Masukkan kode dosen yang akan menggantikan :")
	kodeDosen := utils.Scanner()
	isUpdate, err := service.UpdateMatkul(db, kodeMatkul, kodeDosen)
	if err != nil {
		fmt.Println(err)
	} else {
		if isUpdate {
			fmt.Println("Update Successfull")
		} else {
			fmt.Println("Update Gagal")
		}
	}
	utils.Scanner()
}

func DeleteMatkulView(db *sql.DB) {
	ReadMatkulView(db)
	fmt.Printf("Masukkan kode matkul yang ingin di hapus :")
	KodeMatkul := utils.Scanner()
	isDeleted, err := service.DeleteMatkul(db, KodeMatkul)
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
