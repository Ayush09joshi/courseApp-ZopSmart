package handler

import (
	"strconv"

	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/store"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store store.Store
}

func New(s store.Store) handler {
	return handler{store: s}
}

type response struct {
	Courses []models.Course
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.Get(ctx)
	if err != nil {
		return nil, err
	}

	return response{Courses: resp}, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var cc models.Course
	if err := ctx.Bind(&cc); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, cc)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var up models.Course

	id := ctx.PathParam("id")
	intID, _ := strconv.Atoi(id)

	if err := ctx.Bind(&up); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Update(ctx, intID, up)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {

	id := ctx.PathParam("id")
	intID, _ := strconv.Atoi(id)
	resp, err := h.store.Delete(ctx, intID)

	if err != nil {
		// return nil, err
		panic(err)
	}

	return resp, nil
}
