package repository

import "perpustakaan/entity"

type MainRepository interface {
	FindAll() entity.Data
}
