package service

import "perpustakaan/entity"

type MainService interface {
	FindAll() entity.Data
}
