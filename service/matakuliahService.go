package service

import (
	"database/sql"
	"time"

	"github.com/maulibra/enigma_school/models"
)

func AddMatkul(db *sql.DB, namaMatkul, kodeDosen string) (bool, error) {

	matkul := models.MataKuliah{
		KodeMatkul: time.Now().Format("20060102150405"),
		NamaMatkul: namaMatkul,
		Dosen:      models.Dosen{KodeDosen: kodeDosen},
	}

	err := models.CreateMatkul(db, &matkul)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadMatkul(db *sql.DB) (*[]models.MataKuliah, error) {
	listMatkul, err := models.ReadMatkul(db)
	if err != nil {
		return listMatkul, err
	}
	return listMatkul, nil
}

func ReadMahasiswaMatkulByKode(db *sql.DB, kodeMatkul string) (*[]models.MahasiswaMatkul, error) {
	listMatkul, err := models.ReadMahasiswaMatkulByKode(db, kodeMatkul)
	if err != nil {
		return listMatkul, err
	}
	return listMatkul, nil
}

func ReadMahasiswaMatkulAverage(db *sql.DB) (*[]models.MahasiswaMatkul, error) {
	listMatkul, err := models.ReadMahasiswaMatkulAverage(db)
	if err != nil {
		return listMatkul, err
	}
	return listMatkul, nil
}

func UpdateMatkul(db *sql.DB, kodeMatkul, kodeDosen string) (bool, error) {
	matkul := models.MataKuliah{
		KodeMatkul: kodeMatkul,
		Dosen:      models.Dosen{KodeDosen: kodeDosen},
	}

	err := models.UpdateMatkul(db, &matkul)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteMatkul(db *sql.DB, kodeMatkul string) (bool, error) {
	err := models.DeleteMatkul(db, kodeMatkul)
	if err != nil {
		return false, err
	}
	return true, nil
}
