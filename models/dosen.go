package models

import (
	"database/sql"
	"errors"

	"github.com/maulibra/enigma_school/utils"
)

type Dosen struct {
	KodeDosen string
	NamaDosen string
}

func CrateDosen(db *sql.DB, dosen *Dosen) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_DOSEN)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(dosen.KodeDosen, dosen.NamaDosen); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ReadDosen(db *sql.DB) (*[]Dosen, error) {
	listDosen := []Dosen{}
	stmt, err := db.Prepare(utils.SELECT_DOSEN)
	if err != nil {
		return &listDosen, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return &listDosen, err
	}

	for rows.Next() {
		d := Dosen{}
		err := rows.Scan(&d.KodeDosen, &d.NamaDosen)
		if err != nil {
			return &listDosen, err
		}
		listDosen = append(listDosen, d)
	}
	return &listDosen, nil
}

func UpdateDosen(db *sql.DB, dosen *Dosen) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.UPDATE_DOSEN)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(dosen.NamaDosen, dosen.KodeDosen)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, kode dosen tidak di temukan")
	}
	return tx.Commit()
}

func DeleteDosen(db *sql.DB, kodeDosen string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.DELETE_DOSEN)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(kodeDosen)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete, kode dosen tidak di temukan")
	}

	return tx.Commit()
}
