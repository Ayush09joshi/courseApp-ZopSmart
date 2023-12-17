package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ayush09joshi/courseApp-ZopSmart.git/models"
	"github.com/Ayush09joshi/courseApp-ZopSmart.git/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
)

func initializeHandlerTest(t *testing.T) (*store.MockStore, handler, *gofr.Gofr) {
	ctrl := gomock.NewController(t)

	mockStore := store.NewMockStore(ctrl)
	h := New(mockStore)
	app := gofr.New()

	return mockStore, h, app
}

func TestGet(t *testing.T) {
	tests := []struct {
		desc string
		resp []models.Course
		err  error
	}{
		{"success case", []models.Course{{ID: 0, Name: "sample", Price: 555, Author: "AJoshi"}}, nil},
		{"error case", nil, errors.Error("error fetching course listing")},
	}

	mockStore, h, app := initializeHandlerTest(t)

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		mockStore.EXPECT().Get(ctx).Return(tc.resp, tc.err)

		result, err := h.Get(ctx)

		if tc.err == nil {
			// Assert successful response
			assert.Nil(t, err)
			assert.NotNil(t, result)

			res, ok := result.(response)
			assert.True(t, ok)
			assert.Equal(t, tc.resp, res.Courses)
		} else {
			// Assert error response
			assert.NotNil(t, err)
			assert.Equal(t, tc.err, err)
			assert.Nil(t, result)
		}
	}
}

func TestCreate(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	input := `{"id":6,"name":"mahak","price":989,"author":"kolkata"}`
	expResp := models.Course{ID: 6, Name: "mahak", Price: 989, Author: "kolkata"}

	in := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodPost, "/create", in)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	var em models.Course

	_ = ctx.Bind(&em)

	mockStore.EXPECT().Get(ctx).Return(nil, nil).MaxTimes(2)
	mockStore.EXPECT().Create(ctx, em).Return(expResp, nil).MaxTimes(1)

	resp, err := h.Create(ctx)

	assert.Nil(t, err, "TEST,failed :success case")

	assert.Equal(t, expResp, resp, "TEST, failed:success case")
}

func TestCreate_Error(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	tests := []struct {
		desc    string
		input   string
		expResp interface{}
		err     error
	}{{"create invalid body", `{"id":6,"name":"mahak","price":989}`, models.Course{},
		errors.InvalidParam{Param: []string{"body"}}},
		{"create invalid body", `{{}}`, models.Course{}, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range tests {
		in := strings.NewReader(tc.input)
		req := httptest.NewRequest(http.MethodPost, "/create", in)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		var em models.Course

		_ = ctx.Bind(&em)

		mockStore.EXPECT().Create(ctx, em).Return(tc.expResp.(models.Course), tc.err).MaxTimes(1)

		resp, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Nil(t, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestUpdate(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	input := `{"id":6,"name":"mahak","price":989,"author":"kolkata"}`
	expResp := models.Course{ID: 6, Name: "mahak", Price: 989, Author: "kolkata"}

	in := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodPost, "/update/6", in)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	var em models.Course

	_ = ctx.Bind(&em)

	mockStore.EXPECT().Get(ctx).Return(nil, nil).MaxTimes(2)
	mockStore.EXPECT().Update(ctx, 6, em).Return(expResp, nil).MaxTimes(1)

	resp, err := h.Update(ctx)

	assert.Nil(t, err, "TEST,failed :success case")

	assert.Equal(t, expResp, resp, "TEST, failed:success case")
}

func TestUpdate_Error(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	tests := []struct {
		desc    string
		input   string
		expResp interface{}
		err     error
	}{{"update invalid body", `{"id":6,"name":"mahak","price":989}`, models.Course{},
		errors.InvalidParam{Param: []string{"body"}}},
		{"update invalid body", `{{}}`, models.Course{}, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range tests {
		in := strings.NewReader(tc.input)
		req := httptest.NewRequest(http.MethodPut, "/update/6", in)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		var em models.Course

		_ = ctx.Bind(&em)

		mockStore.EXPECT().Update(ctx, 6, em).Return(tc.expResp.(models.Course), tc.err).MaxTimes(1)

		resp, err := h.Update(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Nil(t, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestDelete(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	input := `{}`
	expResp := models.Course{ID: 6, Name: "mahak", Price: 989, Author: "kolkata"}

	in := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodDelete, "/delete/6", in)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	mockStore.EXPECT().Get(ctx).Return(nil, nil).MaxTimes(2)
	mockStore.EXPECT().Delete(ctx, 6).Return(expResp, nil).MaxTimes(1)

	resp, err := h.Delete(ctx)

	assert.Nil(t, err, "TEST,failed :success case")

	assert.Equal(t, expResp, resp, "TEST, failed:success case")
}

func TestDelete_Error(t *testing.T) {
	mockStore, h, app := initializeHandlerTest(t)

	tests := []struct {
		desc    string
		input   string
		expResp interface{}
		err     error
	}{{"delete invalid body", `{"id":6,"name":"mahak","price":989}`, models.Course{},
		errors.InvalidParam{Param: []string{"body"}}},
		{"delete invalid body", `{{}}`, models.Course{}, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range tests {
		in := strings.NewReader(tc.input)
		req := httptest.NewRequest(http.MethodDelete, "/delete/6", in)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		var em models.Course

		_ = ctx.Bind(&em)

		mockStore.EXPECT().Delete(ctx, em).Return(tc.expResp.(models.Course), tc.err).MaxTimes(1)

		resp, err := h.Delete(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Nil(t, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

