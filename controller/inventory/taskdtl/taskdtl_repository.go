package taskdtl

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type TaskDTLRepository interface {
	InsertTaskDTL(b inventory.TASKDTL) inventory.TASKDTL //fungsi Tambah Lokasi baru
	ReadAllTaskDTL() []inventory.TASKDTL                 //fungsi Melihat semua data Lokasi
	FindTaskDTLByID(task_id uint64) inventory.TASKDTL    //fungsi Mencari Lokasi berdasarkan ID
	UpdateTaskDTL(b inventory.TASKDTL) inventory.TASKDTL //fungsi Memperbaharui Lokasi
	DeleteTaskDTL(b inventory.TASKDTL)                   //fungsi Menghapus Lokasi
}

//Fungsi untuk membuat struct baru sebagai koneksi ke DB dengan GORM
type TASKDTL_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Lokasi
func NewTaskDTLRepository(dbConn *gorm.DB) TaskDTLRepository {
	return &TASKDTL_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Menambahkan Lokasi baru
func (db *TASKDTL_Connection) InsertTaskDTL(b inventory.TASKDTL) inventory.TASKDTL {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Melihat semua Lokasi
func (db *TASKDTL_Connection) ReadAllTaskDTL() []inventory.TASKDTL {
	var taskDTL []inventory.TASKDTL
	db.connection.Find(&taskDTL)
	return taskDTL
}

//Fungsi untuk Melihat Lokasi by ID
func (db *TASKDTL_Connection) FindTaskDTLByID(task_id uint64) inventory.TASKDTL {
	var taskDTL inventory.TASKDTL
	db.connection.Find(&taskDTL, task_id)
	return taskDTL
}

//Fungsi untuk Memperbaharui Lokasi
func (db *TASKDTL_Connection) UpdateTaskDTL(b inventory.TASKDTL) inventory.TASKDTL {
	var taskDTL inventory.TASKDTL
	db.connection.Find(&taskDTL).Where("task_id = ?", b.TASK_ID)
	taskDTL.TASK_ID = b.TASK_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Lokasi
func (db *TASKDTL_Connection) DeleteTaskDTL(b inventory.TASKDTL) {
	db.connection.Delete(&b)
}
