package api

import (
	db "github.com/symyzi/financial-helper/db/gen"
	"github.com/symyzi/financial-helper/util"
)

func RandomCategory() db.Category {
	return db.Category{
		ID:   util.RandomInt(1, 1000),
		Name: util.RandomString(6),
	}
}
