package repository

import (
	"belajar-restful-api/helper"
	"belajar-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	query := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	query := `UPDATE category SET name = ? WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {

	query := `DELETE FROM category WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfErr(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {

	query := `SELECT id, name FROM category WHERE id = ?`

	result, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfErr(err)
	defer result.Close()

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)

		return category, nil

	} else {
		return category, errors.New("category is not found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {

	query := `SELECT id, name FROM category`

	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfErr(err)
	defer result.Close()

	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}
	return categories
}
