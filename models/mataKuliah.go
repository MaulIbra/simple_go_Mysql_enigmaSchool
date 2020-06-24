package models

import (
	"database/sql"
	"errors"

	"github.com/maulibra/enigma_school/utils"
)

type MataKuliah struct {
	KodeMatkul string
	NamaMatkul string
	Dosen      Dosen
}

func CreateMatkul(db *sql.DB, matkul *MataKuliah) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_MATKUL)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(matkul.KodeMatkul, matkul.NamaMatkul, matkul.Dosen.KodeDosen); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ReadMatkul(db *sql.DB) (*[]MataKuliah, error) {
	listMatkul := []MataKuliah{}
	stmt, err := db.Prepare(utils.SELECT_MATKUL)
	if err != nil {
		return &listMatkul, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return &listMatkul, err
	}

	for rows.Next() {
		m := MataKuliah{}
		err := rows.Scan(&m.KodeMatkul, &m.NamaMatkul, &m.Dosen.KodeDosen, &m.Dosen.NamaDosen)
		if err != nil {
			return &listMatkul, err
		}
		listMatkul = append(listMatkul, m)
	}
	return &listMatkul, nil
}

func UpdateMatkul(db *sql.DB, matkul *MataKuliah) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.UPDATE_MATKUL)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(matkul.Dosen.KodeDosen, matkul.KodeMatkul)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal update kode dosen atau kode matkul tidak ditemukan")
	}
	return tx.Commit()
}

func DeleteMatkul(db *sql.DB, kodeMatkul string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_MATKUL)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(kodeMatkul)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, kode matkul tidak di temukan")
	}

	return tx.Commit()
}
