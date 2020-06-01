// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/maei/mongodb_mock_hexa/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// DAOInterfaceArticle is an autogenerated mock type for the DAOInterfaceArticle type
type DAOInterfaceArticle struct {
	mock.Mock
}

// FindArticleByID provides a mock function with given fields: ctx, id
func (_m *DAOInterfaceArticle) FindArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Article); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreArticle provides a mock function with given fields: ctx, article
func (_m *DAOInterfaceArticle) StoreArticle(ctx context.Context, article *domain.Article) error {
	ret := _m.Called(ctx, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
