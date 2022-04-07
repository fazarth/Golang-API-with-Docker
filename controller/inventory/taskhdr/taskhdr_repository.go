package taskhdr

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type TaskHDRRepository interface {
	InsertTaskHDR(b inventory.TASKHDR) inventory.TASKHDR //fungsi Tambah Lokasi baru
	ReadAllTaskHDR() []inventory.TASKHDR                 //fungsi Melihat semua data Lokasi
	FindTaskHDRByID(task_id uint64) inventory.TASKHDR    //fungsi Mencari Lokasi berdasarkan ID
	UpdateTaskHDR(b inventory.TASKHDR) inventory.TASKHDR //fungsi Memperbaharui Lokasi
	DeleteTaskHDR(b inventory.TASKHDR)                   //fungsi Menghapus Lokasi
}

//Fungsi untuk membuat struct baru sebagai koneksi ke DB dengan GORM
type TASKHDR_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Lokasi
func NewTaskHDRRepository(dbConn *gorm.DB) TaskHDRRepository {
	return &TASKHDR_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Menambahkan Lokasi baru
func (db *TASKHDR_Connection) InsertTaskHDR(b inventory.TASKHDR) inventory.TASKHDR {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Melihat semua Lokasi
func (db *TASKHDR_Connection) ReadAllTaskHDR() []inventory.TASKHDR {
	var taskHDR []inventory.TASKHDR
	db.connection.Find(&taskHDR)
	return taskHDR
}

//Fungsi untuk Melihat Lokasi by ID
func (db *TASKHDR_Connection) FindTaskHDRByID(task_id uint64) inventory.TASKHDR {
	var taskHDR inventory.TASKHDR
	db.connection.Find(&taskHDR, task_id)
	return taskHDR
}

//Fungsi untuk Memperbaharui Lokasi
func (db *TASKHDR_Connection) UpdateTaskHDR(b inventory.TASKHDR) inventory.TASKHDR {
	var taskHDR inventory.TASKHDR
	db.connection.Find(&taskHDR).Where("task_id = ?", b.TASK_ID)
	taskHDR.TASK_ID = b.TASK_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Lokasi
func (db *TASKHDR_Connection) DeleteTaskHDR(b inventory.TASKHDR) {
	db.connection.Delete(&b)
}
