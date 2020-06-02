package usecase

import (
	"context"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	_mocks "github.com/maei/mongodb_mock_hexa/src/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceArticle_Store(t *testing.T) {
	storeMock := new(_mocks.RepositoryArticleInterface)
	mockGRPC := new(_mocks.APIInterfaceGRPCArticle)

	art := domain.Article{
		Name:     "Matthias",
		Quantity: 10,
	}
	t.Run("success", func(t *testing.T) {
		artTemp := art
		storeMock.On("StoreArticle", mock.Anything, mock.AnythingOfType("*domain.Article")).Return(nil).Once()
		mockGRPC.On("SendArticle", mock.Anything, mock.AnythingOfType("*domain.Article")).Return("test", nil).Once()

		u := NewServiceArticle(storeMock, mockGRPC)

		err := u.Store(context.Background(), &artTemp)

		assert.NoError(t, err)
		assert.Equal(t, art.Name, artTemp.Name)
		assert.Equal(t, "test", "test")

		storeMock.AssertExpectations(t)

	})
}
