package views

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maulibra/enigma_school/service"
	"github.com/maulibra/enigma_school/utils"
)

func AddDosenView(db *sql.DB) {
	fmt.Println("=====================================================")
	fmt.Println("=                 ADD DOSEN FORM                    =")
	fmt.Println("=====================================================")
	fmt.Printf("Masukkan Nama Dosen :")
	namaDosen := utils.Scanner()
	isRegister, err := service.AddDosen(db, namaDosen)
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

func ReadDosenView(db *sql.DB) {
	fmt.Println("================================================")
	fmt.Println("=              DETAIL DOSEN                    =")
	fmt.Println("================================================")
	fmt.Printf("%-25s \t%-25s\n", "KODE DOSEN", "NAMA DOSEN")
	listDosen, err := service.ReadDosen(db)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listDosen {
			fmt.Printf("%-25s \t%-25s\n", val.KodeDosen, val.NamaDosen)
		}
	}
}

func UpdateDosenView(db *sql.DB) {
	ReadDosenView(db)
	fmt.Printf("Masukkan Kode Dosen yang ingin di update :")
	kodeDosen := utils.Scanner()
	fmt.Printf("Masukkan Nama Dosen :")
	namaDosen := utils.Scanner()
	isUpdate, err := service.UpdateDosen(db, kodeDosen, namaDosen)
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

func DeleteDosenView(db *sql.DB) {
	ReadDosenView(db)
	fmt.Printf("Masukkan nim yang ingin di hapus :")
	kodeDosen := utils.Scanner()
	isDeleted, err := service.DeleteDosen(db, kodeDosen)
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

func AddNilai(db *sql.DB) {
	ReadMatkulView(db)
	fmt.Printf("Masukkan kode matkul :")
	kodeMatkul := utils.Scanner()
	ReadMatkulViewByKodeMatkul(db, kodeMatkul)
	fmt.Printf("Masukkan Nim mahasiswa :")
	nim := utils.Scanner()
	fmt.Printf("Masukkan Nilai (1-100):")
	nilai, _ := strconv.Atoi(utils.Scanner())
	isAdded, err := service.AddNilai(db, nim, kodeMatkul, nilai)
	if err != nil {
		fmt.Println(err)
	} else {
		if isAdded {
			fmt.Println("nilai berhasil di inputkan")
		}
	}
	utils.Scanner()
}
