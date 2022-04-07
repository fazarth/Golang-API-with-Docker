package oborhdr

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type OborHDRRepository interface {
	InsertOborHDR(b inventory.OBORHDR) inventory.OBORHDR //fungsi Tambah Lokasi baru
	ReadAllOborHDR() []inventory.OBORHDR                 //fungsi Melihat semua data Lokasi
	FindOborHDRByID(task_id uint64) inventory.OBORHDR    //fungsi Mencari Lokasi berdasarkan ID
	UpdateOborHDR(b inventory.OBORHDR) inventory.OBORHDR //fungsi Memperbaharui Lokasi
	DeleteOborHDR(b inventory.OBORHDR)                   //fungsi Menghapus Lokasi
}

//Fungsi untuk membuat struct baru sebagai koneksi ke DB dengan GORM
type OBORHDR_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Lokasi
func NewOborHDRRepository(dbConn *gorm.DB) OborHDRRepository {
	return &OBORHDR_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Menambahkan Lokasi baru
func (db *OBORHDR_Connection) InsertOborHDR(b inventory.OBORHDR) inventory.OBORHDR {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Melihat semua Lokasi
func (db *OBORHDR_Connection) ReadAllOborHDR() []inventory.OBORHDR {
	var oborHDR []inventory.OBORHDR
	db.connection.Find(&oborHDR)
	return oborHDR
}

//Fungsi untuk Melihat Lokasi by ID
func (db *OBORHDR_Connection) FindOborHDRByID(task_id uint64) inventory.OBORHDR {
	var oborHDR inventory.OBORHDR
	db.connection.Find(&oborHDR, task_id)
	return oborHDR
}

//Fungsi untuk Memperbaharui Lokasi
func (db *OBORHDR_Connection) UpdateOborHDR(b inventory.OBORHDR) inventory.OBORHDR {
	var oborHDR inventory.OBORHDR
	db.connection.Find(&oborHDR).Where("task_id = ?", b.ORDER_ID)
	oborHDR.ORDER_ID = b.ORDER_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Lokasi
func (db *OBORHDR_Connection) DeleteOborHDR(b inventory.OBORHDR) {
	db.connection.Delete(&b)
}
