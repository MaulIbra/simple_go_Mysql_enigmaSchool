package models

import (
	"database/sql"
	"errors"

	"github.com/maulibra/enigma_school/utils"
)

type Mahasiswa struct {
	Nim           string
	NamaMahasiswa string
	JenisKelamin  int
	Alamat        string
}

func CreateMahasiswa(db *sql.DB, mahasiswa *Mahasiswa) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_MAHASISWA)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(mahasiswa.Nim, mahasiswa.NamaMahasiswa, mahasiswa.JenisKelamin, mahasiswa.Alamat); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ReadMahasiswa(db *sql.DB) (*[]Mahasiswa, error) {
	listMahasiswa := []Mahasiswa{}
	stmt, err := db.Prepare(utils.SELECT_MAHASISWA)
	if err != nil {
		return &listMahasiswa, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return &listMahasiswa, err
	}

	for rows.Next() {
		m := Mahasiswa{}
		err := rows.Scan(&m.Nim, &m.NamaMahasiswa, &m.JenisKelamin, &m.Alamat)
		if err != nil {
			return &listMahasiswa, err
		}
		listMahasiswa = append(listMahasiswa, m)
	}
	return &listMahasiswa, nil
}

func UpdateMahasiswa(db *sql.DB, mahasiswa *Mahasiswa) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.UPDATE_MAHASISWA)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(mahasiswa.NamaMahasiswa, mahasiswa.JenisKelamin, mahasiswa.Alamat, mahasiswa.Nim)
	if err != nil {
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, nim tidak di temukan")
	}
	return tx.Commit()
}

func DeleteMahasiswa(db *sql.DB, nim string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_MAHASISWA)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(nim)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, user id tidak di temukan")
	}

	return tx.Commit()
}
