package models

import (
	"github.com/piyushbihani/go-bookstore-crud/pkg/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	// Catergory []string `json:"catergory"`
}

func init() {
	config.ConnectDataBase()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	if b.ID == 0 {
		DB.Create(&b)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	DB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID=?", Id).Find(&getBook)
	if db.RowsAffected == 0 {
		return nil, db
	}
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var deleteBook Book
	// db := DB.Where("ID=?", Id).Delete(&deleteBook)
	db := DB.Delete(&deleteBook, Id)
	if db.RowsAffected == 0 {
		return nil
	}
	return &deleteBook
}
