package service

import (
	"database/sql"
	"time"

	"github.com/maulibra/enigma_school/models"
)

func AddMahasiswa(db *sql.DB, namaMahasiswa, alamat string, jenisKelamin int) (bool, error) {

	mahasiswa := models.Mahasiswa{
		Nim:           time.Now().Format("20060102150405"),
		NamaMahasiswa: namaMahasiswa,
		JenisKelamin:  jenisKelamin,
		Alamat:        alamat,
	}

	err := models.CreateMahasiswa(db, &mahasiswa)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadMahasiswa(db *sql.DB) (*[]models.Mahasiswa, error) {
	listMahasiswa, err := models.ReadMahasiswa(db)
	if err != nil {
		return listMahasiswa, err
	}
	return listMahasiswa, nil
}

func UpdateMahasiswa(db *sql.DB, nim, namaMahasiswa, alamat string, jenisKelamin int) (bool, error) {
	mahasiswa := models.Mahasiswa{
		Nim:           nim,
		NamaMahasiswa: namaMahasiswa,
		JenisKelamin:  jenisKelamin,
		Alamat:        alamat,
	}

	err := models.UpdateMahasiswa(db, &mahasiswa)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteMahasiswa(db *sql.DB, nim string) (bool, error) {
	err := models.DeleteMahasiswa(db, nim)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AmbilMatakuliah(db *sql.DB, nim, kodeMatkul, nilai string) (bool, error) {
	mahasiswaMatkul := models.MahasiswaMatkul{
		Mahasiswa:  models.Mahasiswa{Nim: nim},
		MataKuliah: models.MataKuliah{KodeMatkul: kodeMatkul},
		Nilai:      nilai,
	}

	err := models.CreateMahasiswaMatkul(db, &mahasiswaMatkul)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ListMataKuliahByNim(db *sql.DB, nim string) (*[]models.MahasiswaMatkul, error) {
	listMatkul, err := models.ReadMahasiswaMatkul(db, nim)
	if err != nil {
		return listMatkul, err
	}
	return listMatkul, nil
}
