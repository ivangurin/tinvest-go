package user_repo_test

import (
	"context"
	"testing"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/suite_provider"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	type testCase struct {
		Name       string
		UserID     int64
		UserExists bool
		User       *model.User
	}

	testCases := []testCase{
		{
			Name:       "User not exists",
			UserID:     int64(gofakeit.Uint32()),
			UserExists: false,
		},
		{
			Name:       "User exists",
			UserID:     int64(gofakeit.Uint32()),
			UserExists: true,
			User: &model.User{
				ChatID:    int64(gofakeit.Uint32()),
				Username:  gofakeit.Username(),
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Role:      gofakeit.Word(),
				Token:     gofakeit.Sentence(5),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			sp, cancel := suite_provider.NewSuiteProvider()
			defer cancel()

			if !tc.UserExists {
				exists, err := sp.GetUserRepo(ctx).IsUserExists(ctx, tc.UserID)
				require.NoError(t, err)
				require.False(t, exists)

				user, err := sp.GetUserRepo(ctx).GetUserByID(ctx, tc.UserID)
				require.NoError(t, err)
				require.Nil(t, user)

				return
			}

			tc.User.ID = tc.UserID

			user, err := sp.GetUserRepo(ctx).CreateUser(ctx, tc.User)
			require.NoError(t, err)
			require.NotNil(t, user)

			exists, err := sp.GetUserRepo(ctx).IsUserExists(ctx, tc.UserID)
			require.NoError(t, err)
			require.True(t, exists)

			user, err = sp.GetUserRepo(ctx).GetUserByID(ctx, tc.UserID)
			require.NoError(t, err)
			require.NotNil(t, user)
			require.Equal(t, tc.User.Username, user.Username)
			require.Equal(t, tc.User.FirstName, user.FirstName)
			require.Equal(t, tc.User.LastName, user.LastName)
			require.Equal(t, tc.User.Role, user.Role)
			require.Equal(t, tc.User.Token, user.Token)
			require.NotNil(t, user.CreatedAt)
			require.Nil(t, user.ChangedAt)
			require.Nil(t, user.DeletedAt)

			tc.User = user
			tc.User.ChatID = int64(gofakeit.Uint32())
			tc.User.Username = gofakeit.Username()
			tc.User.FirstName = gofakeit.FirstName()
			tc.User.LastName = gofakeit.LastName()
			tc.User.Role = gofakeit.Word()
			tc.User.Token = gofakeit.Sentence(5)

			err = sp.GetUserRepo(ctx).UpdateUser(ctx, tc.User)
			require.NoError(t, err)

			user, err = sp.GetUserRepo(ctx).GetUserByID(ctx, tc.UserID)
			require.NoError(t, err)
			require.NotNil(t, user)
			require.Equal(t, tc.User.Username, user.Username)
			require.Equal(t, tc.User.FirstName, user.FirstName)
			require.Equal(t, tc.User.LastName, user.LastName)
			require.Equal(t, tc.User.Role, user.Role)
			require.Equal(t, tc.User.Token, user.Token)
			require.NotNil(t, user.CreatedAt)
			require.Equal(t, tc.User.CreatedAt, user.CreatedAt)
			require.NotNil(t, user.ChangedAt)
			require.Nil(t, user.DeletedAt)

			tc.User = user
			err = sp.GetUserRepo(ctx).DeleteUser(ctx, tc.UserID)
			require.NoError(t, err)

			user, err = sp.GetUserRepo(ctx).GetUserByID(ctx, tc.UserID)
			require.NoError(t, err)
			require.NotNil(t, user)
			require.Equal(t, tc.User.Username, user.Username)
			require.Equal(t, tc.User.FirstName, user.FirstName)
			require.Equal(t, tc.User.LastName, user.LastName)
			require.Equal(t, tc.User.Role, user.Role)
			require.Equal(t, tc.User.Token, user.Token)
			require.NotNil(t, user.CreatedAt)
			require.Equal(t, tc.User.CreatedAt, user.CreatedAt)
			require.NotNil(t, user.ChangedAt)
			require.Equal(t, tc.User.ChangedAt, user.ChangedAt)
			require.NotNil(t, user.DeletedAt)
		})
	}
}
