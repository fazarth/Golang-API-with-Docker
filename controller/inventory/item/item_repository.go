package item

import (
	//import package dari Models
	"backend/models/inventory"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type ItemRepository interface {
	InsertItem(b inventory.ITEM) inventory.ITEM //fungsi Tambah Item baru
	ReadAllItem() []inventory.ITEM              //fungsi Melihat semua data Item
	FindItemByID(item_ID uint64) inventory.ITEM //fungsi Mencari Item berdasarkan ID
	UpdateItem(b inventory.ITEM) inventory.ITEM //fungsi Memperbaharui Item
	DeleteItem(b inventory.ITEM)                //fungsi Menghapus Item
}

type ITEM_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Location DTL
func NewItemRepository(dbConn *gorm.DB) ItemRepository {
	return &ITEM_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Create Item baru
func (db *ITEM_Connection) InsertItem(b inventory.ITEM) inventory.ITEM {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Memuat semua Item
func (db *ITEM_Connection) ReadAllItem() []inventory.ITEM {
	var Item []inventory.ITEM
	db.connection.Find(&Item)
	return Item
}

//Fungsi untuk Memuat Item by ID
func (db *ITEM_Connection) FindItemByID(Item_ID uint64) inventory.ITEM {
	var Item inventory.ITEM
	db.connection.Find(&Item, Item_ID)
	return Item
}

//Fungsi untuk Memperbaharui Item
func (db *ITEM_Connection) UpdateItem(b inventory.ITEM) inventory.ITEM {
	var Item inventory.ITEM
	db.connection.Find(&Item).Where("Item_ID = ?", b.ITEM_ID)
	Item.ITEM_ID = b.ITEM_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Item
func (db *ITEM_Connection) DeleteItem(b inventory.ITEM) {
	db.connection.Delete(&b)
}
