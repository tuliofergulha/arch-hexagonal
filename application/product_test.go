package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/tuliofergulha/arch-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Price = -10
	isValid, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
	require.False(t, isValid)

	product.Price = 10
	product.Status = "invalid"
	isValid, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())
	require.False(t, isValid)
}
