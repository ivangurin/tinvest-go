package history_repo_test

import (
	"context"
	"testing"
	"tinvest-go/internal/pkg/suite_provider"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestHistory(t *testing.T) {

	type testCase struct {
		Name    string
		UserID  int64
		Command string
	}

	testCases := []testCase{
		{
			Name:    "Create history record 1",
			UserID:  int64(gofakeit.Uint32()),
			Command: gofakeit.Sentence(5),
		},
		{
			Name:    "Create history record 2",
			UserID:  int64(gofakeit.Uint32()),
			Command: gofakeit.Sentence(5),
		},
		{
			Name:    "Create history record 3",
			UserID:  int64(gofakeit.Uint32()),
			Command: gofakeit.Sentence(5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			sp, cancel := suite_provider.NewSuiteProvider()
			defer cancel()

			recordID, err := sp.GetHistoryRepo(ctx).CreateRecord(ctx, tc.UserID, tc.Command)
			require.NoError(t, err)
			require.NotEmpty(t, recordID)

			historyRecord, err := sp.GetHistoryRepo(ctx).GetRecords(ctx, recordID)
			require.NoError(t, err)
			require.NotNil(t, historyRecord)
			require.Equal(t, tc.UserID, historyRecord.UserID)
			require.Equal(t, tc.Command, historyRecord.Command)
			require.False(t, historyRecord.CreatedAt.IsZero())
		})
	}
}
