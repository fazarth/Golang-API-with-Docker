package container

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type ContainerRepository interface {
	InsertContainer(b inventory.CONTAINER) inventory.CONTAINER //fungsi Tambah Kontainer baru
	ReadAllContainer() []inventory.CONTAINER                   //fungsi Melihat semua data Kontainer
	FindContainerByID(container_ID uint64) inventory.CONTAINER //fungsi Mencari Kontainer berdasarkan ID
	UpdateContainer(b inventory.CONTAINER) inventory.CONTAINER //fungsi Memperbaharui Kontainer
	DeleteContainer(b inventory.CONTAINER)                     //fungsi Menghapus Kontainer
}

type CONTAINER_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Location DTL
func NewContainerRepository(dbConn *gorm.DB) ContainerRepository {
	return &CONTAINER_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Create Container baru
func (db *CONTAINER_Connection) InsertContainer(b inventory.CONTAINER) inventory.CONTAINER {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Memuat semua Container
func (db *CONTAINER_Connection) ReadAllContainer() []inventory.CONTAINER {
	var Container []inventory.CONTAINER
	db.connection.Find(&Container)
	return Container
}

//Fungsi untuk Memuat Container by ID
func (db *CONTAINER_Connection) FindContainerByID(Container_ID uint64) inventory.CONTAINER {
	var Container inventory.CONTAINER
	db.connection.Find(&Container, Container_ID)
	return Container
}

//Fungsi untuk Memperbaharui Container
func (db *CONTAINER_Connection) UpdateContainer(b inventory.CONTAINER) inventory.CONTAINER {
	var Container inventory.CONTAINER
	db.connection.Find(&Container).Where("Container_ID = ?", b.CONT_ID)
	Container.CONT_ID = b.CONT_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Container
func (db *CONTAINER_Connection) DeleteContainer(b inventory.CONTAINER) {
	db.connection.Delete(&b)
}
