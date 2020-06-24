package service

import (
	"database/sql"
	"time"

	"github.com/maulibra/enigma_school/models"
	"github.com/maulibra/enigma_school/utils"
)

func AddDosen(db *sql.DB, namaDosen string) (bool, error) {

	dosen := models.Dosen{
		KodeDosen: time.Now().Format("20060102150405"),
		NamaDosen: namaDosen,
	}

	err := models.CrateDosen(db, &dosen)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadDosen(db *sql.DB) (*[]models.Dosen, error) {
	listDosen, err := models.ReadDosen(db)
	if err != nil {
		return listDosen, err
	}
	return listDosen, nil
}

func UpdateDosen(db *sql.DB, kodeDosen, namaDosen string) (bool, error) {
	dosen := models.Dosen{
		KodeDosen: kodeDosen,
		NamaDosen: namaDosen,
	}

	err := models.UpdateDosen(db, &dosen)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteDosen(db *sql.DB, kodeDosen string) (bool, error) {
	err := models.DeleteDosen(db, kodeDosen)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddNilai(db *sql.DB, nim, kodeMatkul string, nilai int) (bool, error) {
	err := models.UpdateMahasiswaMatkul(db, nim, kodeMatkul, utils.KonversiNilai(nilai), nilai)
	if err != nil {
		return false, err
	}
	return true, nil
}
