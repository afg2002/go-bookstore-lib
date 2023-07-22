package repository

import (
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"
)

var data entity.Data

type MainRepositoryImpl struct{}

func NewMainRepositoryImpl() *MainRepositoryImpl {
	return &MainRepositoryImpl{}
}

func (repository *MainRepositoryImpl) FindAll() entity.Data {

	// Get the book data
	con := db.ConnectionDB()
	sql := "SELECT * FROM buku"
	result, err := con.Query(sql)
	helper.PanicIfError(err)

	book := entity.Book{}
	var resp []entity.Book
	for result.Next() {
		err = result.Scan(&book.ID, &book.Cover, &book.Judul, &book.Harga, &book.Pengarang, &book.Kategori, &book.Penerbit, &book.Tahun, &book.Stok, &book.Deskripsi)
		helper.PanicIfError(err)

		resp = append(resp, book)
	}

	for _, e := range resp {
		//fmt.Println(i)
		data.BookData = append(data.BookData, e)
	}
	//fmt.Println(data)
	return data

}
