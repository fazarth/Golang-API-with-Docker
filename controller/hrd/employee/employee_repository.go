package employee

import (
	//import package dari Models
	"backend/models/hrd"

	"gorm.io/gorm" //Import package GORM
)

//Mendeklarasikan Class dengan type Interface untuk mendefinisikan fungsi apa saja yang akan digunakan dalam file ini
type EmployeeRepository interface {
	InsertEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE        //fungsi Tambah Employee baru
	GetAllEmployee() []hrd.EMPLOYEE                    //fungsi Melihat semua data Employee
	FindEmployeeByID(container_ID uint64) hrd.EMPLOYEE //fungsi Mencari Employee berdasarkan ID
	UpdateEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE        //fungsi Memperbaharui Employee
	DeleteEmployee(b hrd.EMPLOYEE)                     //fungsi Menghapus Employee
}

type EMPLOYEE_Connection struct {
	connection *gorm.DB
}

//Fungsi untuk membuat instanse Location DTL
func NewEmployeeRepository(dbConn *gorm.DB) EmployeeRepository {
	return &EMPLOYEE_Connection{
		connection: dbConn,
	}
}

//Fungsi untuk Create Employee baru
func (db *EMPLOYEE_Connection) InsertEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

//Fungsi untuk Memuat semua Employee
func (db *EMPLOYEE_Connection) GetAllEmployee() []hrd.EMPLOYEE {
	var Employee []hrd.EMPLOYEE
	db.connection.Find(&Employee)
	return Employee
}

//Fungsi untuk Memuat Employee by ID
func (db *EMPLOYEE_Connection) FindEmployeeByID(Employee_ID uint64) hrd.EMPLOYEE {
	var Employee hrd.EMPLOYEE
	db.connection.Find(&Employee, Employee_ID)
	return Employee
}

//Fungsi untuk Memperbaharui Employee
func (db *EMPLOYEE_Connection) UpdateEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE {
	var Employee hrd.EMPLOYEE
	db.connection.Find(&Employee).Where("Employee_ID = ?", b.EMPLOYEE_ID)
	Employee.EMPLOYEE_ID = b.EMPLOYEE_ID
	db.connection.Updates(&b)
	return b
}

//Fungsi untuk Menghapus Employee
func (db *EMPLOYEE_Connection) DeleteEmployee(b hrd.EMPLOYEE) {
	db.connection.Delete(&b)
}
