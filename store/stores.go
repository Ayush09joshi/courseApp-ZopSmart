package store

import (
	"fmt"

	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type course struct{}

func New() Store {
	return course{}
}

// GET
func (c course) Get(ctx *gofr.Context) ([]models.Course, error) {

	querySelect := "SELECT * FROM courses"

	rows, err := ctx.DB().QueryContext(ctx, querySelect)
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

	fmt.Println("Get ran Successfully!")

	return courses, nil
}

// CREATE
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

	fmt.Println("Inserted Successfully!")
	//Insertion of data completes here!

	// Now, use a separate SELECT query to fetch the inserted data
	queryFind := "SELECT * FROM courses WHERE id = ?"
	// Use QueryRowContext to execute the SELECT query and get a single row result
	err = ctx.DB().QueryRowContext(ctx, queryFind, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Author)

	// Handle the error if any
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}
	fmt.Println("Selected Successfully!")
	return resp, nil
}

// UPDATE
func (c course) Update(ctx *gofr.Context, id int, crt models.Course) (models.Course, error) {
	var resp models.Course

	queryUpdate := "UPDATE courses SET id=?, name=?, price=?, author=? WHERE id=?"
	// Execute the INSERT query
	_, err := ctx.DB().ExecContext(ctx, queryUpdate, crt.ID, crt.Name, crt.Price, crt.Author, id)

	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}

	fmt.Println("Updated Successfully!")

	// Now, use a separate SELECT query to fetch the inserted data
	queryFind := "SELECT id, name, price, author FROM courses WHERE id = ?"

	// Use QueryRowContext to execute the SELECT query and get a single row result
	err = ctx.DB().QueryRowContext(ctx, queryFind, crt.ID).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Author)

	// Handle the error if any
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}
	fmt.Println("Selected Successfully!")
	return resp, nil
}

// DELETE
func (c course) Delete(ctx *gofr.Context, id int) (models.Course, error) {
	var resp models.Course

	//Now to Delete that Desired Row.
	queryDelete := "DELETE FROM courses WHERE id=?"
	_, err := ctx.DB().ExecContext(ctx, queryDelete, id)
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}
	fmt.Println("Deleted Successfully!")
	return resp, nil
}
