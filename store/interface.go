package store

import (
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"gofr.dev/pkg/gofr"
)

type Store interface {
	Get(ctx *gofr.Context) ([]models.Course, error)
	Create(ctx *gofr.Context, course models.Course) (models.Course, error)
	Update(ctx *gofr.Context, id int, course models.Course) (models.Course, error)
	Delete(ctx *gofr.Context, id int) (models.Course, error)
}
