package views

import (
	"database/sql"
	"fmt"

	"github.com/maulibra/enigma_school/service"
	"github.com/maulibra/enigma_school/utils"
)

func ReportingByMataKuliah(db *sql.DB) {
	ReadMatkulView(db)
	fmt.Println("Masukkan Kode Mata Kuliah :")
	kodeMatkul := utils.Scanner()
	ReadMatkulViewByKodeMatkul(db, kodeMatkul)
}

func ReportingAverageMatkul(db *sql.DB) {
	fmt.Println("=========================================================================================================")
	fmt.Println("=                                        REPORTING MATKUL                                               =")
	fmt.Println("=========================================================================================================")
	fmt.Printf("%-25s \t%-25s \t%-25s \t%-10s\n", "KODE MATKUL", "NAMA MATKUL", "KODE DOSEN", "RATA-RATA")
	listMatkul, err := service.ReadMahasiswaMatkulAverage(db)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range *listMatkul {
			fmt.Printf("%-25s \t%-25s \t%-25s \t%-10v\n", val.MataKuliah.KodeMatkul, val.MataKuliah.NamaMatkul, val.MataKuliah.Dosen.KodeDosen, val.Nilai)
		}
	}
}
