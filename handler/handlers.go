package handler

import (
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
