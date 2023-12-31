package backend

import (
	"fmt"
	"testing"
)

func TestInsertData(t *testing.T) {
	dbname := "penggajian"
	nama := "Abidin"
	status := "Aktif"
	jabatan := "Staff Administrasi"
	gaji := "RP 4.000.000"

	hasil := InsertDataKaryawan(dbname, nama, status, jabatan, gaji)
	fmt.Println(hasil)

}

func TestInsertHonor(t *testing.T) {
	dbname := "penggajian"
	honor := Honor{
		Nama:    "Saepul Anwar",
		Status:  "Aktif",
		Jabatan: "Staff",
		Gaji:    "RP 1.500.000",
	}
	insertedID := InsertHonor(dbname, honor)
	if insertedID == nil {
		t.Error("Failed to insert honor")
	}
}

func TestInsertJob(t *testing.T) {
	dbname := "penggajian"
	job := Job{
		Namajob: "Staff IT",
	}
	insertedID := InsertJob(dbname, job)
	if insertedID == nil {
		t.Error("Failed to insert job")
	}
}

func TestInsertTeam(t *testing.T) {
	dbname := "penggajian"
	team := Team{
		Nama:      "Uzumaki Asep",
		Deskripsi: "anjay asep",
	}
	insertedID := InsertTeam(dbname, team)
	if insertedID == nil {
		t.Error("Failed to insert team")
	}
}

func TestInsertAbout(t *testing.T) {
	dbname := "penggajian"
	about := About{
		Isi_satu: "Ini isi satu",
		Isi_dua:  "Ini isi dua",
		Image:    "image.jpg",
	}
	insertedID := InsertAbout(dbname, about)
	if insertedID == nil {
		t.Error("Failed to insert about")
	}
}

func TestGetDataKaryawan(t *testing.T) {
	stats := "Aktif"
	data := GetDataKaryawan(stats)
	fmt.Println(data)
}
func TestGetDataHonor(t *testing.T) {
	stats := "Aktif"
	data := GetDataHonor(stats)
	fmt.Println(data)
}
func TestGetDataJob(t *testing.T) {
	stats := "Staff Administrasi"
	data := GetDataJob(stats)
	fmt.Println(data)
}
func TestGetDataTeam(t *testing.T) {
	stats := "Uzumaki Memet"
	data := GetDataTeam(stats)
	fmt.Println(data)
}
