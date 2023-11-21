package usecase

import (
	// jwt "menu-service/component"
	menu "menu-service/interfaces"
)

type menuUC struct {
	menuRepo menu.Repository
	// tokenprovider jwt.TokenProvider
	cacheRepo menu.CacheRepository
}

func NewMenuUC(menuRepo menu.Repository, cacheRepo menu.CacheRepository) *menuUC {
	return &menuUC{menuRepo: menuRepo, cacheRepo: cacheRepo}
}
