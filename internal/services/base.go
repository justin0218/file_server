package services

import (
	"file_server/store"
)

type baseService struct {
	Redis  store.Redis
	Config store.Config
}
