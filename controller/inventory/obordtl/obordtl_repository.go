package obordtl

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type OborDTLRepository interface {
	InsertOborDTL(b inventory.OBORDTL) inventory.OBORDTL //fungsi Tambah Lokasi baru
	ReadAllOborDTL() []inventory.OBORDTL                 //fungsi Melihat semua data Lokasi
	FindOborDTLByID(task_id uint64) inventory.OBORDTL    //fungsi Mencari Lokasi berdasarkan ID
	UpdateOborDTL(b inventory.OBORDTL) inventory.OBORDTL //fungsi Memperbaharui Lokasi
	DeleteOborDTL(b inventory.OBORDTL)                   //fungsi Menghapus Lokasi
}

//Fungsi untuk membuat struct baru sebagai koneksi ke DB dengan GORM
type OBORDTL_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Lokasi
func NewOborDTLRepository(dbConn *gorm.DB) OborDTLRepository {
	return &OBORDTL_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Menambahkan Lokasi baru
func (db *OBORDTL_Connection) InsertOborDTL(b inventory.OBORDTL) inventory.OBORDTL {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Melihat semua Lokasi
func (db *OBORDTL_Connection) ReadAllOborDTL() []inventory.OBORDTL {
	var oborDTL []inventory.OBORDTL
	db.connection.Find(&oborDTL)
	return oborDTL
}

//Fungsi untuk Melihat Lokasi by ID
func (db *OBORDTL_Connection) FindOborDTLByID(task_id uint64) inventory.OBORDTL {
	var oborDTL inventory.OBORDTL
	db.connection.Find(&oborDTL, task_id)
	return oborDTL
}

//Fungsi untuk Memperbaharui Lokasi
func (db *OBORDTL_Connection) UpdateOborDTL(b inventory.OBORDTL) inventory.OBORDTL {
	var oborDTL inventory.OBORDTL
	db.connection.Find(&oborDTL).Where("task_id = ?", b.TASK_ID)
	oborDTL.TASK_ID = b.TASK_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Lokasi
func (db *OBORDTL_Connection) DeleteOborDTL(b inventory.OBORDTL) {
	db.connection.Delete(&b)
}
