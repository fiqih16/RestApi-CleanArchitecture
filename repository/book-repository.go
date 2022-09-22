package repository

import (
	"go_api/entity"

	"gorm.io/gorm"
)

// BookRepository is a contract what book repository can do
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book) 
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepository creates a new instance of BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}