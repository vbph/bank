package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbph/bank/utils"
)

func TestCreateAccount(t *testing.T) {
	createAcc(t)
}

func TestReadAccount(t *testing.T) {
	acc := createAcc(t)

	readAcc, err := testQueries.ReadAccount(context.Background(), acc.Email)

	require.NoError(t, err)
	require.NotEmpty(t, readAcc)

	require.Equal(t, readAcc.Email, acc.Email)
	require.Equal(t, readAcc.Balance, acc.Balance)
	require.Equal(t, readAcc.CreatedAt, acc.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	acc := createAcc(t)

	newPassword := "new_password"

	hashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)

	args := ChangePasswordParams{
		ID:       acc.ID,
		Password: hashedPassword,
	}

	updatedAcc, err := testQueries.ChangePassword(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAcc)

	require.Equal(t, updatedAcc.ID, acc.ID)

	require.NoError(t, utils.VerifyPassword(newPassword, updatedAcc.Password))
}

func TestDeleteAccount(t *testing.T) {
	acc := createAcc(t)

	deletedAcc, err := testQueries.DeleteAccount(context.Background(), acc.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deletedAcc)

	require.Equal(t, deletedAcc.ID, acc.ID)
}

func createAcc(t *testing.T) Account {
	password := "init_password"

	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	args := CreateAccountParams{
		Email:    utils.RandomEmail(),
		Password: hashedPassword,
		Balance:  utils.RandomAmount(100, 1000),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.NoError(t, utils.VerifyPassword(password, acc.Password))

	return acc
}
