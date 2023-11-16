package usecase

import (
	// jwt "menu-service/component"
	menu "menu-service/interfaces"
)

type menuUC struct {
	menuRepo menu.Repository
	// tokenprovider jwt.TokenProvider
}

func NewMenuUC(menuRepo menu.Repository) *menuUC {
	return &menuUC{menuRepo: menuRepo}
}
