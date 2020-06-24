package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/maulibra/enigma_school/utils"
	"github.com/maulibra/enigma_school/views"
)

func connectDB() (*sql.DB, error) {

	dbUser := utils.GetEnv("--dbUser", "root")
	dbPassword := utils.GetEnv("--dbPassword", "")
	dbHost := utils.GetEnv("--dbHost", "127.0.0.1")
	dbPort := utils.GetEnv("--dbPort", "3306")
	schemaName := utils.GetEnv("--dbSchema", "live_code_db")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, schemaName)

	db, err := utils.InitDB("mysql", dataSource)

	return db, err
}

func main() {
	db, _ := connectDB()
	for {
		utils.ClearScreen()
		utils.MainScreen()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			MahasiswaScreen(db)
			break
		case "2":
			MatkulScreen(db)
			break
		case "3":
			DosenScreen(db)
			break
		case "4":
			views.AddNilai(db)
			break
		case "5":
			LaporanScreen(db)
			break
		case "6":
			os.Exit(0)
		default:
			break
		}
	}
}

func MahasiswaScreen(db *sql.DB) {
	for {
		utils.ClearScreen()
		utils.MahasiswaScreen()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			utils.ClearScreen()
			views.AddMahasiswaView(db)
			break
		case "2":
			utils.ClearScreen()
			views.DeleteMahasiswa(db)
			break
		case "3":
			utils.ClearScreen()
			views.ReadMahasiswa(db)
			utils.Scanner()
			break
		case "4":
			utils.ClearScreen()
			views.AmbilMatakuliah(db)
			break
		case "5":
			utils.ClearScreen()
			views.ListMatkulByNim(db)
			break
		case "6":
			utils.ClearScreen()
			views.UpdateMahasiswa(db)
			break
		default:
			main()
		}

	}
}

func DosenScreen(db *sql.DB) {
	for {
		utils.ClearScreen()
		utils.DosenScreen()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			utils.ClearScreen()
			views.AddDosenView(db)
			break
		case "2":
			utils.ClearScreen()
			views.DeleteDosenView(db)
			break
		case "3":
			utils.ClearScreen()
			views.ReadDosenView(db)
			utils.Scanner()
			break
		case "4":
			utils.ClearScreen()
			views.UpdateDosenView(db)
			break
		default:
			main()
		}

	}
}

func MatkulScreen(db *sql.DB) {
	for {
		utils.ClearScreen()
		utils.MatkulScreen()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			utils.ClearScreen()
			views.AddMatkulView(db)
			break
		case "2":
			utils.ClearScreen()
			views.DeleteMatkulView(db)
			break
		case "3":
			utils.ClearScreen()
			views.ReadMatkulView(db)
			utils.Scanner()
			break
		case "4":
			utils.ClearScreen()
			views.UpdateMatkulView(db)
		default:
			main()
		}

	}
}

func LaporanScreen(db *sql.DB) {
	for {
		utils.ClearScreen()
		utils.LaporanScreen()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			utils.ClearScreen()
			views.ReportingByMataKuliah(db)
			utils.Scanner()
			break
		case "2":
			utils.ClearScreen()
			views.ReportingAverageMatkul(db)
			utils.Scanner()
			break
		default:
			main()
		}

	}
}
