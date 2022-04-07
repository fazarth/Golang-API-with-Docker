package location

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type LocationRepository interface {
	InsertLocation(b inventory.LOCATION) inventory.LOCATION //fungsi Tambah Lokasi baru
	ReadAllLocation() []inventory.LOCATION                  //fungsi Melihat semua data Lokasi
	FindLocationByID(location_ID uint64) inventory.LOCATION //fungsi Mencari Lokasi berdasarkan ID
	UpdateLocation(b inventory.LOCATION) inventory.LOCATION //fungsi Memperbaharui Lokasi
	DeleteLocation(b inventory.LOCATION)                    //fungsi Menghapus Lokasi
}

//Fungsi untuk membuat struct baru sebagai koneksi ke DB dengan GORM
type LOCATION_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Lokasi
func NewLocationRepository(dbConn *gorm.DB) LocationRepository {
	return &LOCATION_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Menambahkan Lokasi baru
func (db *LOCATION_Connection) InsertLocation(b inventory.LOCATION) inventory.LOCATION {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Melihat semua Lokasi
func (db *LOCATION_Connection) ReadAllLocation() []inventory.LOCATION {
	var location []inventory.LOCATION
	db.connection.Find(&location)
	return location
}

//Fungsi untuk Melihat Lokasi by ID
func (db *LOCATION_Connection) FindLocationByID(location_ID uint64) inventory.LOCATION {
	var location inventory.LOCATION
	db.connection.Find(&location, location_ID)
	return location
}

//Fungsi untuk Memperbaharui Lokasi
func (db *LOCATION_Connection) UpdateLocation(b inventory.LOCATION) inventory.LOCATION {
	var location inventory.LOCATION
	db.connection.Find(&location).Where("location_ID = ?", b.LOCN_ID)
	location.LOCN_ID = b.LOCN_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Lokasi
func (db *LOCATION_Connection) DeleteLocation(b inventory.LOCATION) {
	db.connection.Delete(&b)
}
