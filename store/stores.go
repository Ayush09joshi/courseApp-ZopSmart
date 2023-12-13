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


//GET
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

	return courses, nil
}


//CREATE
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
	//Insertion of data completes here!


	// Now, use a separate SELECT query to fetch the inserted data
	querySelect := "SELECT * FROM courses WHERE id = ?"

	// Use QueryRowContext to execute the SELECT query and get a single row result
	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Author)

	// Handle the error if any
	if err != nil {
		return models.Course{}, errors.DB{Err: err}
	}

	
	return resp, nil
}


//UPDATE
func (c course) Update(ctx *gofr.Context, id int, upd models.Course) (models.Course, error) {
	var resp models.Course

	queryUpdate := "UPDATE courses SET name=?, price=?, author=? WHERE id=?"

	_, err := ctx.DB().ExecContext(ctx, queryUpdate, upd.Name, upd.Price, upd.Author, id)
	if err != nil {
		// return models.Course{}, errors.DB{Err: err}
		panic(err)
	}

	fmt.Println("Updated Successfully!!")

	return resp, nil
}


//DELETE
func (c course) Delete(ctx *gofr.Context, id int) (models.Course, error) {
	var resp models.Course

	//Need to find the entry before we can delete it.
	querySelect := "SELECT * FROM courses WHERE id=?"
	_ = ctx.DB().QueryRowContext(ctx, querySelect, id).Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Author)
	//HERE WE FOUND OUR DESIRED ROW

	//Now to Delete that Desired Row.
	queryDelete := "DELETE FROM courses WHERE id=?"
	_, err := ctx.DB().ExecContext(ctx, queryDelete, id)
	if err != nil {
		// return models.Course{}, errors.DB{Err: err}
		panic(err)
	}

	return resp, nil
}           