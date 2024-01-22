package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/sql/database"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"user_id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}

func databaseCategoryToCategory(category database.Category) Category {
	return Category{
		ID:        category.ID,
		Name:      category.Name,
		UserID:    category.UserID,
		Createdat: category.Createdat,
		Updatedat: category.Updatedat,
	}
}

func DatabaseCategoriesToCategories(categories []database.Category) []Category {
	categorySlice := make([]Category, 0)

	for _, category := range categories {
		categorySlice = append(categorySlice, databaseCategoryToCategory(category))
	}

	return categorySlice
}
