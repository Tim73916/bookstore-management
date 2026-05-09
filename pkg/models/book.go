package models

import (
	"bookstore/pkg/config"
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) CreateBook() (*Book, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("database connection not established")
	}

	if b.Name == "" {
		return nil, errors.New("book name cannot be empty")
	}

	result := db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}

	return b, nil
}

func GetAllBooks() ([]Book, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("database connection not established")
	}

	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("database connection not established")
	}

	var book Book
	result := db.Where("ID = ?", id).First(&book)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}

	return &book, nil
}

func DeleteBook(id int64) error {
	db := config.GetDB()
	if db == nil {
		return errors.New("database connection not established")
	}

	result := db.Where("ID = ?", id).Delete(&Book{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

func UpdateBook(book *Book) (*Book, error) {
	db := config.GetDB()
	if db == nil {
		return nil, errors.New("database connection not established")
	}

	result := db.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func MigrateDB() error {
	db := config.GetDB()
	if db == nil {
		return errors.New("database connection not established")
	}

	return db.AutoMigrate(&Book{})
}
