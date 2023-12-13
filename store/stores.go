package store

import (
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type course struct{}

func New() Store {
	return course{}
}

func (c course) Get(ctx *gofr.Context) ([]models.Course, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT id,name,price,author FROM courses")
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	courses := make([]models.Course, 0)

	for rows.Next() {
		var c models.Course

		err = rows.Scan(&c.ID, &c.Name, &c.Price, &c.Author)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		courses = append(courses, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return courses, nil
}

func (c course) Create(ctx *gofr.Context, crt models.Course) (models.Course, error) {
	var resp models.Course

	queryInsert := "INSERT INTO courses (id, name, price, author) VALUES (?, ?, ?, ?)"

	// Execute the INSERT query
	result, err := ctx.DB().ExecContext(ctx, queryInsert, crt.ID, crt.Name, crt.Price, crt.Author)

	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}

	// Now, use a separate SELECT query to fetch the inserted data
	querySelect := "SELECT id, name, price, author FROM courses WHERE id = ?"

	// Use QueryRowContext to execute the SELECT query and get a single row result
	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Author)

	// Handle the error if any
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}

	
	return resp, nil
}
