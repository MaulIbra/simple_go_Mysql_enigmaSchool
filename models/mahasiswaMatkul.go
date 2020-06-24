package models

import (
	"database/sql"

	"github.com/maulibra/enigma_school/utils"
)

type MahasiswaMatkul struct {
	Mahasiswa  Mahasiswa
	MataKuliah MataKuliah
	Nilai      string
	NilaiAngka int
}

func CreateMahasiswaMatkul(db *sql.DB, mahasiswaMatkul *MahasiswaMatkul) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_MAHASISWA_MATKUL)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(mahasiswaMatkul.Mahasiswa.Nim, mahasiswaMatkul.MataKuliah.KodeMatkul, mahasiswaMatkul.Nilai); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ReadMahasiswaMatkul(db *sql.DB, nim string) (*[]MahasiswaMatkul, error) {
	listMahasiswaMatkul := []MahasiswaMatkul{}
	stmt, err := db.Prepare(utils.SELECT_MAHASISWA_MATKUL)
	if err != nil {
		return &listMahasiswaMatkul, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(nim)
	if err != nil {
		return &listMahasiswaMatkul, err
	}

	for rows.Next() {
		m := MahasiswaMatkul{}
		err := rows.Scan(&m.Mahasiswa.Nim, &m.MataKuliah.KodeMatkul, &m.MataKuliah.NamaMatkul, &m.MataKuliah.Dosen.KodeDosen, &m.Nilai)
		if err != nil {
			return &listMahasiswaMatkul, err
		}
		listMahasiswaMatkul = append(listMahasiswaMatkul, m)
	}
	return &listMahasiswaMatkul, nil
}

func UpdateMahasiswaMatkul(db *sql.DB, nim, kodeMatkul, nilai string, nilaiNumber int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_MAHASISWA_MATKUL)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(nilai, nilaiNumber, kodeMatkul, nim); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ReadMahasiswaMatkulByKode(db *sql.DB, kodeMatkul string) (*[]MahasiswaMatkul, error) {
	listMahasiswaMatkul := []MahasiswaMatkul{}
	stmt, err := db.Prepare(utils.SELECT_MATKUL_BY_KODE)
	if err != nil {
		return &listMahasiswaMatkul, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(kodeMatkul)
	if err != nil {
		return &listMahasiswaMatkul, err
	}

	for rows.Next() {
		m := MahasiswaMatkul{}
		err := rows.Scan(&m.Mahasiswa.Nim, &m.MataKuliah.KodeMatkul, &m.MataKuliah.NamaMatkul, &m.MataKuliah.Dosen.KodeDosen, &m.Nilai)
		if err != nil {
			return &listMahasiswaMatkul, err
		}
		listMahasiswaMatkul = append(listMahasiswaMatkul, m)
	}
	return &listMahasiswaMatkul, nil
}

func ReadMahasiswaMatkulAverage(db *sql.DB) (*[]MahasiswaMatkul, error) {
	listMahasiswaMatkul := []MahasiswaMatkul{}
	stmt, err := db.Prepare(utils.SELECT_REPORT_MATKUL)
	if err != nil {
		return &listMahasiswaMatkul, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return &listMahasiswaMatkul, err
	}

	for rows.Next() {
		m := MahasiswaMatkul{}
		err := rows.Scan(&m.MataKuliah.KodeMatkul, &m.MataKuliah.NamaMatkul, &m.MataKuliah.Dosen.KodeDosen, &m.Nilai)
		if err != nil {
			return &listMahasiswaMatkul, err
		}
		listMahasiswaMatkul = append(listMahasiswaMatkul, m)
	}
	return &listMahasiswaMatkul, nil
}
