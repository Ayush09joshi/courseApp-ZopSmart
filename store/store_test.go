package store

import (
	"context"
	"testing"

	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func TestCoreLayer(*testing.T) {
	app := gofr.New()

	//Seeder Initilizing
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	createTable(app)
}

func createTable(app *gofr.Gofr) {
	//Droping Table
	_, err := app.DB().Exec("DROP TABLE IF EXISTS courses")
	if err != nil {
		return
	}

	//Creating Table
	_, err = app.DB().Exec("CREATE TABLE IF NOT EXISTS courses (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, price INT NOT NULL, author VARCHAR(255) NOT NULL)")
	if err != nil {
		return
	}
}


// CREATE
func TestCreate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		course  models.Course
		mockErr error
		err     error
	}{
		{"Valid case", models.Course{ID: 2, Name: "test123", Price: 999, Author: "AJ09"}, nil, nil},
		{"DB error", models.Course{ID: 6, Name: "test323", Price: 999, Author: "AJ09"}, errors.DB{}, errors.DB{Err: errors.DB{}}},
	}

	for i, tc := range tests {
		// Set up the expectations for the INSERT query.
		mock.ExpectExec("INSERT INTO courses (id, name, price, author) VALUES (?, ?, ?, ?)").
			WithArgs(tc.course.ID, tc.course.Name, tc.course.Price, tc.course.Author).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		// Set up the expectations for the SELECT query.
		rows := sqlmock.NewRows([]string{"id", "name", "price", "author"}).
			AddRow(tc.course.ID, tc.course.Name, tc.course.Price, tc.course.Author)

		mock.ExpectQuery("SELECT * FROM courses WHERE id = ?").
			WithArgs(tc.course.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		store := New()

		resp, err := store.Create(ctx, tc.course)

		ctx.Logger.Log(resp)

		assert.IsType(t, tc.err, err, "TEST[%d], failed. \n%s", i, tc.desc)	
	}
}


// GET
func TestGet(t *testing.T) {

	ctx := gofr.NewContext(nil, nil, gofr.New())

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		courses []models.Course
		mockErr error
		err     error
	}{
		{"Valid case with courses", []models.Course{
			{ID: 1, Name: "AA1", Price: 999, Author: "AJ09"},
			{ID: 2, Name: "AA2", Price: 1999, Author: "A09J"},
		}, nil, nil},
		{"Valid case with no courses", []models.Course{}, nil, nil},
		{"Error case", nil, errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {

		rows := sqlmock.NewRows([]string{"id", "name", "price", "author"})

		for _, cor := range tc.courses {
			rows.AddRow(cor.ID, cor.Name, cor.Price, cor.Author)
		}

		mock.ExpectQuery("SELECT * FROM courses").
			WillReturnRows(rows).WillReturnError(tc.mockErr)

		store := New()
		resp, err := store.Get(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.courses, resp, "TEST[%d], failed.\n%s", i, tc.desc)

	}
}


// UPDATE
func TestUpdate(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}
	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		course  models.Course
		id      int
		mockErr error
		err     error
	}{
		{"Valid case", models.Course{ID: 2, Name: "test1123", Price: 9999, Author: "A09"}, 2, nil, nil},
		{"DB error", models.Course{ID: 2, Name: "test1123", Price: 999, Author: "AJ09"}, 1, errors.DB{}, errors.DB{Err: errors.DB{}}},
	}

	for i, tc := range tests {
		// Set up the expectations for the INSERT query.
		mock.ExpectExec("UPDATE courses SET id=?, name=?, price=?, author=? WHERE id=?").
			WithArgs(tc.course.ID, tc.course.Name, tc.course.Price, tc.course.Author, tc.id).
			WillReturnResult(sqlmock.NewResult(int64(tc.id), 1)).
			WillReturnError(tc.mockErr)

		// Set up the expectations for the SELECT query.
		rows := sqlmock.NewRows([]string{"id","name", "price", "author"}).
			AddRow(tc.course.ID, tc.course.Name, tc.course.Price, tc.course.Author)

			mock.ExpectQuery("SELECT id, name, price, author FROM courses WHERE id = ?").
			WithArgs(tc.course.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)


		store := New()
		resp, err := store.Update(ctx, tc.id, tc.course)
		ctx.Logger.Log(resp)
		assert.IsType(t, tc.err, err, "TEST[%d], failed. \n%s", i, tc)
	}
}


//Delete
func TestDelete(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}
	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	
	tests := []struct {
		desc    string
		course  models.Course
		id      int
		mockErr error
		err     error
	}{
		{"Valid case", models.Course{ID: 2, Name: "test123", Price: 999, Author: "AJ09"}, 2, nil, nil},
		{"DB error", models.Course{ID: 2, Name: "test23", Price: 999, Author: "AJ09"}, 1, errors.DB{}, errors.DB{Err: errors.DB{}}},
	}

	for i, tc := range tests {

		// rows := sqlmock.NewRows([]string{"id", "name", "price", "author"}).
		// 	AddRow(tc.course.ID, tc.course.Name, tc.course.Price, tc.course.Author)

		// mock.ExpectQuery("SELECT * FROM courses WHERE id = ?").
		// 	WithArgs(tc.course.ID).
		// 	WillReturnRows(rows).
		// 	WillReturnError(tc.mockErr)


		// Set up the expectations for the Delete query.
		mock.ExpectExec("DELETE FROM courses WHERE id=?").
			WithArgs(tc.id).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		store := New()
		resp, err := store.Delete(ctx, tc.id)
		ctx.Logger.Log(resp)
		assert.IsType(t, tc.err, err, "TEST[%d], failed. \n%s", i, tc)
	}
}